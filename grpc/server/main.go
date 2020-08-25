package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	student "z.cn/20200825-rpc/grpc/proto"
)

const (
	port = ":8080"
)

type server struct {
	student.UnimplementedGoSutdentServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetSutdents(ctx context.Context, in *student.SendParam) (*student.GetResponse, error) {
	students := make([]map[string]interface{},0)
	for i := 0; i < 5; i++ {
		tmp := make(map[string]interface{},3)
		tmp["name"] = "stu" + strconv.Itoa(i)
		tmp["age"] = i
		tmp["married"] = (i % 2 == 0)
		students = append(students,tmp)
	}
	data,err := json.Marshal(students)
	if err != nil {
		return &student.GetResponse{HttpCode: 500, Response: err.Error()},err
	}
	return &student.GetResponse{HttpCode: 200,Response:string(data)},nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//注册服务并保留方法
	s := grpc.NewServer()
	student.RegisterGoSutdentServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}