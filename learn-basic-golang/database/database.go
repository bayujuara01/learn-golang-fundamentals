package database

import "fmt"

type Connection struct {
	Name    string
	Address string
	Port    int32
}

var connection Connection

func init() {
	fmt.Println("Init dipanggil")
	connection = Connection{"Test", "https", 8888}
}

func GetDatabase() *Connection {
	return &connection
}
