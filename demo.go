package main

//go:generate make -f Makefile

import (
	"crypto/rand"
	"encoding/binary"
	"io/ioutil"
	"log"
	"strconv"

	"os"

	"time"

	"github.com/kiwamunet/zopflipng/binding"
)

const (
	srcPath = "testdata/demo.png"
	dirPath = "testdata/"
)

func main() {

	var b []byte
	var err error

	b, err = getImageData()
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	log.Println("ZopfliPng Starting .....")

	z := &binding.ZopfliPng{}

	// ------ zopfli pram settings ------
	z.Src = b
	z.Opt.Lossy8bit = false
	z.Opt.FilterStrategies = append(z.Opt.FilterStrategies, binding.StrategyOne)
	z.Opt.FilterStrategies = append(z.Opt.FilterStrategies, binding.StrategyTwo)
	z.Opt.FilterStrategies = append(z.Opt.FilterStrategies, binding.StrategyThree)
	z.Opt.NumIterations = 14
	z.Opt.Keepchunks = append(z.Opt.Keepchunks, "iCCP")
	// ------ zopfli pram settings ------

	var e binding.Error
	b, e = z.ZopfliPng()

	if e.Code != 0 {
		log.Printf("Error::: Code: %d / Desc: %s", e.Code, e.Description)
		return
	}

	OutputFile(b, dirPath+random()+".png")
	elapsed0 := time.Since(start)
	log.Printf("elapsed time: %.3f secs", elapsed0.Seconds())
}

func getImageData() (b []byte, e error) {
	b, e = ioutil.ReadFile(srcPath)
	return
}

func random() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func OutputFile(b []byte, path string) (e error) {
	file, err := os.Create(path)
	if err != nil {
		e = err
	}
	defer file.Close()
	file.Write(b)
	return
}
