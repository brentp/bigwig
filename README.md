bigwig
======
bigwig reads bigwig files via Devon Ryan's C library.

## Usage

```Go
package main

import (
	"log"

	"github.com/brentp/bigwig"
)

func main() {
	bw, err := bigwig.Open("test.bw")
	log.Println(err)

	// give chrom, start, end, get []float64
	vals := bw.Values("1", 0, 9)
	log.Println(vals)

	// give chrom, start, end, get float64 of non-nil values
	mean := bw.Mean("1", 0, 9)
	log.Println(mean)
}
```

#### type BigWig

```go
type BigWig struct {
}
```

BigWig hold the methods that access the underlying C library.

#### func  Open

```go
func Open(path string) (*BigWig, error)
```
Open takes a path (possibly remote) and returns a BigWig or an error

#### func (*BigWig) Close

```go
func (bw *BigWig) Close() error
```
Close

#### func (*BigWig) Mean

```go
func (bw *BigWig) Mean(chrom string, start int, end int) float64
```
Mean accepts a location and returns the mean of the non-nil values in that
range.

#### func (*BigWig) Values

```go
func (bw *BigWig) Values(chrom string, start int, end int) []float64
```
Values accepts a location and returns a slice of the non-nil values in that
range.
