package main

import (
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"log"
	"net"
	"z.cn/20200825-rpc/grpc/server/conf"
	"z.cn/20200825-rpc/grpc/server/serviceimpl"
	"z.cn/20200825-rpc/grpc/service"
)

func main(){
	dsn := getCmdAgrs()
	db,err := sqlx.Connect("mysql",dsn)
	if err != nil{
		log.Fatalf("connect db failed ,err",err)
		return
	}
	err = db.Ping()
	if err!= nil {
		log.Fatalf("connect db ping() failed ,err",err)
	}
	log.Println("connect db success !")
	ls , err := net.Listen("tcp",":80")
	if err != nil{
		log.Fatalf("listen :80 failed ,err",err)
		return
	}
	server := grpc.NewServer()
	studentService := serviceimpl.NewStudentServiceServer(db)
	//注册服务
	service.RegisterStudentServiceServer(server,studentService)


	server.Serve(ls)
	defer db.Close()
}

func getCmdAgrs()(dsn string){
	c := conf.DbConf{}
	flag.StringVar(&c.Host,"host",conf.HOST,"db host")
	flag.Int64Var(&c.Port,"port",int64(conf.PORT),"db host")
	flag.StringVar(&c.USERNAME,"username",conf.USERNAME,"db password")
	flag.StringVar(&c.PASS,"pass",conf.PASSWORD,"db password")
	flag.StringVar(&c.Params,"params",conf.PARAMS,"db host")
	flag.StringVar(&c.DbName,"dbname",conf.DBNAME,"db host")
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",c.USERNAME,c.PASS,c.Host,c.Port,c.DbName,c.Params)
	return
}