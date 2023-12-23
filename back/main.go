package main

import "github.com/GrandOichii/messager-app/back/router"

func main() {
	r := router.CreateRouter()

	r.Run()
}
