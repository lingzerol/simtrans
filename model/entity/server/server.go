package server

const (
	MaxCachceNum = 10
)

type DeviceCacheItem struct {
	ItemID       uint64 `json:"item_id" form:"item_id"  binding:"required"`
	TransferType int64  `json:"transfer_type" form:"transfer_type"  binding:"required"`
	Content      string `json:"content" form:"content" binding:"required"`
	TimeStamp    int64  `json:"timestamp" form:"timestamp" binding:"required"`
}

type CacheRefreshParams struct {
	DeviceName string             `json:"device_name" form:"device_name" binding:"required"`
	DeviceID   uint64             `json:"device_id" form:"device_id" binding:"required"`
	CacheList  []*DeviceCacheItem `json:"cache_list" form:"cache_list" binding:"required"`
	TimeStamp  int64              `json:"timestamp" form:"timestamp" binding:"required"`
}
