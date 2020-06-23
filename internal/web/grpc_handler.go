package web

import (
	"context"
	"time"

	web_v1 "github.com/betterchen/go-project-tmpl/api/proto/demo/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// DemoServer ...
type DemoServer struct {
	// ORM...
}

// GetFoo ...
func (s DemoServer) GetFoo(ctx context.Context, req *web_v1.GetFooReq) (resp *web_v1.GetFooResp, err error) {
	resp = &web_v1.GetFooResp{}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	cli := web_v1.NewDemoClient(conn)

	return cli.GetBar(ctx, &web_v1.GetBarReq{})
}

// GetBar ...
func (s DemoServer) GetBar(ctx context.Context, req *web_v1.GetBarReq) (resp *web_v1.GetFooResp, err error) {
	resp = &web_v1.GetFooResp{
		Id: 001,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
		},
		Data: "你好",
	}

	return
}
