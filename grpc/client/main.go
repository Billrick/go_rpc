package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	sutdent "z.cn/20200825-rpc/grpc/proto"
)

const (
	address     = "localhost:8080"      //ip + port 服务地址
	defaultName = "world"
)

func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := sutdent.NewGoSutdentClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetSutdents(ctx, &sutdent.SendParam{Address: address,Method: ""})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("getstudents: %s", r.GetResponse())
}
