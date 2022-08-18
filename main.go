package main

import (
	"fmt"
	"strconv"

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

	for i := 1; i <= 10; i++ {
		createQueryResult, createQueryError := lib.Create(
			"CREATE (n:Sample) SET n.message = $message RETURN n.message + ', from node ' + id(n)",
			map[string]interface{}{"message": "hello, world" + strconv.Itoa(i)},
		)

		if createQueryError != nil {
			fmt.Println("Create Query Error ====>", createQueryError.Error())
		}

		fmt.Println("Create Result ====>", createQueryResult)
	}

	matchQueryResult, matchQueryError := lib.Match(
		"MATCH (n) RETURN n",
		nil,
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

	removeQueryResult, removeQueryError := lib.Remove(
		"MATCH (n{message: 'hello, world1'}) REMOVE n.message return n",
		nil,
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

	deleteQueryError := lib.Delete(
		"MATCH (n{message: 'hello, world2'}) Delete n return n",
		nil,
	)

	if deleteQueryError != nil {
		fmt.Println("Delete Query Error ====>", deleteQueryError.Error())
		return
	}

	allDeleteQueryError := lib.Delete(
		"MATCH (n) DETACH DELETE n",
		nil,
	)

	if allDeleteQueryError != nil {
		fmt.Println("All Delete Query Error ====>", deleteQueryError.Error())
		return
	}
}

func main() {
	start()
}
