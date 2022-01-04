package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "github.com/essayZW/microdemo/content/proto"
)

// Content 内容服务
type Content struct {
	contents []*ContentInfo
}

// Query 通过ID查询内容
func (content *Content) Query(ctx context.Context, id *pb.ContentId, resp *pb.ContentRep) error {
	if id.Id < 0 {
		return errors.New("Content id must greater than 0 or equal 0")
	}
	if int(id.Id) >= len(content.contents) {
		return fmt.Errorf("Content %d does not exists", id.Id)
	}
	log.Printf("Query content id %d", id.Id)
	c := content.contents[id.Id]
	resp.Id = c.ID
	resp.Name = c.Name
	resp.Description = c.Discription
	resp.Userid = c.Userid
	return nil
}

// ContentInfo 内容信息
type ContentInfo struct {
	ID          int32
	Name        string
	Discription string
	Userid      int32
}

// New 创建Content service
func New(contents []*ContentInfo) *Content {
	return &Content{
		contents: contents,
	}
}
