package main

import (
	"fmt"

	"github.com/awakelife93/go-neo4j-sample/neo4j"
)

func start() {
	error := neo4j.Initialize()
	fmt.Println("Initialize Error ====>", error)
}

func main() {
	start()
}
