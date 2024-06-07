package iprecoder

import (
	"github.com/xinyangme001/V2bX/limiter"
)

type IpRecorder interface {
	SyncOnlineIp(Ips []limiter.UserIpList) ([]limiter.UserIpList, error)
}
