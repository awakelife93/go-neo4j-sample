package main

import (
	"fmt"

	"github.com/awakelife93/go-neo4j-sample/lib"
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
		"CREATE (n:Sample) SET n.message = $message RETURN n.message + ', from node ' + id(n)",
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
		fmt.Println("Match Result Label ====>", matchQueryResult.Labels())
	} else {
		fmt.Println("Match Node is Nil")
	}

	removeQueryResult, removeQueryError := lib.RemoveQuery(
		"MATCH (n{message: 'hello, world22'}) REMOVE n.message return n",
		map[string]interface{}{"message": "hello, world"},
	)

	if removeQueryError != nil {
		fmt.Println("Remove Query Error ====>", removeQueryError.Error())
		return
	}

	if removeQueryResult != nil {
		fmt.Println("Remove Result Label ====>", removeQueryResult.Labels())
	} else {
		fmt.Println("Remove Node is Nil")
	}
}

func main() {
	start()
}
