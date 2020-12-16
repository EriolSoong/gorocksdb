package gorocksdb

import (
	"fmt"
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
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	writeOptions := CreateWriteOptions()
	defer writeOptions.Destroy()
	err = db.Put(writeOptions, "hello", "world")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	readOptions := CreateReadOptions()
	defer readOptions.Destroy()
	val ,err := db.Get(readOptions, "hello")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(val)
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
	defer db.Close()

	writeOptions := CreateWriteOptions()
	defer writeOptions.Destroy()

	for i := 0; i < b.N; i++ {
		err = db.Put(writeOptions, "hello", "world")
		if err != nil {
			b.Error(err.Error())
		}
	}
}