package consent_manager

import (
	"github.com/jakehl/goid"
)

// DataTypeDocument mapping from DataType to a document on couchdb
type DataTypeDocument struct {
	DBId     string `json:"_id"`
	Revision string `json:"_revision"`
	ID       string `json:"id"`
	Display  string `json:"display"`
	UUID     string `json:"uuid"`
	*DataType
}

// MappingFromDataType maps DataType to a document
func MappingFromDataType(src *DataType) *DataTypeDocument {
	id := src.ID
	if len(id) < 1 {
		id = goid.NewV4UUID().String()
	}
	return &DataTypeDocument{
		DBId:     id,
		ID:       id,
		Display:  src.Display,
		UUID:     id,
		DataType: src,
	}
}
