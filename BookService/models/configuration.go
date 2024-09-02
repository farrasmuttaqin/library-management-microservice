package models

type ApplicationConfiguration struct {
	Name           string
	Port           int
	ServerTimeHour int
	ClientTimeHour int
	TimeZone       string
}

type DatabaseMySQLConfiguration struct {
	Driver   string
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

type RedisConfiguration struct {
	Host      string
	Port      int
	Password  string
	TTLSecond int
}

type GRPCConfiguration struct {
	UserServiceGRPCAddress string
}
