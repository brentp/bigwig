// bigwig reads bigwig files via Devon Ryan's C library.
package bigwig

/*
#cgo CFLAGS: -g -Wall -O3 -fpic
#cgo LDFLAGS: -lcurl -lz -lm
#include "bigWig.h"
#include "stdlib.h"

*/
import "C"
import (
	"fmt"
	"unsafe"
)

// BigWig holds the methods that access the underlying C library.
type BigWig struct {
	bw   *C.bigWigFile_t
	path string
}

// Open takes a path (possibly remote) and returns a BigWig or an error
func Open(path string) (*BigWig, error) {
	bw := &BigWig{}

	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	res := C.bwOpen(cpath, nil)
	if res == nil {
		return nil, fmt.Errorf("couldn't open bigwig file: %s", path)
	}
	bw.bw = res
	return bw, nil
}

// Close
func (bw *BigWig) Close() error {
	C.bwClose(bw.bw)
	return nil
}

// Values accepts a location and returns a slice of the non-nil values in that range.
func (bw *BigWig) Values(chrom string, start int, end int) []float64 {
	cchrom := C.CString(chrom)
	defer C.free(unsafe.Pointer(cchrom))

	intervals := C.bwGetValues(bw.bw, cchrom, C.uint32_t(start), C.uint32_t(end), C.int(0))
	defer C.bwDestroyOverlappingIntervals(intervals)
	tmp := (*[1 << 30]C.float)(unsafe.Pointer(intervals.value))
	// TODO: dont copy
	m := make([]float64, int(intervals.l))
	for i := 0; i < int(intervals.l); i++ {
		m[i] = float64(tmp[i])
	}
	return m
}

// Mean accepts a location and returns the mean of the non-nil values in that range.
func (bw *BigWig) Mean(chrom string, start int, end int) float64 {
	cchrom := C.CString(chrom)
	defer C.free(unsafe.Pointer(cchrom))
	// C.int32(1) tells it to use 1 bin
	res := C.bwStats(bw.bw, cchrom, C.uint32_t(start), C.uint32_t(end), C.uint32_t(1), C.mean)
	tmp := (*[1]C.double)(unsafe.Pointer(res))
	v := float64(tmp[0])
	C.free(unsafe.Pointer(res))
	return v
}
