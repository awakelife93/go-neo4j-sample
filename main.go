package main

import (
	"fmt"
	"strconv"

	"github.com/awakelife93/go-neo4j-sample/lib"
)

func init() {
	initializeResult, initializeError := lib.Initialize()

	if initializeError != nil {
		fmt.Println("Initialize Error ====>", initializeError.Error())
		lib.Clear()
	}

	fmt.Println(initializeResult)
}

func start() {
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
		for i := 0; i < len(matchQueryResult); i++ {
			fmt.Println("Match Result Id ====>", matchQueryResult[i].Id())
			fmt.Println("Match Result Label ====>", matchQueryResult[i].Labels())
			fmt.Println("Match Result Props ====>", matchQueryResult[i].Props())
		}
	} else {
		fmt.Println("Match Node is Nil")
	}

	updateQueryResult, updateQueryError := lib.Update(
		"MATCH (n{message: 'hello, world5'}) SET n.message = 'hi' return n",
		nil,
	)

	if updateQueryError != nil {
		fmt.Println("Update Query Error ====>", updateQueryError.Error())
		return
	}

	if updateQueryResult != nil {
		fmt.Println("Update Result Id ====>", updateQueryResult.Id())
		fmt.Println("Update Result Label ====>", updateQueryResult.Labels())
		fmt.Println("Update Result Props ====>", updateQueryResult.Props())
	} else {
		fmt.Println("Update Node is Nil")
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
		fmt.Println("Remove Result Id ====>", removeQueryResult.Id())
		fmt.Println("Remove Result Label ====>", removeQueryResult.Labels())
		fmt.Println("Remove Result Props ====>", removeQueryResult.Props())
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
