package gorocksdb

import (
	"runtime"
	"testing"
)

func TestCreateIterator(t *testing.T) {
	dbPath := "./tmp"
	options := CreateOptions()
	options.CreateIfMissing(true)
	options.IncreaseParallelism(runtime.NumCPU())
	options.OptimizeLevelStyleCompaction(0)
	defer options.Destroy()

	db, err := OpenDB(options, dbPath)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	readOpts := CreateReadOptions()
	defer readOpts.Destroy()

	iter, err := CreateIterator(db, readOpts)
	if err != nil {
		t.Error(err)
	}
	defer iter.Destroy()

	//iter.SeekForPrev("key")
	//endKey := iter.Key()
	//for iter.Seek("hello"); iter.Key() <= endKey; iter.Next() {
	//	t.Log(iter.Key(), ":", iter.Value())
	//}

	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
		t.Log(iter.Key(), ":", iter.Value())
	}
}
