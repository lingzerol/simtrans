package server

import (
	"errors"

	"github.com/lingzerol/simtrans/library/errno"
	server_entity "github.com/lingzerol/simtrans/model/entity/server"
	"github.com/lingzerol/simtrans/model/mdb"
	"gorm.io/gorm"
)

type CacheItem struct {
	ID           uint64 `gorm:"id;bigint(64);primary_key;AUTO_INCREMENT" json:"id" form:"id"`
	DeviceName   string `gorm:"device_name;varchar(256);not null" json:"device_name" form:"device_name"`
	DeviceID     uint64 `gorm:"device_id;bigint(64);not null" json:"device_id" form:"device_id"`
	ItemID       uint64 `gorm:"item_id;bigint(64);key;not null" json:"item_id" form:"item_id"`
	TransferType int64  `gorm:"transfer_type;int(32);not null" json:"transfer_type" form:"transfer_type"`
	Content      string `gorm:"content;text;not null" json:"content" form:"content"`
	CreatedAt    int64  `gorm:"created_at;bigint(64)" json:"created_at" form:"created_at"`
	UpdatedAt    int64  `gorm:"updated_at;bigint(64)" json:"updated_at" form:"updated_at"`
	DeletedAt    int64  `gorm:"deleted_at;bigint(64)" json:"deleted_at" form:"deleted_at"`
}

func (t *CacheItem) TableName() string {
	return "transfer_item"
}

func (t *CacheItem) GetDeviceCaches(deviceName string, deviceID uint64) ([]*server_entity.DeviceCacheItem, error) {
	var list []*server_entity.DeviceCacheItem
	model := mdb.GetDBInstance().Model(&CacheItem{}).Select("*").
		Where("device_name = ? and device_id = ?", deviceName, deviceID).Order("updated_at desc")
	err := model.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return make([]*server_entity.DeviceCacheItem, 0), nil
	}
	if err != nil {
		return make([]*server_entity.DeviceCacheItem, 0), errno.WrapCodeError(errno.DatabaseError, err)
	}
	return list, nil
}

func (t *CacheItem) InsertCaches(deviceName string, deviceID uint64,
	caches []*server_entity.DeviceCacheItem) error {
	list := make([]CacheItem, 0)
	for _, item := range caches {
		list = append(list, CacheItem{
			DeviceName:   deviceName,
			DeviceID:     deviceID,
			ItemID:       item.ItemID,
			TransferType: item.TransferType,
			Content:      item.Content,
			CreatedAt:    item.TimeStamp,
			UpdatedAt:    item.TimeStamp,
		})
	}
	err := mdb.GetDBInstance().Model(&CacheItem{}).Create(&list).Error
	return err
}

func (t *CacheItem) UpdateCache(deviceName string, deviceID uint64,
	cache *server_entity.DeviceCacheItem) error {
	if cache == nil {
		return errno.WrapCodeErrorf(errno.ParamsError, "param is nil")
	}
	item := CacheItem{
		ItemID:       cache.ItemID,
		TransferType: cache.TransferType,
		Content:      cache.Content,
		CreatedAt:    cache.TimeStamp,
		UpdatedAt:    cache.TimeStamp,
	}
	model := mdb.GetDBInstance().Model(&CacheItem{}).
		Select("item_id, transfer_type, content, created_at, updated_at").
		Where("device_name = ? and device_id = ? and item_id = ?",
			deviceName, deviceID, cache.ItemID)
	err := model.Updates(&item).Error
	return err
}

func (t *CacheItem) DeleteCache(deviceName string, deviceID uint64,
	itemID uint64) error {
	err := mdb.GetDBInstance().Model(&CacheItem{}).
		Delete(&CacheItem{}, "device_name = ? and device_id = ? and item_id = ?",
			deviceName, deviceID, itemID,
		).Error
	return err
}
