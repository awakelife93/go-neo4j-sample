package main

import (
	"fmt"

	"github.com/awakelife93/go-neo4j-sample/neo4j"
)

func start() {
	initializeResult, initializeError := neo4j.Initialize()

	if initializeError != nil {
		fmt.Println("Initialize Error ====>", initializeError.Error())
		neo4j.ClearDriver()
		neo4j.ClearSession()
		return
	}

	fmt.Println(initializeResult)

	queryResult, queryError := neo4j.Query(
		"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
		map[string]interface{}{"message": "hello, world"},
	)

	// todo: TLS error: Remote end closed the connection, check that TLS is enabled on the server 트러블 슈팅
	if queryError != nil {
		fmt.Println("Query Error ====>", queryError.Error())
		return
	}

	fmt.Println(queryResult)
}

func main() {
	start()
}
