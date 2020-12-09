package gorocksdb

//#cgo CFLAGS: -I${SRCDIR}/lib/include
//#include "c.h"
import "C"
import "errors"

type BackupEngine struct {
	cbe *C.rocksdb_backup_engine_t
}

func OpenBackupEngine(opts *Options, backupPath string) (*BackupEngine, error) {
	var cerr *C.char

	be := new(BackupEngine)
	be.cbe = C.rocksdb_backup_engine_open(opts.cOpts, C.CString(backupPath), &cerr)
	if cerr != nil {
		return nil, errors.New(C.GoString(cerr))
	}

	return be, nil
}

func (be *BackupEngine) NewBackup(db *RocksDB) error  {
	var cerr *C.char

	C.rocksdb_backup_engine_create_new_backup(be.cbe, db.cdb, &cerr)
	if cerr != nil {
		return errors.New(C.GoString(cerr))
	}

	return nil
}

func (be *BackupEngine) RestoreDBFromLatestBackup(dbPath, backupPath string, restoreOpts *RestoreOptions) error {
	var cerr *C.char

	C.rocksdb_backup_engine_restore_db_from_latest_backup(
		be.cbe, C.CString(dbPath), C.CString(backupPath), restoreOpts.cRestoreOpts, &cerr)

	if cerr != nil {
		return errors.New(C.GoString(cerr))
	}

	return nil
}