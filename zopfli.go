package zopfli

/*
#cgo CPPFLAGS: -I./internal/zopfli/src/zopfli
#cgo CPPFLAGS: -I./internal/zopfli/src/zopflipng
#cgo LDFLAGS: -lstdc++

#include "internal/zopfli/src/zopflipng/zopflipng_lib.h"
#include "internal/bridge_zopflipng.h"
*/
import "C"
import "unsafe"

const (
	StrategyZero C.int = iota
	StrategyOne
	StrategyTwo
	StrategyThree
	StrategyFour
	StrategyMinSum
	StrategyEntropy
	StrategyPredefined
	StrategyBruteForce
	NumFilterStrategies /* Not a strategy but used for the size of this enum */

)

type opt struct {
	Lossy8bit          bool // convert 16-bit per channel image to 8-bit per. defalut false
	FilterStrategies   []C.int
	Keepchunks         []string
	NumIterations      int
	NumIterationsLarge int
}

type ZopfliPng struct {
	Src []byte
	Opt opt
}

func (z *ZopfliPng) ZopfliPng() ([]byte, Error) {

	if z.Src == nil {
		return z.Src, errorCode(2).sError()
	}
	o := z.setPramZopfliPng()
	defer func() {
		C.freeopts(unsafe.Pointer(&o))
	}()

	var res unsafe.Pointer
	defer C.free(res)

	var resLen C.size_t

	r := C.CZopfliPNGOptimize((*C.uchar)(unsafe.Pointer(&z.Src[0])), C.size_t(len(z.Src)), unsafe.Pointer(&o), C.int(0), (**C.uchar)(unsafe.Pointer(&res)), (*C.size_t)(unsafe.Pointer(&resLen)))
	return C.GoBytes(res, C.int(resLen)), errorCode(r).sError()

}

func (z *ZopfliPng) setPramZopfliPng() (o C.struct_CZopfliPNGOptions) {
	o = C.struct_CZopfliPNGOptions{}
	C.CZopfliPNGSetDefaults(unsafe.Pointer(&o))

	o.lossy_8bit = boolToCint(z.Opt.Lossy8bit)

	if len(z.Opt.FilterStrategies) > 0 {
		C.setfilters(&o, C.int(len(z.Opt.FilterStrategies)), (*C.int)(&z.Opt.FilterStrategies[0]))
	}

	if len(z.Opt.Keepchunks) > 0 {
		chunks := make([]*C.char, len(z.Opt.Keepchunks))
		for i, v := range z.Opt.Keepchunks {
			chunks[i] = C.CString(v)
		}
		C.setchunks(&o, C.int(len(z.Opt.Keepchunks)), (**C.char)(&chunks[0]))
		for _, v := range chunks {
			C.free(unsafe.Pointer(v))
		}
	}

	if z.Opt.NumIterations >= 1 {
		o.num_iterations = C.int(z.Opt.NumIterations)
	}

	if z.Opt.NumIterationsLarge >= 1 {
		o.num_iterations_large = C.int(z.Opt.NumIterationsLarge)
	}
	return
}

func boolToCint(b bool) C.int {
	if b {
		return C.int(0)
	}
	return C.int(1)
}
