package gorocksdb

import (
	"os"
	"runtime"
	"testing"
)

var (
	db 			*RocksDB
	dbPath 		string
	options 	*Options
	writeOpts 	*WriteOptions
	readOpts 	*ReadOptions
)

func init() {
	dbPath = "./tmp/iter"
	_ = os.Remove(dbPath)
	options = CreateOptions()
	options.CreateIfMissing(true)
	options.IncreaseParallelism(runtime.NumCPU())
	options.OptimizeLevelStyleCompaction(0)

	var data = map[string]string {
		"alia": 	"20",
		"alice": 	"30",
		"eric":		"28",
		"Eriol":	"18",
		"nora":		"18",
		"zara":		"1000",
	}

	var err error
	db, err = OpenDB(options, dbPath)
	if err != nil {
		panic(err)
	}

	writeOpts = CreateWriteOptions()
	for key, value := range data {
		err = db.Put(writeOpts, key, value)
		if err != nil {
			panic(err)
		}
	}
	readOpts = CreateReadOptions()
}

func clear() {
	iter, err := CreateIterator(db, readOpts)
	if err != nil {
		panic(err)
	}
	defer iter.Destroy()
	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
	}
}

func TestIterator_SeekToFirst(t *testing.T) {
	iter, err := CreateIterator(db, readOpts)
	if err != nil {
		t.Fatal(err)
	}
	defer iter.Destroy()

	iter.SeekToFirst()
	if !iter.Valid() || iter.Key() != "alia" {
		t.Errorf("first value is : %s, expected is : %s", iter.Key(), "alia")
	}
}

func TestIterator_SeekToLast(t *testing.T) {
	iter, err := CreateIterator(db, readOpts)
	if err != nil {
		t.Fatal(err)
	}
	defer iter.Destroy()

	if !iter.Valid() || iter.Key() != "zara" {
		t.Errorf("first value is : %s, expected is : %s", iter.Key(), "alia")
	}
}

func TestIterator_Seek(t *testing.T) {
	iter, err := CreateIterator(db, readOpts)
	if err != nil {
		t.Fatal(err)
	}
	defer iter.Destroy()

	iter.Seek("e")
	if iter.Key() != "eric" {
		t.Fatal("seek failed")
	}
}

func TestIterator_SeekForPrev(t *testing.T) {
	iter, err := CreateIterator(db, readOpts)
	if err != nil {
		t.Fatal(err)
	}
	defer iter.Destroy()

	iter.SeekForPrev("e")
	if iter.Key() != "alice" {
		t.Fatal("seek failed")
	}
}
