package main

import (
	"log"

	"github.com/asim/go-micro/plugins/config/source/etcd/v4"
	"go-micro.dev/v4/config"
)

func main() {
	etcdSource := etcd.NewSource(
		etcd.WithAddress("172.17.0.2:2379"),
		etcd.Auth("root", "beihai"),
		etcd.WithPrefix("/micro/config"),
		etcd.StripPrefix(true),
	)
	conf, err := config.NewConfig(
		config.WithSource(etcdSource),
	)
	if err != nil {
		log.Panicln(err)
	}
	// Watch 不需要前缀
	// 监听的是Watcher下悬挂的节点，比如/watcher/xss
	go func() {
		watcher, err := conf.Watch("watcher")
		if err != nil {
			log.Panic(err)
		}
		for {
			value, err := watcher.Next()
			if err != nil {
				log.Println(err)
			}
			log.Printf("Watcher: %s", string(value.Bytes()))
		}
	}()
}
