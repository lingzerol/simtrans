package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/lingzerol/simtrans/library/encrypt"
	"github.com/lingzerol/simtrans/library/errno"
	"github.com/lingzerol/simtrans/library/utils"
	"github.com/lingzerol/simtrans/model/config"
	data "github.com/lingzerol/simtrans/model/data/server"
	"github.com/lingzerol/simtrans/model/entity"
	server_entity "github.com/lingzerol/simtrans/model/entity/server"
	"github.com/lingzerol/simtrans/model/logger"
	"github.com/united-manufacturing-hub/expiremap/pkg/expiremap"
)

type ServerConnection struct {
	ID                  uint64
	DeviceName          string
	ctx                 context.Context
	cancel              context.CancelFunc
	conn                net.Conn
	authMessage         chan []byte
	processedCommandIds *expiremap.ExpireMap[uint64, bool]
}

func NewServerConnection(ctx context.Context, conn net.Conn) *ServerConnection {
	var exMap = expiremap.NewEx[uint64, bool](time.Minute, 2*entity.CommandExpireTime)
	id, _ := utils.RandomID()
	nctx, cancel := context.WithCancel(ctx)
	serverConnection := &ServerConnection{
		ID:                  id,
		ctx:                 nctx,
		cancel:              cancel,
		conn:                conn,
		processedCommandIds: exMap,
	}
	go serverConnection.HearBeat()

}

func (s *ServerConnection) Listen() {
	if s == nil {
		logger.GetLogger().Warn("empty connection")
		return
	}
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				logger.GetLogger().Info(fmt.Sprintf("connection %s done", s.ID))
				s.conn.Close()
				return
			}
		}
	}()

	buf := make([]byte, entity.MaxBufferSize)
	for {
		_, err := s.conn.Read(buf)
		if err != nil {
			logger.GetLogger().Warn(fmt.Sprintf("read message error: ", err))
			s.conn.Close()
			s.cancel()
			return
		}
		secretKey := config.GetSecretKey(s.DeviceName)
		if secretKey == "" {
			s.conn.Close()
			s.cancel()
			return
		}

		var command server_entity.ServerCommand
		err = json.Unmarshal(buf, &command)
		if err != nil {
			logger.GetLogger().Warn(fmt.Sprintf("message parse error: ", err))
			continue
		}
		s.Command(&command)
	}
}

func (s *ServerConnection) Command(command *server_entity.ServerCommand) error {
	if command.TimeStamp+entity.CommandExpireTime <= time.Now().Unix() {
		return errno.WrapCodeErrorf(errno.RequrestTimeOut, "request is timeout")
	}
	_, ok := s.processedCommandIds.Load(command.ID)
	if ok {
		return errno.NewDefaultCodeError(errno.RepeatRequest)
	}
	s.processedCommandIds.Set(command.ID, true)
	switch command.Command.Type {
	case server_entity.CopyCommand:
		s.Copy(command)
	case server_entity.PutCommand:
		s.Put(command)
	case server_entity.PasteCommand:
		s.Paste(command)
	case server_entity.DeleteCommand:
		s.Delete(command)
	case server_entity.RefreshCacheCommand:
		s.RefreshCache(command)
	}
}

func (s *ServerConnection) Copy(command *server_entity.ServerCommand) error {

}

func (s *ServerConnection) Put(command *server_entity.ServerCommand) error {

}

func (s *ServerConnection) Paste(command *server_entity.ServerCommand) error {

}

func (s *ServerConnection) DefaultPaste() error {

}

func (s *ServerConnection) Delete(command *server_entity.ServerCommand) error {

}

func (s *ServerConnection) RefreshCache(command *server_entity.ServerCommand) error {
	var params server_entity.CacheRefreshParams
	err := json.Unmarshal([]byte(command.Content), &params)
	if err != nil {
		return errno.WrapCodeError(errno.ParamsError, err)
	}
	err = data.NewCacheSrv().RefreshCache(&params)
	command.Type = server_entity.SendRefreshCacheCommand
	GetServer().SendOtherConnectionsCommand(s.ID, command)
	return err
}

func (s *ServerConnection) SendRefreshCache(command *server_entity.ServerCommand) error {

}

func (s *ServerConnection) CheckAuth() (bool, error) {
	if s == nil {
		return false, errno.WrapCodeErrorf(errno.AuthFailed, fmt.Sprintf("connection is nil"))
	}
	params := entity.AuthParams{
		Command: entity.Command{
			Type:      server_entity.AuthCommand,
			TimeStamp: time.Now().Unix(),
		},
	}
	data, err := json.Marshal(params)
	if err != nil {
		return false, errno.NewCodeError(errno.ParamsError, "json marshal error", err)
	}
	_, err = s.conn.Write([]byte(data))
	if err != nil {
		return false, errno.NewCodeError(errno.ConnSendFailed, "send check auth failed", err)
	}
	var message []byte
	select {
	case <-time.After(time.Duration(entity.AuthTimeOut)):
	case message = <-s.authMessage:
	}
	var authMessage entity.AuthResonse
	err = json.Unmarshal(message, &authMessage)
	if err != nil {
		return false, errno.NewCodeError(errno.ConnSendFailed, "parse check auth message failed", err)
	}
	s.DeviceName = authMessage.DeviceName
	secretKey := config.GetSecretKey(authMessage.DeviceName)
	if secretKey == "" {
		return false, nil
	}
	decrypt_message, err := encrypt.AesDecrypt(authMessage.Message, secretKey)
	if err != nil {
		return false, err
	}
	if decrypt_message == entity.DefaultAuthMessage {
		return true, nil
	}
	return false, nil
}

func (s *ServerConnection) HearBeat() {
	for {
		ok, err := s.CheckAuth()
		if err != nil {
			logger.GetLogger().Warn(fmt.Sprintf("connecion %d hearbeat failed, error: ", s.ID), err)
			s.cancel()
			return
		}
		if !ok {
			logger.GetLogger().Warn(fmt.Sprintf("connecion %d hearbeat failed", s.ID), err)
			s.cancel()
			return
		}
		time.Sleep(time.Duration(entity.HearBeatTimeSpan))
	}
}
