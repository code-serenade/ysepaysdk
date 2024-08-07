package main

import (
	"fmt"
	"log"

	"github.com/Centny/rediscache"
	ysepaysdk "github.com/code-serenade/ysepaysdk/api"
	"github.com/codingeasygo/util/xmap"
	"github.com/codingeasygo/web"
	"github.com/wfunc/go/session"
)

func main() {
	// 初始化配置
	{

		ysepaysdk.FindPubKey = findPubKey

		// 订单回调
		ysepaysdk.NotifyPayDone = func(orderID string, result xmap.M) (err error) {
			err = donePay(orderID, result)
			return
		}

		// 商户开户/支付业务回调
		ysepaysdk.NotifyCustDone = func(orderID string, result xmap.M) (err error) {
			err = doneCust(orderID, result)
			return
		}
	}
	rediscache.InitRedisPool("redis.loc:6379?db=3")
	sb := session.NewDbSessionBuilder()
	sb.Redis = rediscache.C
	sb.ShowLog = false
	web.Shared.Builder = sb
	ysepaysdk.Handle("", web.Shared)
	go web.HandleSignal()
	port := ":8080"
	log.Printf("start fp service on %v", port)

	var err = web.ListenAndServe(port)
	if err != nil {
		panic(err)
	}
}

func findPubKey(orderID string) (key []byte, err error) {
	// 业务逻辑
	err = fmt.Errorf("findPubKey not implemented")
	return
}

func donePay(orderID string, result xmap.M) (err error) {
	// 业务逻辑
	log.Printf("orderID: %v, result: %v", orderID, result)
	return
}

func doneCust(orderID string, result xmap.M) (err error) {
	// 业务逻辑
	log.Printf("orderID: %v, result: %v", orderID, result)
	return
}
