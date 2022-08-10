package neo4j

import (
	"errors"
	"fmt"
	"os"

	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/thoas/go-funk"
)

var driverObject neo4j.Driver = nil
var sessionObject neo4j.Session = nil

func connect() error {
	uri := os.Getenv("uri")
	username := os.Getenv("username")
	password := os.Getenv("password")

	if funk.IsEmpty(uri) || funk.IsEmpty(uri) || funk.IsEmpty(uri) {
		return errors.New("Empty Connect Info")
	}

	driver, error := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))

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

	session, error := driverObject.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	if error != nil {
		return error
	}

	sessionObject = session

	return nil
}

func Initialize() error {
	fmt.Println("Neo4j initialize")

	connectError := connect()

	if connectError != nil {
		return connectError
	}

	setupSessionError := setupSession()

	if setupSessionError != nil {
		return setupSessionError
	}

	return nil
}

func getSession() neo4j.Session {
	return sessionObject
}

func ClearDriver() {
	fmt.Println("Clear Neo4j Driver")

	if !funk.IsEmpty(driverObject) {
		driverObject.Close()
		driverObject = nil
	}
}

func ClearSession() {
	fmt.Println("Clear Neo4j Session")

	if !funk.IsEmpty(sessionObject) {
		sessionObject.Close()
		sessionObject = nil
	}
}
