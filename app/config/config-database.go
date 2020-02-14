package config

type cnf struct {
	Host     string
	Port     uint32
	DBName   string
	User     string
	Password string
}

//Config Return
var Config = cnf{
	Host:     "localhost",
	Port:     5432,
	DBName:   "postgres",
	User:     "postgres",
	Password: "P@ssw0rd",
}
