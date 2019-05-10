package conf

// Db настройки подключения к базе
var Db = db{
	"DB_NAME",
	"PASS",
	"127.0.0.1",
	3306,
	"USER_NAME",
}

type db struct {
	User string
	Pass string
	Host string
	Port int
	Name string
}
