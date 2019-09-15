package consent_manager

import (
	"fmt"

	couchdb "github.com/leesper/couchdb-golang"
)

const (
	defaultBaseURL = "http://localhost:5984/"
)

// DBConfig is the database config to access couchdb
type DBConfig struct {
	BaseURL  string
	DBName   string
	Username string
	Password string
}

// Connect creates a new connection with couchdb running instance
func Connect(config *DBConfig) *couchdb.Database {
	configuration := config
	if config == nil {
		configuration = &DBConfig{
			BaseURL:  defaultBaseURL,
			DBName:   "db",
			Username: "admin",
			Password: "admin",
		}
	}
	server, err := couchdb.NewServer(configuration.BaseURL)

	if err != nil {
		panic(fmt.Sprintf("Cannot connect on couchdb: %v", err))
	}

	token, err := server.Login(configuration.Username, configuration.Password)
	if err != nil {
		panic(fmt.Sprintf("Could not login on server: %v", err))
	}

	fmt.Print("Logged in.", token)

	database, err := server.Get(configuration.DBName)
	if err != nil {
		panic(fmt.Sprintf("Cannot connect on database: %v", err))
	}
	return database
}
