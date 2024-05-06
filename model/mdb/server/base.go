package server

import "github.com/lingzerol/simtrans/model/mdb"

func InitServerTable() error {
	err := mdb.GetDBInstance().AutoMigrate(CacheItem{})
	if err != nil {
		return err
	}
	return nil
}
