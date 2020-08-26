package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	service "z.cn/20200825-rpc/grpc/service"
)

const (
	address     = "localhost:80"      //ip + port 服务地址
	version = "v1.0"
)

func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := service.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stu := service.Student{Name: "zhangsan",Age: 18}
	response, err := c.Create(ctx, &service.CreateRequest{Version: version,Stu: &stu})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("create: %s", response)

	readResponse,err := c.Read(ctx,&service.ReadRequest{Version: version})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("read: %s", readResponse)
}
