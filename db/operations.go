package db

import (
	typesv1 "backend/gen/types/v1"
	"database/sql"
)

func CreateOperation(db *sql.DB, operationId int64, operationType typesv1.OperationType, source string) (*typesv1.Operation, error) {
	prep, err := db.Prepare("insert into operations (id, type, source) value (?, ?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = prep.Exec(operationId, operationType.Number(), source)

	// operationを返すべきか否か。
	return nil, err
}

func GetOperationWithOperationId(db *sql.DB, operationId int64) (*typesv1.Operation, error) {
	prep, err := db.Prepare("select type, source from operations where id = ?")
	if err != nil {
		return nil, err
	}

	var opType int32
	var source string
	err = prep.QueryRow(operationId).Scan(&opType, &source)
	if err != nil {
		return nil, err
	}

	destinations, err := GetOperationDestinations(db, operationId)
	if err != nil {
		return nil, err
	}

	return &typesv1.Operation{
		Id:          operationId,
		Type:        typesv1.OperationType(opType),
		Source:      source,
		Destination: destinations,
	}, nil
}
