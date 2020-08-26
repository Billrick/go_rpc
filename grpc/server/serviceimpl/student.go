package serviceimpl

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"z.cn/20200825-rpc/grpc/service"
)

const (
	port = ":8080"
	version = "v1.0"
)
//实现接口的结构体
type StudentServiceServer struct {
	Db *sqlx.DB
}

func NewStudentServiceServer(db *sqlx.DB) *StudentServiceServer {
	return &StudentServiceServer{
		Db: db,
	}
}

func (s *StudentServiceServer)checkVersion(inversion string) error{
	if  inversion != version {
		return fmt.Errorf("wrong api version")
	}
	if len(inversion) < 1{
		return fmt.Errorf("invalid api version")
	}
	return nil
}

func (s *StudentServiceServer)Create(ctx context.Context,req *service.CreateRequest) (res *service.CreateResponse,err error){
	if err := s.checkVersion(req.Version); err!=nil{
		return nil,status.Error(codes.Unknown,err.Error())
	}
	conn,err := s.Db.Conn(ctx)
	if err != nil {
		return nil,status.Errorf(codes.Unknown,"connect db failed ,err %s",err.Error())
	}
	defer conn.Close()
	sqlStr := "insert into student(name,age) values(?,?)"
	result, err := conn.ExecContext(ctx, sqlStr, req.Stu.Name, req.Stu.Age)
	if err != nil {
		return nil,status.Errorf(codes.Unknown,"insert student failed ,err %s",err.Error())
	}
	id,err := result.LastInsertId()
	if err != nil{
		return nil,status.Errorf(codes.Unknown,"insert student get id failed ,err %s",err.Error())
	}
	return &service.CreateResponse{Version: version,Id: id},nil
}
func (s *StudentServiceServer)Update(ctx context.Context,req *service.UpdateRequest) (res *service.UpdateResponse,err error){
	if err := s.checkVersion(req.Version); err!=nil{
		return nil,status.Error(codes.Unknown,err.Error())
	}
	conn,err := s.Db.Conn(ctx)
	if err != nil {
		return nil,status.Errorf(codes.Unknown,"connect db failed ,err %s",err.Error())
	}
	defer conn.Close()
	sqlStr := "update student set name = ?, age = ? where id = ?"
	result , err := conn.ExecContext(ctx,sqlStr,req.Stu.Name,req.Stu.Age,req.Stu.Id)
	if err != nil {
		return nil,status.Errorf(codes.Unknown,"update student failed ,err %s",err.Error())
	}
	rows,err := result.RowsAffected()
	if err != nil{
		return nil,status.Errorf(codes.Unknown,"update student get rowsaffected failed ,err %s",err.Error())
	}
	return &service.UpdateResponse{Version: version,Rows: rows},nil
}
func (s *StudentServiceServer)Remove(ctx context.Context,req *service.RemoveRequest) (res *service.RemoveResponse,err error){
	if err := s.checkVersion(req.Version); err!=nil{
		return nil,status.Error(codes.Unknown,err.Error())
	}
	conn,err := s.Db.Conn(ctx)
	if err != nil {
		return nil,status.Errorf(codes.Unknown,"connect db failed ,err %s",err.Error())
	}
	defer conn.Close()
	sqlStr := "delete from student where id = ?"
	result , err := conn.ExecContext(ctx,sqlStr,req.Id)
	if err != nil {
		return nil,status.Errorf(codes.Unknown,"delete student failed ,err %s",err.Error())
	}
	rows,err := result.RowsAffected()
	if err != nil{
		return nil,status.Errorf(codes.Unknown,"delete student get RowsAffected failed ,err %s",err.Error())
	}
	return &service.RemoveResponse{Version: version,Rows: rows},nil
}
func (s *StudentServiceServer)Read(ctx context.Context,req *service.ReadRequest) (res *service.ReadResponse,err error){
	if err := s.checkVersion(req.Version); err!=nil{
		return nil,status.Error(codes.Unknown,err.Error())
	}
	conn,err := s.Db.Conn(ctx)
	if err != nil {
		return nil,status.Errorf(codes.Unknown,"connect db failed ,err %s",err.Error())
	}
	sqlStr := "select id , name , age from student "
	param := make([]interface{},0)
	if req.Id != 0{
		sqlStr = sqlStr + " where id = ?"
		param = append(param,req.Id)
	}
	rows , err := conn.QueryContext(ctx,sqlStr, param...)
	if err != nil {
		return nil,status.Errorf(codes.Unknown,"select student failed ,err %s",err.Error())
	}
	defer conn.Close()
	defer rows.Close()
	students := make([]*service.Student,0)
	for rows.Next() {
		stu := &service.Student{}
		if err:= rows.Scan(&stu.Id,&stu.Name,&stu.Age);err != nil{
			return nil,status.Errorf(codes.Unknown,"scan student rows failed ,err %s",err.Error())
		}
		students = append(students,stu)
	}
	return &service.ReadResponse{Version: version,Students: students},nil
}
