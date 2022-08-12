package main

import (
	"fmt"

	"github.com/awakelife93/go-neo4j-sample/lib"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func start() {
	initializeResult, initializeError := lib.Initialize()

	if initializeError != nil {
		fmt.Println("Initialize Error ====>", initializeError.Error())
		lib.Clear()
		return
	}

	fmt.Println(initializeResult)

	createQueryResult, createQueryError := lib.CreateQuery(
		"CREATE (a:Sample) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
		map[string]interface{}{"message": "hello, world"},
	)

	if createQueryError != nil {
		fmt.Println("Create Query Error ====>", createQueryError.Error())
		return
	}

	fmt.Println("Create Result ====>", createQueryResult)

	matchQueryResult, matchQueryError := lib.MatchQuery(
		"MATCH (n) RETURN n",
		map[string]interface{}{},
	)

	if matchQueryError != nil {
		fmt.Println("Match Query Error ====>", matchQueryError.Error())
		return
	}

	if matchQueryResult != nil {
		var result = matchQueryResult.(neo4j.Node)

		fmt.Println("Match Result Id ====>", result.Id())
		fmt.Println("Match Result Label ====>", result.Labels())
		fmt.Println("Match Result Props ====>", result.Props())
	}
}

func main() {
	start()
}
