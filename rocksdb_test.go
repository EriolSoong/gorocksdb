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

	db.Put(writeOptions, "key1", "value1")
	db.Put(writeOptions, "key2", "value2")
	db.Put(writeOptions, "key3", "value3")
	db.Put(writeOptions, "赵信", "50")
	db.Put(writeOptions, "阿阿", "20")
	db.Put(writeOptions, "吃货", "30")
	db.Put(writeOptions, "阿乐", "10")
	db.Put(writeOptions, "hello1", "world1")
	db.Put(writeOptions, "hello2", "world2")
	db.Put(writeOptions, "hello3", "world3")

	readOptions := CreateReadOptions()
	defer readOptions.Destroy()
	val ,err := db.Get(readOptions, "hello")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(val)
}