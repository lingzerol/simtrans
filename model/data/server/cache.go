package server

import (
	"sort"

	"github.com/lingzerol/simtrans/library/errno"
	server_entity "github.com/lingzerol/simtrans/model/entity/server"
	server_mdb "github.com/lingzerol/simtrans/model/mdb/server"
)

type CacheSrv struct {
	CacheItemDB *server_mdb.CacheItem
}

func NewCacheSrv() *CacheSrv {
	return &CacheSrv{
		CacheItemDB: &server_mdb.CacheItem{},
	}
}

func (c *CacheSrv) RefreshCache(params *server_entity.CacheRefreshParams) error {
	if params == nil {
		return errno.WrapCodeErrorf(errno.NilError, "input params nil")
	}
	sort.Slice(params.CacheList, func(i, j int) bool {
		return params.CacheList[i].TimeStamp > params.CacheList[j].TimeStamp
	})
	caches, err := c.CacheItemDB.GetDeviceCaches(params.DeviceName, params.DeviceID)
	if err != nil {
		return err
	}
	insertCaches := make([]*server_entity.DeviceCacheItem, 0)
	validNum := 0
	i, j := 0, 0
	existsCaches := make(map[uint64]*server_entity.DeviceCacheItem)
	for _, item := range caches {
		existsCaches[item.ItemID] = item
	}
	for {
		if i >= len(params.CacheList) && j >= len(caches) {
			break
		}
		var item *server_entity.DeviceCacheItem
		if i >= len(params.CacheList) {
			item = caches[j]
			j += 1
		}
		if j >= len(caches) {
			item = params.CacheList[i]
			i += 1
		}
		if item == nil {
			if caches[j].TimeStamp >= params.CacheList[i].TimeStamp {
				item = caches[j]
				j += 1
			} else {
				item = params.CacheList[i]
				i += 1
			}
		}
		if validNum < server_entity.MaxCachceNum {
			if existItem, ok := existsCaches[item.ItemID]; ok {
				if existItem.TimeStamp < item.TimeStamp {
					err := c.CacheItemDB.UpdateCache(params.DeviceName, params.DeviceID, item)
					if err != nil {
						return err
					}
				}
			} else {
				insertCaches = append(insertCaches, item)
			}
			validNum += 1
		} else if existItem, ok := existsCaches[item.ItemID]; ok {
			err := c.CacheItemDB.DeleteCache(params.DeviceName, params.DeviceID, existItem.ItemID)
			if err != nil {
				return err
			}
		}
	}
	if len(insertCaches) > 0 {
		err := c.CacheItemDB.InsertCaches(params.DeviceName, params.DeviceID, insertCaches)
		if err != nil {
			return err
		}
	}
	return nil
}
