package neo4j

import (
	"time"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func Query(cypher string, params map[string]interface{}) (string, error) {
	sessionObject := getSession()

	queryResult, writeTransactionError := sessionObject.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			result, transactionError := transaction.Run(cypher, params)

			if transactionError != nil {
				return nil, transactionError
			}

			if result.Next() {
				return result.Record().Values, nil
			}

			return nil, result.Err()
		},

		func(config *neo4j.TransactionConfig) {
			config.Timeout = 60 * time.Second
		})

	if writeTransactionError != nil {
		return "", writeTransactionError
	}

	return queryResult.(string), nil
}
