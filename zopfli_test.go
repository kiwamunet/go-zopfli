package zopfli

import (
	"io/ioutil"
	"testing"

	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

func TestZopfliPng(t *testing.T) {

	b, err := ioutil.ReadFile("testdata/demo.png")
	if err != nil {
		t.Fatalf("%v", err)
	}

	z := &ZopfliPng{}
	z.Src = b
	z.Opt.Lossy8bit = false
	z.Opt.FilterStrategies = append(z.Opt.FilterStrategies, StrategyOne)
	z.Opt.FilterStrategies = append(z.Opt.FilterStrategies, StrategyTwo)
	z.Opt.FilterStrategies = append(z.Opt.FilterStrategies, StrategyThree)
	z.Opt.NumIterations = 14
	z.Opt.Keepchunks = append(z.Opt.Keepchunks, "iCCP")

	img, e := z.ZopfliPng()
	if e.Code != 0 {
		t.Fatalf("Error::: Code: %d / Desc: %s", e.Code, e.Description)
	}

	if err := ioutil.WriteFile("testdata/demo-result.png", img, os.ModePerm); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestZopfliPng_Goroutine(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/demo.png")
	if err != nil {
		t.Fatalf("%v", err)
	}

	z := &ZopfliPng{}
	z.Src = b

	var v int32
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&v, 1)
			s := "testdata/demo-results-" + strconv.Itoa(int(v)) + ".png"
			img, e := z.ZopfliPng()
			if e.Code != 0 {
				wg.Done()
				t.Fatalf("Error::: Code: %d / Desc: %s", e.Code, e.Description)
			}
			if err := ioutil.WriteFile(s, img, os.ModePerm); err != nil {
				wg.Done()
				t.Fatalf("%v %s", err, s)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
