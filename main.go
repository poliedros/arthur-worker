package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-mail/mail"

	"github.com/joho/godotenv"
)

func notify() {
	m := mail.NewMessage()

	m.SetHeader("From", os.Getenv("FROM"))

	m.SetHeader("To", os.Getenv("TO"))

	url := os.Getenv("URL")
	title := fmt.Sprintf("%s OFF", url)
	m.SetHeader("Subject", title)

	body := fmt.Sprintf("Check <a href='%s'>%s</a> now", url, url)
	m.SetBody("text/html", body)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Println("Port casting error.")
	}

	d := mail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("FROM"), os.Getenv("PASSWORD"))

	if err := d.DialAndSend(m); err != nil {

		panic(err)

	}

	fmt.Println("Email sent.")
}

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
		notify()
		return
	}
	defer res.Body.Close()

	fmt.Println("Status code: ", res.Status)
}
