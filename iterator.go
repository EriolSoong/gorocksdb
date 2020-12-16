package gorocksdb

//#cgo CFLAGS: -I${SRCDIR}/lib/include
//#include "c.h"
import "C"
import "errors"

type Iterator struct {
	cIter *C.rocksdb_iterator_t
}

func CreateIterator(db *RocksDB, rOpts *ReadOptions) (*Iterator, error) {
	if db == nil {
		return nil, errors.New("db is nil")
	}
	if rOpts == nil {
		return nil, errors.New("rOpts is nil")
	}

	iter := new(Iterator)
	iter.cIter = C.rocksdb_create_iterator(db.cdb, rOpts.cReadOpts)
	if iter.cIter == nil {
		return nil, errors.New("create read iterator failed")
	}

	return iter, iter.error()
}

func (iter *Iterator) Destroy() {
	C.rocksdb_iter_destroy(iter.cIter)
}

func (iter *Iterator) SeekToFirst() {
	C.rocksdb_iter_seek_to_first(iter.cIter)
}

func (iter *Iterator) SeekToLast() {
	C.rocksdb_iter_seek_to_last(iter.cIter)
}

func (iter *Iterator) Seek(prefixKey string) {
	C.rocksdb_iter_seek(iter.cIter, C.CString(prefixKey), C.size_t(len(prefixKey)))
}

func (iter *Iterator) SeekForPrev(prefixKey string) {
	C.rocksdb_iter_seek_for_prev(iter.cIter, C.CString(prefixKey), C.size_t(len(prefixKey)))
}

func (iter *Iterator) Next() {
	C.rocksdb_iter_next(iter.cIter)
}

func (iter *Iterator) Prev() {
	C.rocksdb_iter_prev(iter.cIter)
}

func (iter *Iterator) Key() string {
	var key *C.char
	var size C.size_t
	key = C.rocksdb_iter_key(iter.cIter, &size)
	if key == nil {
		panic("get key failed")
	}
	return C.GoStringN(key, C.int(size))
}

func (iter *Iterator) Value() string {
	var value *C.char
	var size C.size_t
	value = C.rocksdb_iter_value(iter.cIter, &size)
	if value == nil {
		panic("get value failed")
	}
	return C.GoStringN(value, C.int(size))
}

func (iter *Iterator) Valid() bool {
	if C.rocksdb_iter_valid(iter.cIter) != 0 {
		return true
	}
	return false
}

func (iter *Iterator) error() error {
	if iter == nil || iter.cIter == nil {
		return errors.New("iterator is invalid")
	}
	var cerr *C.char
	C.rocksdb_iter_get_error(iter.cIter, &cerr)
	if cerr != nil {
		return errors.New(C.GoString(cerr))
	}
	return nil
}