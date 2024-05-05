package server

type TransferItem struct {
	ID           uint64 `gorm:"id;bigint(64);primary_key;AUTO_INCREMENT" json:"id" form:"id"`
	DeviceName   string `gorm:"device_name;varchar(256);not null" json:"device_name" form:"device_name"`
	DeviceID     uint64 `gorm:"device_id;bigint(64);not null" json:"device_id" form:"device_id"`
	TransferType int64  `gorm:"transfer_type;int(32);not null" json:"transfer_type" form:"transfer_type"`
	Content      string `gorm:"content;text;not null" json:"content" form:"content"`
	Transferred  int64  `gorm:"transferred;int(32);not null" json:"transferred" form:"transferred"`
	CreatedAt    int64  `gorm:"created_at;bigint(64)" json:"created_at" form:"created_at"`
	UpdatedAt    int64  `gorm:"updated_at;bigint(64)" json:"updated_at" form:"updated_at"`
	DeletedAt    int64  `gorm:"deleted_at;bigint(64)" json:"deleted_at" form:"deleted_at"`
}
