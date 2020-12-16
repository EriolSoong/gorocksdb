package gorocksdb

//#cgo CFLAGS: -I${SRCDIR}/lib/include
//#include "c.h"
import "C"
import "errors"

type WriteBatch struct {
	cwBath *C.rocksdb_writebatch_t
}

func CreateWriteBatch() *WriteBatch {
	return &WriteBatch{
		C.rocksdb_writebatch_create(),
	}
}

func CreateBatchFrom(reserved string) *WriteBatch {
	return &WriteBatch{
		C.rocksdb_writebatch_create_from(C.CString(reserved), C.size_t(len(reserved))),
	}
}
func (wBatch *WriteBatch) Count() int {
	return int(C.rocksdb_writebatch_count(wBatch.cwBath))
}

func (wBatch *WriteBatch) Clear() {
	C.rocksdb_writebatch_clear(wBatch.cwBath)
}

func (wBatch *WriteBatch) Destroy() {
	C.rocksdb_writebatch_destroy(wBatch.cwBath)
}

func (wBatch *WriteBatch) Put(key, value string) error {
	C.rocksdb_writebatch_put(wBatch.cwBath,
		C.CString(key), C.size_t(len(key)),
		C.CString(value), C.size_t(len(value)))
	return nil
}

//TODO
func (wBatch *WriteBatch) PutList() error {

	return nil
}

func (wBatch *WriteBatch) Merge(key, value string) error {
	C.rocksdb_writebatch_merge(wBatch.cwBath,
		C.CString(key), C.size_t(len(key)),
		C.CString(value), C.size_t(len(value)))
	return nil
}

func (wBatch *WriteBatch) Delete(key, value string) error {
	C.rocksdb_writebatch_delete(wBatch.cwBath,
		C.CString(key), C.size_t(len(key)))
	return nil
}

func (wBatch *WriteBatch) DeleteRange(startkey, endKey string) error {
	C.rocksdb_writebatch_delete_range(wBatch.cwBath,
		C.CString(startkey), C.size_t(len(startkey)),
		C.CString(endKey), C.size_t(len(endKey)))
	return nil
}

func (wBatch *WriteBatch) SavePoint() error {
	C.rocksdb_writebatch_set_save_point(wBatch.cwBath)
	return nil
}

func (wBatch *WriteBatch) RollbackToSavePoint() error {
	var err *C.char
	C.rocksdb_writebatch_rollback_to_save_point(wBatch.cwBath, &err)
	if err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}

func (wBatch *WriteBatch) PopSavePoint() error {
	var err *C.char
	C.rocksdb_writebatch_pop_save_point(wBatch.cwBath, &err)
	if err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}
