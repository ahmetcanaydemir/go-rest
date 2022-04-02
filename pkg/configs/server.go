package configs

type server struct {
	Config struct {
		DbConnectionString string
		Port               string
	}
}

var Server server
