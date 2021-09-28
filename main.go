package main

import (
	"test.com/config"
	"test.com/routes"
)

func main() {
    r := routes.SetupRouter()
    config.InitApp()
    defer config.FireStoreClient.Close()
    r.Run()
}

