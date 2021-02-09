package main

import (
	"go.uber.org/zap"
	"gohipernetFake"
	"strconv"
	"strings"
)

type configAppServer struct {
	GameName			string

	RoomMaxCount 		int32
	RoomStartNum 		int32
	RoomMaxUserCount 	int32
}

type ChatServer struct {
	ServerIndex int
	IP			string
	Port 		int

}

func createAnsStartServer(netConfig gohipernetFake.NetworkConfig, appConfig configAppServer) {
	gohipernetFake.NTELIB_LOG_INFO("CreateServer !!!")

	var server ChatServer

	if server.setIPAddress(netConfig.BindAddress) == false {
		gohipernetFake.NTELIB_LOG_ERROR("server address parsing failed")
		return
	}

	
}

func (server *ChatServer) setIPAddress(ipAddress string) bool {
	results := strings.Split(ipAddress, ":")
	if len(results) != 2 {
		return false
	}

	server.IP = results[0]
	server.Port, _ = strconv.Atoi(results[1])

	gohipernetFake.NTELIB_LOG_INFO("Server Address", zap.String("IP", server.IP), zap.Int("Port", server.Port))
}