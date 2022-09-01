package lib

import (
	"fmt"
	"time"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func writeTransaction(cypher string, params map[string]interface{}) (interface{}, error) {
	fmt.Println("Start WriteTransaction")
	session := getSession()

	defer renewalSession(session)

	queryResult, writeTransactionError := session.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			result, transactionError := transaction.Run(cypher, params)

			if transactionError != nil {
				return nil, transactionError
			}

			if result.Next() {
				return result.Record().Values(), nil
			}

			return nil, result.Err()
		},

		func(config *neo4j.TransactionConfig) {
			config.Timeout = 60 * time.Second
		})

	if writeTransactionError != nil {
		return "", writeTransactionError
	}

	return queryResult, nil
}

func readTransaction(cypher string, params map[string]interface{}) (interface{}, error) {
	fmt.Println("Start ReadTransaction")
	session := getSession()

	defer renewalSession(session)

	queryResult, writeTransactionError := session.ReadTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			result, transactionError := transaction.Run(cypher, params)

			var nodes []neo4j.Node = nil

			if transactionError != nil {
				return nil, transactionError
			}

			for result.Next() {
				nodes = append(nodes, result.Record().GetByIndex(0).(neo4j.Node))
			}

			resultError := result.Err()
			if resultError != nil {
				return nil, resultError
			}

			return nodes, nil
		},

		func(config *neo4j.TransactionConfig) {
			config.Timeout = 60 * time.Second
		})

	if writeTransactionError != nil {
		return nil, writeTransactionError
	}

	return queryResult, nil
}

func Create(cypher string, params map[string]interface{}) (string, error) {
	queryResult, error := writeTransaction(cypher, params)

	if error != nil {
		return "", error
	}

	// * Since the return type is different for each crud, it is defined as interface{},
	// * so create is forced to infer because it is an array type.
	var resultArray = queryResult.([]interface{})
	var result string = resultArray[0].(string)

	return result, nil
}

func Match(cypher string, params map[string]interface{}) ([]neo4j.Node, error) {
	var result []neo4j.Node = nil
	queryResult, error := readTransaction(cypher, params)

	if error != nil {
		return nil, error
	}

	if queryResult != nil {
		result = queryResult.([]neo4j.Node)
	}

	return result, nil
}

func Remove(cypher string, params map[string]interface{}) (neo4j.Node, error) {
	var result neo4j.Node = nil
	queryResult, error := writeTransaction(cypher, params)

	if error != nil {
		return nil, error
	}

	if queryResult != nil {
		var queryResultArray = queryResult.([]interface{})
		result = queryResultArray[0].(neo4j.Node)
	}

	return result, nil
}

func Delete(cypher string, params map[string]interface{}) error {
	_, error := writeTransaction(cypher, params)

	if error != nil {
		return error
	}

	return nil
}

func Update(cypher string, params map[string]interface{}) (neo4j.Node, error) {
	var result neo4j.Node = nil
	queryResult, error := writeTransaction(cypher, params)

	if error != nil {
		return nil, error
	}

	if queryResult != nil {
		var queryResultArray = queryResult.([]interface{})
		result = queryResultArray[0].(neo4j.Node)
	}

	return result, error
}
