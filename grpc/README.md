## GRPC
### 1.安装 golang 的proto工具包
```bash 
go get -u github.com/golang/protobuf/proto
```
### 2.安装 goalng 的proto编译支持
```bash 
go get -u github.com/golang/protobuf/protoc-gen-go 
```

### 3.生成go文件
```bash 
                           输出的目录  proto所在目录
cmd:protoc --go_out=plugins=grpc:./ ./spider.proto
```