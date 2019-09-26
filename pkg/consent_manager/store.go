package consent_manager

import (
	"encoding/json"
	"fmt"

	couchdb "github.com/joseteodoro/consent-enforcement/pkg/couchdb"
)

// DataTypeByID loads the DataTypeDocument by its ID
func DataTypeByID(ID string) (*DataTypeDocument, error) {
	connection, err := couchdb.Connect(nil)
	if err != nil {
		panic(err)
	}

	bytes, err := connection.LoadRaw(ID)
	if err != nil {
		return nil, err
	}
	var target DataTypeDocument
	if err := json.Unmarshal(bytes, &target); err != nil {
		return nil, err
	}

	return &target, nil
}

// StoreDataType stores the DataTypeDocument
func StoreDataType(dataType *DataType) error {
	connection, err := couchdb.Connect(nil)
	if err != nil {
		return err
	}
	document := MappingFromDataType(dataType)
	_, err = connection.Store(document)
	return err
}

type dataTypes struct {
	docs []*DataType
}

// ListDataTypes list all datatypes
func ListDataTypes() ([]*DataType, error) {
	connection, err := couchdb.Connect(nil)
	if err != nil {
		panic(err)
	}

	bytes, err := connection.QueryJSONRaw(`
	{
		"selector": {
		   "type": "DataType"
		}
	 }
	`)
	if err != nil {
		return nil, err
	}
	fmt.Printf(string(bytes))
	var rows dataTypes
	if err := json.Unmarshal(bytes, &rows); err != nil {
		return nil, err
	}

	return rows.docs, nil
}
