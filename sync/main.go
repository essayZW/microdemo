package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/asim/go-micro/plugins/sync/etcd/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/sync"
)

func main() {
	var outputer bool
	flag.BoolVar(&outputer, "o", false, "Is outputer")
	flag.Parse()
	logger.Info(outputer)

	etcdSync := etcd.NewSync(
		&etcd.Auth{
			Username: "root",
			Password: "beihai",
		},
		sync.Nodes("172.17.0.2:2379"),
	)
	etcdSync.Init()
	if outputer {
		i := 0
		for {
			etcdSync.Lock("output")
			logger.Info(i)
			time.Sleep(time.Second)
			i++
			etcdSync.Unlock("output")
		}
	} else {
		var num int
		lock := false
		for {
			_, err := fmt.Scanf("%d", &num)
			if err != nil {
				break
			}
			lock = !lock
			if lock {
				etcdSync.Lock("output")
				logger.Info("Lock")
			} else {
				etcdSync.Unlock("output")
				logger.Info("UnLock")
			}
		}
	}

}
