package main

import (
	"api/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando api!!")

	router := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", router))
}
