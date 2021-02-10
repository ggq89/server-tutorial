package main

import (
	"go.uber.org/zap"
	"gohipernetFake"
	"strconv"
	"strings"
	"time"
)

type EchoServer struct {
	ServerIndex int
	IP 			string
	Port 		int
}

func createServer(netConfig gohipernetFake.NetworkConfig) {
	gohipernetFake.NTELIB_LOG_INFO("Create Server !!!")

	var server EchoServer

	if server.setIPAddress(netConfig.BindAddress) == false {
		gohipernetFake.NTELIB_LOG_ERROR("fail. server address")
		return
	}

	networkFunctor := gohipernetFake.SessionNetworkFunctors{}

	networkFunctor.OnConnect = server.OnConnect
	networkFunctor.OnReceive = server.OnReceive
	networkFunctor.OnReceiveBufferedData = nil
	networkFunctor.OnClose = server.OnClose
	networkFunctor.PacketTotalSizeFunc = gohipernetFake.PacketTotalSize
	networkFunctor.PacketHeaderSize = gohipernetFake.PACKET_HEADER_SIZE
	networkFunctor.IsClientSession = true

	gohipernetFake.NetLibStartNetwork(&netConfig, networkFunctor)

	server.Stop()
}

func (server *EchoServer) setIPAddress(ipAddress string) bool {
	results := strings.Split(ipAddress, ":")
	if len(results) != 2 {
		return false
	}

	server.IP = results[0]
	server.Port, _ = strconv.Atoi(results[1])

	gohipernetFake.NTELIB_LOG_INFO(
		"Server Address", zap.String("IP", server.IP),
		zap.Int("PORT", server.Port))
	return true
}

func (server *EchoServer) OnConnect(sessionIndex int32, sessionUniqueID uint64) {
	gohipernetFake.NTELIB_LOG_INFO("client OnConnect",
		zap.Int32("sessionIndex", sessionIndex),
		zap.Uint64("sessionUniqueID", sessionUniqueID))
}

func (server *EchoServer) OnReceive(sessionIndex int32, sessionUniqueID uint64, data []byte) bool {
	gohipernetFake.NTELIB_LOG_INFO("OnReceive",
		zap.Int32("sessionIndex", sessionIndex),
		zap.Uint64("sessionUniqueID", sessionUniqueID),
		zap.Int("packetSize", len(data)))

	gohipernetFake.NetLibISendToClient(sessionIndex, sessionUniqueID, data)
	return true
}

func (server *EchoServer) OnClose(sessionIndex int32, sessionUniqueID uint64) {
	gohipernetFake.NTELIB_LOG_INFO("client OnCloseClientSession",
		zap.Int32("sessionIndex", sessionIndex),
		zap.Uint64("sessionUniqueID", sessionUniqueID))
}

func (server *EchoServer) Stop() {
	gohipernetFake.NTELIB_LOG_INFO("chatServer Stop !!!")

	gohipernetFake.NetLib_StopServer()

	gohipernetFake.NTELIB_LOG_INFO("chatServer Stop Waiting...")
	time.Sleep(1 * time.Second)
}