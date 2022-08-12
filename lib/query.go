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

			if transactionError != nil {
				return nil, transactionError
			}

			for result.Next() {
				returnedMap := result.Record().GetByIndex(0)
				return returnedMap, nil
			}

			resultError := result.Err()
			if resultError != nil {
				return nil, resultError
			}

			return nil, result.Err()
		},

		func(config *neo4j.TransactionConfig) {
			config.Timeout = 60 * time.Second
		})

	if writeTransactionError != nil {
		return nil, writeTransactionError
	}

	return queryResult, nil
}

func CreateQuery(cypher string, params map[string]interface{}) (string, error) {
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

func MatchQuery(cypher string, params map[string]interface{}) (interface{}, error) {
	queryResult, error := readTransaction(cypher, params)

	if error != nil {
		return nil, error
	}

	return queryResult, nil
}
