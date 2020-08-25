package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Request struct {
	Size int
}
type Response struct {
	Student []map[string]interface{}
}

func main(){
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}
	// 调用方法  Student = 向外暴露结构名,GetStudents = 结构体向外暴露的方法名
	tmp := Response{}
	err = conn.Call("Student.GetStudents", Request{10},&tmp)
	if err != nil {
		log.Fatalf("call Student.GetStudents remote method fialed ,err :",err)
		return
	}
	fmt.Printf("get students success , repsonse data : %#v\n",tmp)
}