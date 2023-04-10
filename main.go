package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error getting environment variables.")
		return
	}

	fmt.Println(os.Getenv("URL"))
	res, err := http.Get(os.Getenv("URL"))

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Status code: ", res.Status)
}
