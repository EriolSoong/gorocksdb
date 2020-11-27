package gorocksdb

//#cgo CFLAGS: -I${SRCDIR}/lib/include
//#include "c.h"
import "C"

type Options struct {
	cOpts *C.rocksdb_options_t
}

func CreateOptions() *Options {
	opts := new(Options)
	opts.cOpts = C.rocksdb_options_create()
	if opts.cOpts != nil {
		return opts
	}
	return nil
}

func (opts *Options) Destroy()  {
	C.rocksdb_options_destroy(opts.cOpts)
}

func (opts *Options) IncreaseParallelism(totalThreads int) {
	C.rocksdb_options_increase_parallelism(opts.cOpts, C.int(totalThreads))
}

func (opts *Options) OptimizeLevelStyleCompaction(memtableMemoryBudget uint64) {
	C.rocksdb_options_optimize_level_style_compaction(opts.cOpts, C.uint64_t(memtableMemoryBudget))
}

func (opts *Options) CreateIfMissing(yes bool) {
	do := 0
	if yes {
		do = 1
	}
	C.rocksdb_options_set_create_if_missing(opts.cOpts, C.uchar(do))
}