package main

import (
	"log"

	"github.com/ValentinZoia/gin/internal/router"
)

func main() {
	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
