package cron

import (
	"time"

	"github.com/freedomkk-qfeng/windows-agent/g"
	"github.com/open-falcon/common/model"
)

func SyncTrustableIps() {
	if g.Config().Heartbeat.Enabled && g.Config().Heartbeat.Addr != "" {
		go syncTrustableIps()
	}
}

func syncTrustableIps() {

	duration := time.Duration(g.Config().Heartbeat.Interval) * time.Second

	for {
		time.Sleep(duration)

		var ips string
		err := g.HbsClient.Call("Agent.GetWhiteIPList", model.NullRpcRequest{}, &ips)
		if err != nil {
			g.Logger().Println("ERROR: call Agent.GetWhiteIPList fail", err)
			continue
		}

		g.SetTrustableIps(ips)
	}
}
