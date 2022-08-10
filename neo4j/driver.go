package neo4j

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/thoas/go-funk"
)

var driverObject neo4j.Driver = nil
var sessionObject neo4j.Session = nil

func connect() error {
	uri := os.Getenv("uri")
	username := os.Getenv("username")
	password := os.Getenv("password")

	if funk.IsEmpty(uri) {
		uri = "bolt://localhost:7687"
	}

	if funk.IsEmpty(username) {
		username = "neo4j"
	}

	if funk.IsEmpty(password) {
		password = "neo4j"
	}

	driver, error := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""), func(config *neo4j.Config) {
		config.MaxConnectionLifetime = 60 * 60 * time.Second
		config.MaxConnectionPoolSize = 50
		config.ConnectionAcquisitionTimeout = 2 * time.Minute
	})

	if error != nil {
		return error
	}

	driverObject = driver

	return nil
}

func setupSession() error {
	if funk.IsEmpty(driverObject) {
		return errors.New("Empty driverObject")
	}

	session, error := driverObject.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite, DatabaseName: "neo4j"})

	if error != nil {
		return error
	}

	sessionObject = session

	return nil
}

func Initialize() (string, error) {
	fmt.Println("Neo4j initialize")

	connectError := connect()

	if connectError != nil {
		return "", connectError
	}

	setupSessionError := setupSession()

	if setupSessionError != nil {
		return "", setupSessionError
	}

	return "initialize success", nil
}

func getSession() neo4j.Session {
	return sessionObject
}

func ClearDriver() {
	fmt.Println("Clear Neo4j Driver")

	if driverObject != nil {
		driverObject.Close()
		driverObject = nil
	}
}

func ClearSession() {
	fmt.Println("Clear Neo4j Session")

	if sessionObject != nil {
		sessionObject.Close()
		sessionObject = nil
	}
}
