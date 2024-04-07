package constant

import "time"

var (
	// allow change in test
	IdleTimeOut = 10 * time.Second

	Version     = "1.1.4-dev"
	GitBranch   string
	GitRevision string
	BuildTime   string
	StartTime   = time.Now().Local()
)

const (
	DialTimeOut = 3 * time.Second

	SmuxGCDuration       = 30 * time.Second
	SmuxMaxAliveDuration = 10 * time.Minute
	SmuxMaxStreamCnt     = 5

	// todo add udp buffer size
	BUFFER_POOL_SIZE = 1024      // support 512 connections
	BUFFER_SIZE      = 20 * 1024 // 20KB the maximum packet size of shadowsocks is about 16 KiB
)

// relay type
const (
	RelayTypeRaw  = "raw"
	RelayTypeWS   = "ws"
	RelayTypeWSS  = "wss"
	RelayTypeMWSS = "mwss"
	RelayTypeMTCP = "mtcp"
)
