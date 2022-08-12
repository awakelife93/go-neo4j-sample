package lib

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/thoas/go-funk"
)

var driverObject neo4j.Driver = nil

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
		password = "test"
	}

	driver, error := neo4j.NewDriver(
		uri,
		neo4j.BasicAuth(username, password, ""),
		func(config *neo4j.Config) {
			config.MaxConnectionLifetime = 60 * 60 * time.Second
			config.MaxConnectionPoolSize = 50
			config.ConnectionAcquisitionTimeout = 2 * time.Minute
			config.Encrypted = false
		})

	if error != nil {
		return error
	}

	driverObject = driver

	return nil
}

func setupSession() (neo4j.Session, error) {
	if funk.IsEmpty(driverObject) {
		return nil, errors.New("Empty driverObject")
	}

	session, error := driverObject.NewSession(
		neo4j.SessionConfig{
			AccessMode: neo4j.AccessModeWrite,
		})

	if error != nil {
		return session, error
	}

	return session, error
}

func Initialize() (string, error) {
	fmt.Println("Neo4j initialize")

	connectError := connect()

	if connectError != nil {
		return "", connectError
	}

	return "initialize success", nil
}

func getSession() neo4j.Session {
	session, error := setupSession()

	if error != nil {
		fmt.Println("getSession Error ====> ", error.Error())
	}

	return session
}

func clearDriver() {
	fmt.Println("Clear Neo4j Driver")

	if driverObject != nil {
		driverObject.Close()
		driverObject = nil
	}
}

func clearSession(session neo4j.Session) {
	fmt.Println("Clear Neo4j Session")

	if session != nil {
		error := session.Close()

		if error != nil {
			fmt.Println("clearSession Error ====> ", error)
		}

		session = nil
	}
}

func renewalSession(session neo4j.Session) {
	defer setupSession()
	clearSession(session)
}

func Clear() {
	clearDriver()
}
