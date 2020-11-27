package gorocksdb

import (
	"fmt"
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