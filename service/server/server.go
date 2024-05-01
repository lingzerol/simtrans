package server

import (
	"bufio"
	"fmt"
	"net"

	"github.com/lingzerol/simtrans/library/config"
	"github.com/lingzerol/simtrans/library/logger"
)

func InitServer(configPath string) {
	serverConfig, err := config.InitServerConfig(configPath)
	if err != nil {
		panic("[server] init failed: " + err.Error())
	}
	if serverConfig == nil {
		panic("[server] init failed with unknown error")
	}

	// 监听TCP连接
	listener, err := net.Listen("tcp", serverConfig.Listen)
	if err != nil {
		panic("[server]: listen port failed" + err.Error())
	}
	defer listener.Close()
	logger.GetLogger().Info("[server] listen " + serverConfig.Listen)

	for {
		conn, err := listener.Accept()
		if err != nil || conn == nil {
			logger.GetLogger().Warn("[server] connect client "+conn.RemoteAddr().Network()+" failed:", err)
			continue
		}
		logger.GetLogger().Info("[server] client establish connection success:", conn.RemoteAddr().Network())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 读取客户端发送的消息
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// 打印收到的消息
		fmt.Println("收到客户端消息:", scanner.Text())

		// 回复客户端
		response := "收到您的消息: " + scanner.Text()
		conn.Write([]byte(response + "\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取客户端消息时发生错误:", err)
	}
}
