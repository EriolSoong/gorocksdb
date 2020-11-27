package gorocksdb

//#cgo CFLAGS: -I${SRCDIR}/lib/include
//#include "c.h"
import "C"

type RestoreOptions struct {
	cRestoreOpts *C.rocksdb_restore_options_t
}

func CreateRestoreOptions() *RestoreOptions {
	restoreOpts := new(RestoreOptions)
	restoreOpts.cRestoreOpts = C.rocksdb_restore_options_create()
	return restoreOpts
}

func (restoreOpts *RestoreOptions) Destroy()  {
	C.rocksdb_restore_options_destroy(restoreOpts.cRestoreOpts)
}
