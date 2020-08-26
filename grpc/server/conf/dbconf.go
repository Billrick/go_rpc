package conf

const (
	HOST = "127.0.0.1"
	PORT = 3306
	PARAMS = "parseTime=true"
	USERNAME = "root"
	PASSWORD = "123456"
	DBNAME = "test"
)

type DbConf struct {
	Host string
	Port int64
	USERNAME string
	PASS string
	DbName string
	Params string
}