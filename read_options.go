package gorocksdb

//#cgo CFLAGS: -I${SRCDIR}/lib/include
//#include "c.h"
import "C"

type ReadOptions struct {
	cReadOpts *C.rocksdb_readoptions_t
}

func CreateReadOptions() *ReadOptions {
	readOpts := new(ReadOptions)
	readOpts.cReadOpts = C.rocksdb_readoptions_create()
	return readOpts
}

func (readOpts *ReadOptions) Destroy()  {
	C.rocksdb_readoptions_destroy(readOpts.cReadOpts)
}
