package server

import "github.com/lingzerol/simtrans/model/entity"

type ServerCommand struct {
	entity.Command
}

const (
	AuthCommand             = "auth"
	CopyCommand             = "copy"
	PutCommand              = "put"
	PasteCommand            = "paste"
	DeleteCommand           = "delete"
	RefreshCacheCommand     = "refresh_cache"
	SendRefreshCacheCommand = "send_refresh_cache"
)
