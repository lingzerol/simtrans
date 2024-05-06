package server

import (
	server_entity "github.com/lingzerol/simtrans/model/entity/server"
)

type Connection interface {
	GetConnectionID() uint64
	Command(*server_entity.ServerCommand) error
	Copy(*server_entity.ServerCommand) error
	Put(*server_entity.ServerCommand) error
	Paste(*server_entity.ServerCommand) error
	DefaultPaste() error
	CheckAuth() bool
	HearBeat() error
	Alive() error
}
