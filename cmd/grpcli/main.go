package main

import (
	"context"
	"flag"
	"log"
	"time"

	web_v1 "github.com/betterchen/go-project-tmpl/api/proto/demo/v1"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", "", "assign grpc server's address")

func main() {
	flag.Parse()

	testGetFoo()
}

func testGetFoo() {
	// 建立连接
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	// 创建gRPC客户端
	cli := web_v1.NewDemoClient(conn)

	resp, err := cli.GetFoo(context.Background(), &web_v1.GetFooReq{})
	if err != nil {
		log.Panic(err)
	}

	log.Printf(
		"\nresp:\nid %d\ncreated_at: %v\ndata: %s",
		resp.GetId(),
		time.Unix(resp.GetCreatedAt().GetSeconds(), 0),
		resp.GetData(),
	)
}
