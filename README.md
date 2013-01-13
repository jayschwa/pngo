PNGo
====

PNGo reads in images and writes out PNGs.

Useful for quickly testing custom Go image decoders.
Add `import _ "custom/image/pkg"` to `main.go` and rebuild.

## Build

``` bash
export GOPATH=$PWD
go build pngo
```

## Run

``` bash
./pngo this.jpg that.gif
```
Produces `this.png` and `that.png`.
