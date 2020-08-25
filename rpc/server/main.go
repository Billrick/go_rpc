package main

import (
	"log"
	"net/http"
	"net/rpc"
	"strconv"
)

func main(){
	//注册向外提供服务的结构体, 在server.register中 suitableMethods(s.typ, true) 方法会获取所有方法
	//通过反射获取方法及参数信息
	rpc.Register(new(Student))
	//把服务处理绑定到http协议上
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Student struct {
	Stus []map[string]interface{}
}

type Request struct {
	Size int
}
type Response struct {
	Student []map[string]interface{}
}

func (s *Student)GetStudents(r Request, res *Response) error{
	//初始化数据列表
	s.Stus = make([]map[string]interface{},r.Size,r.Size)
	for i := 0; i < r.Size; i++ {
		tmp := make(map[string]interface{},3)
		tmp["name"] = "stu" + strconv.Itoa(i)
		tmp["age"] = i
		tmp["married"] = (i % 2 == 0)
		s.Stus[i] = tmp
	}
	res.Student = s.Stus
	return nil
}


