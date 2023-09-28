package config

type Config struct {
	Http     *Http
	DataBase *DataBase
}

type Http struct {
	Port string
	Host string
}

type DataBase struct {
	Port     string
	Host     string
	User     string
	Password string
	Ssl      string
	Driver   string
	DSN      string
}
