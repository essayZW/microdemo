package main

import (
	"log"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	content "github.com/essayZW/microdemo/content/proto"
	"github.com/essayZW/microdemo/content/service"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {

	etcdRegistry := etcd.NewRegistry(
		registry.Addrs("172.17.0.2:2379"),
		etcd.Auth("root", "beihai"),
	)

	srv := micro.NewService(
		micro.Name("content"),
		micro.Registry(etcdRegistry),
	)

	contentService := service.New([]*service.ContentInfo{
		{
			ID:          1,
			Name:        "test content",
			Discription: "content for test",
			Userid:      1,
		},
		{
			ID:          2,
			Name:        "essay's content",
			Discription: "content published by essay",
			Userid:      1,
		},
	})
	content.RegisterContentHandler(srv.Server(), contentService)

	srv.Init()
	if err := srv.Run(); err != nil {
		log.Panic(err)
	}
}
