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
	return fmt.Sprintf("http://%s:%d/%s/%s", config.host, config.port, config.db, endpoint)
}

func (config *ConnectionConfig) documentURL(endpoint string) string {
	return fmt.Sprintf("http://%s:%d/%s/%s", config.host, config.port, config.db, endpoint)
}

func (config *ConnectionConfig) baseURL() string {
	return fmt.Sprintf("http://%s:%d", config.host, config.port)
}

func (config *ConnectionConfig) dbURL() string {
	return config.endpointURL("")
}

// Connect create a new connection if the connection config is correct
func Connect(config *ConnectionConfig) (*Connection, error) {
	configuration := config
	if config == nil {
		configuration = &ConnectionConfig{
			host: "127.0.0.1",
			port: 5984,
			db:   "db",
		}
	}
	response, err := http.Get(configuration.baseURL())
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if !strings.Contains(string(data), "version") {
		return nil, err
	}

	return &Connection{config: configuration}, nil
}

const postContentType = "application/json"

type couchResponse struct {
	Ok     bool   `json:"ok"`
	Err    string `json:"error"`
	Reason string `json:"reason"`
}

func validateResult(data []byte) error {
	var res couchResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}
	if len(res.Err) > 0 {
		return fmt.Errorf("%s Reason: %s", res.Err, res.Reason)
	}
	return nil
}

// Store stores a document on the connected database
func (connection *Connection) Store(document interface{}) (data []byte, err error) {
	jsonValue, _ := json.Marshal(document)
	buffer := bytes.NewBuffer(jsonValue)
	url := connection.config.dbURL()
	resp, err := http.Post(url, postContentType, buffer)

	if err != nil {
		return nil, err
	}
	data, _ = ioutil.ReadAll(resp.Body)
	if err = validateResult(data); err != nil {
		return nil, err
	}
	return data, nil
}

// QueryJSON queries the server using the queryJson as part of a mango query
func (connection *Connection) QueryJSON(queryJSON string) (*[]map[string]interface{}, error) {
	buffer := bytes.NewBufferString(queryJSON)
	url := connection.config.endpointURL("_find")
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

// QueryJSONRaw queries the server using the queryJson as part of a mango query and returns the raw
func (connection *Connection) QueryJSONRaw(queryJSON string) ([]byte, error) {
	buffer := bytes.NewBufferString(queryJSON)
	url := connection.config.endpointURL("_find")
	resp, err := http.Post(url, postContentType, buffer)

	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	return []byte(data), nil
}

// Load loads a document by its _id
func (connection *Connection) Load(id string) (*map[string]interface{}, error) {
	url := connection.config.endpointURL(fmt.Sprintf("/%s", id))
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	byt := []byte(data)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		return nil, err
	}

	return &dat, nil
}

// LoadRaw loads a document by its _id returning the json string as byte slice
func (connection *Connection) LoadRaw(id string) ([]byte, error) {
	url := connection.config.endpointURL(fmt.Sprintf("/%s", id))
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	return []byte(data), nil
}
