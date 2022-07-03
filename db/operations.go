package db

import (
	typesv1 "backend/gen/types/v1"
	"database/sql"
)

func CreateOperation(db *sql.DB, operationId int64, operationType typesv1.OperationType, source string, p1 string, p2 string, p3 string) (*typesv1.Operation, error) {
	prep, err := db.Prepare("insert into operations (id, type, source, param1, param2, param3) value (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = prep.Exec(operationId, operationType.Number(), source, p1, p2, p3)

	// operationを返すべきか否か。
	return nil, err
}

func GetOperationWithOperationId(db *sql.DB, operationId int64) (*typesv1.Operation, error) {
	prep, err := db.Prepare("select type, source, param1, param2, param3 from operations where id = ?")
	if err != nil {
		return nil, err
	}

	var opType int32
	var source string
	var param1 sql.NullString
	var param2 sql.NullString
	var param3 sql.NullString
	err = prep.QueryRow(operationId).Scan(&opType, &source, &param1, &param2, &param3)
	if err != nil {
		return nil, err
	}

	destinations, err := GetOperationDestinations(db, operationId)
	if err != nil {
		return nil, err
	}

	var p1 *string
	if param1.Valid {
		p1 = &param1.String
	}
	var p2 *string
	if param2.Valid {
		p2 = &param2.String
	}
	var p3 *string
	if param3.Valid {
		p3 = &param3.String
	}

	return &typesv1.Operation{
		Id:          operationId,
		Type:        typesv1.OperationType(opType),
		Source:      source,
		Destination: destinations,
		Param1:      p1,
		Param2:      p2,
		Param3:      p3,
	}, nil
}
