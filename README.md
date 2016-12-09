# go-zopfli


Go go-zopfli is a Go bind to [google/zopfli](https://github.com/google/zopfli).


# Install

```
go get github.com/kiwamunet/go-zopfli
go generate
```

# Use

```
z := &binding.ZopfliPng{}

// ------ zopfli pram settings ------
z.Src = b
z.Opt.Lossy8bit = false
z.Opt.FilterStrategies = append(z.Opt.FilterStrategies, binding.StrategyOne)
z.Opt.NumIterations = 14
z.Opt.NumIterationsLarge = 5
z.Opt.Keepchunks = append(z.Opt.Keepchunks, "iCCP")
// ------ zopfli pram settings ------

var e binding.Error
b, e = z.ZopfliPng()
if e.Code != 0 {
	log.Printf("Error::: Code: %d / Desc: %s", e.Code, e.Description)
	return
}
```

# Params


```
// Convert 16-bit per channel images to 8-bit per channel
Lossy8bit

// Filter strategies to try
FilterStrategies
    
// PNG chunks to keep
// chunks to literally copy over from the original PNG to the resulting one
Keepchunks

// Zopfli number of iterations
NumIterations

// Zopfli number of iterations on large images
NumIterationsLarge
```