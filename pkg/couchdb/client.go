package couchdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	strings "strings"
)

// ConnectionConfig is a configuration to connect on couchdb
type ConnectionConfig struct {
	host string
	port int
	db   string
}

// Connection is a connection with couchdb
type Connection struct {
	config *ConnectionConfig
}

func (config *ConnectionConfig) endpointURL(endpoint string) string {
	return fmt.Sprintf("http://%s:%d/%s", config.host, config.port, endpoint)
}

func (config *ConnectionConfig) baseURL() string {
	return config.endpointURL("")
}

func (config *ConnectionConfig) dbURL() string {
	return config.endpointURL(config.db)
}

// Connect create a new connection if the connection config is correct
func Connect(config *ConnectionConfig) (*Connection, error) {
	response, err := http.Get(config.baseURL())
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if !strings.Contains(string(data), "version") {
		return nil, err
	}

	return &Connection{config: config}, nil
}

const postContentType = "application/json"

// Store stores a document on the connected database
func (connection *Connection) Store(document *interface{}) (body string, err error) {
	jsonValue, _ := json.Marshal(document)
	buffer := bytes.NewBuffer(jsonValue)
	url := connection.config.dbURL()
	resp, err := http.Post(url, postContentType, buffer)

	if err != nil {
		return "", err
	}
	data, _ := ioutil.ReadAll(resp.Body)
	return string(data), nil
}

// QueryJSON queries the server using the queryJson as part of a mango query
func (connection *Connection) QueryJSON(queryJSON string) (*[]map[string]interface{}, error) {
	buffer := bytes.NewBufferString(queryJSON)
	url := connection.config.dbURL()
	resp, err := http.Post(url, postContentType, buffer)

	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	byt := []byte(data)
	var dat []map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		return nil, err
	}

	return &dat, nil
}
