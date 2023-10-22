package main

import (
	"belajar-golang-dasar/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()

	LogDB(result)
	
}

func LogDB(db *database.Connection) {
	fmt.Println(db.Name)
}
