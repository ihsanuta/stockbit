package main

import (
	"fmt"
	"os"
	"stockbit/app/global"
	"stockbit/app/route"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env is not loaded properly")
		fmt.Println(err)
		os.Exit(2)
	}
}

func main() {
	g := gin.Default()
	route.Init(g)

	serverHost := os.Getenv(global.EnvServerHost)
	serverPort := os.Getenv(global.EnvServerPort)
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)

	g.Run(serverString)
}
