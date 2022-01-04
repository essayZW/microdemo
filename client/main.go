package main

import (
	"context"
	"log"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	pb "github.com/essayZW/microdemo/content/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs("172.17.0.2:2379"),
		etcd.Auth("root", "beihai"),
	)

	src := micro.NewService(
		micro.Registry(etcdRegistry),
	)

	client := pb.NewContentService("content", src.Client())

	for i := 0; i < 10; i++ {
		resp, err := client.Query(context.TODO(), &pb.ContentId{
			Id: 1,
		})
		if err != nil {
			log.Panicln(err)
		}
		log.Println(resp)
	}
}
