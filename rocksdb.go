package gorocksdb

//#cgo CFLAGS: -I${SRCDIR}/lib/include
//#cgo LDFLAGS: -L${SRCDIR}/lib -lrocksdb -lstdc++ -lm
//#include "c.h"
import "C"
import (
	"errors"
)

type RocksDB struct {
	cdb *C.rocksdb_t
}

// 初始化数据库
// opts: 初始化数据库之前首先要构建并配置Options，详情见Options对象
// dbPath: 数据库的存储路径，指定到文件夹
func OpenDB(opts *Options, dbPath string) (*RocksDB, error) {
	var cerr *C.char
	db := new(RocksDB)
	db.cdb = C.rocksdb_open(opts.cOpts, C.CString(dbPath), &cerr)
	if cerr != nil {
		return nil, errors.New(C.GoString(cerr))
	} else if db.cdb == nil {
		return nil, errors.New("unknown error")
	}
	return db, nil
}

func (db *RocksDB) Put(wOpts *WriteOptions, key string, value string) error {
	var cerr *C.char

	C.rocksdb_put(db.cdb, wOpts.cWriteOpts,
		C.CString(key), C.size_t(len(key)),
		C.CString(value), C.size_t(len(value)), &cerr)

	if cerr != nil {
		return errors.New(C.GoString(cerr))
	}

	return nil
}

func (db *RocksDB) Get(rOpts *ReadOptions, key string) (string, error) {
	var cerr *C.char
	var valLen C.size_t

	val := C.rocksdb_get(db.cdb, rOpts.cReadOpts, C.CString(key), C.size_t(len(key)), &valLen, &cerr)
	if cerr != nil {
		return "", errors.New(C.GoString(cerr))
	}

	return C.GoString(val), nil
}

func (db *RocksDB) Write(wOpts *WriteOptions, wBatch *WriteBatch) error  {
	var cerr *C.char
	C.rocksdb_write(db.cdb, wOpts.cWriteOpts, wBatch.cwBath, &cerr)
	if cerr != nil {
		return errors.New(C.GoString(cerr))
	}
	return nil
}

func (db *RocksDB) Delete(wOpts *WriteOptions, key string) error {
	var cerr *C.char

	C.rocksdb_delete(db.cdb, wOpts.cWriteOpts, C.CString(key), C.size_t(len(key)), &cerr)
	if cerr != nil {
		return errors.New(C.GoString(cerr))
	}

	return nil
}

func (db *RocksDB) Close() {
	C.rocksdb_close(db.cdb)
}