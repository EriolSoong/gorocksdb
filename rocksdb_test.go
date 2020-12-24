package gorocksdb

import (
	"os"
	"runtime"
	"testing"
)

func TestOpenDB(t *testing.T) {
	dbPath := "./tmp"
	options := CreateOptions()
	defer options.Destroy()
	options.IncreaseParallelism(runtime.NumCPU())
	options.OptimizeLevelStyleCompaction(0)
	options.CreateIfMissing(true)

	db, err := OpenDB(options, dbPath)
	if err != nil {
		t.Error(err.Error())
	}
	defer func() {
		db.Close()
		os.Remove(dbPath)
	}()
}

func TestRocksDB_Put(t *testing.T) {
	dbPath := "./tmp"
	options := CreateOptions()
	defer options.Destroy()
	options.IncreaseParallelism(runtime.NumCPU())
	options.OptimizeLevelStyleCompaction(0)
	options.CreateIfMissing(true)

	db, err := OpenDB(options, dbPath)
	if err != nil {
		t.Error(err.Error())
	}
	defer func() {
		db.Close()
		os.Remove(dbPath)
	}()

	writeOptions := CreateWriteOptions()
	defer writeOptions.Destroy()
	err = db.Put(writeOptions, "hello", "world")
	if err != nil {
		t.Error(err.Error())
	}

	readOptions := CreateReadOptions()
	defer readOptions.Destroy()
	val ,err := db.Get(readOptions, "hello")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(val)
}

func BenchmarkRocksDB_Put(b *testing.B) {
	dbPath := "./tmp"
	_ = os.Remove(dbPath)
	options := CreateOptions()
	defer options.Destroy()
	options.IncreaseParallelism(runtime.NumCPU())
	options.OptimizeLevelStyleCompaction(0)
	options.CreateIfMissing(true)

	db, err := OpenDB(options, dbPath)
	if err != nil {
		b.Error(err.Error())
	}
	defer func() {
		db.Close()
		os.Remove(dbPath)
	}()

	writeOptions := CreateWriteOptions()
	defer writeOptions.Destroy()

	for i := 0; i < b.N; i++ {
		err = db.Put(writeOptions, "hello", "world")
		if err != nil {
			b.Error(err.Error())
		}
	}
}