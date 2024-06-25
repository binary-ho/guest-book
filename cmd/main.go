package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"guestbook/cmd/application"
)

func main() {
	guestBookApplication := application.InitApplication()
	fmt.Println("Hello, 世界!")
	guestBookApplication.Run()
}
