package gorocksdb

//#cgo CFLAGS: -I${SRCDIR}/lib/include
//#include "c.h"
import "C"

type WriteOptions struct {
	cWriteOpts *C.rocksdb_writeoptions_t
}

func CreateWriteOptions() *WriteOptions {
	writeOpts := new(WriteOptions)
	writeOpts.cWriteOpts = C.rocksdb_writeoptions_create()
	return writeOpts
}

func (writeOpts *WriteOptions) Destroy()  {
	C.rocksdb_writeoptions_destroy(writeOpts.cWriteOpts)
}