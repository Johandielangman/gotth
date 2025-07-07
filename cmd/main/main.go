// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~
//      /\_/\
//     ( o.o )
//      > ^ <
//
// Author: Johan Hanekom
// Date: July 2025
//
// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~

package main

import (
	"fmt"
	"gotth/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("Starting... 🚀")

	// ====> INITIALIZE A NEW APP
	app := server.NewApp()

	// ====> START THE APPS
	app.Start()
}
