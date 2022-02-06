package infrastructure

type Config struct {
	DB struct {
		Production struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
		Test struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {

	c := new(Config)
	
	c.DB.Production.Host = "mysql-0.mysql"
	c.DB.Production.Username = "go_user"
	c.DB.Production.Password = "password"
	c.DB.Production.DBName = "go_database"

	c.DB.Test.Host = "mysql"
	c.DB.Test.Username = "go_user"
	c.DB.Test.Password = "password"
	c.DB.Test.DBName = "go_database"

	c.Routing.Port = ":8080"

	return c
}
