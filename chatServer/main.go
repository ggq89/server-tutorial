package main

import (
	"flag"
	"gohipernetFake"
)

func main() {
	gohipernetFake.NetLibInitLog()

	netConfig, appConfig := parseAppConfig()
	netConfig.WriteNetworkConfig(true)

	// createAnsStartServer(netConfig, appConfig)
}

func parseAppConfig() (gohipernetFake.NetworkConfig, configAppServer) {

	gohipernetFake.NTELIB_LOG_INFO("[[Setting NetworkConfig]]")

	appConfig := configAppServer {
		GameName: "chatServer",
		RoomMaxCount: 100,
		RoomStartNum: 0,
		RoomMaxUserCount: 4,
	}

	netConfig := gohipernetFake.NetworkConfig{}

	flag.BoolVar(&netConfig.IsTcp4Addr, "c_isTcp4Addr", true, "bool flag")
	flag.StringVar(&netConfig.BindAddress, "c_BindAddress", "127.0.0.0:11021", "string flag")
	flag.IntVar(&netConfig.MaxSessionCount, "c_MaxSessionCount", 0, "int flag")
	flag.IntVar(&netConfig.MaxPacketSize, "c_MaxPacketSize", 0, "int flag")

	return netConfig, appConfig
}