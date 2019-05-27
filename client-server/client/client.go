package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func sendRequest(reqType string, addr string, id string) (string, error) {
	if id == "" {
		return "", errors.New("No specified Id")
	}
	client := &http.Client{}
	req, err := http.NewRequest(reqType, "http://"+addr+"/users/"+id+"/", nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	contents, _ := ioutil.ReadAll(resp.Body)
	return string(contents), nil
}

func addUser(reader *bufio.Reader, addr string) {
	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Surname: ")
	surname, _ := reader.ReadString('\n')
	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')

	name = name[:len(name)-1]
	surname = surname[:len(surname)-1]
	email = email[:len(email)-1]

	v := url.Values{
		"name":    {name},
		"surname": {surname},
		"email":   {email},
	}

	resp, err := http.PostForm("http://"+addr+"/users/", v)
	if err != nil {
		log.Println(err)
	} else {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
		}
		fmt.Println(string(contents))
	}
}

func main() {

	var addr string
	flag.StringVar(&addr, "address", "127.0.0.1:8000", "Server IPv4 address")
	flag.Parse()

	fmt.Printf("Target IP- %s \n", addr)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Command: ")

		b, _, _ := reader.ReadLine()
		command := string(b)
		splitet := strings.Split(command, " ")

		switch splitet[0] {
		case "adduser":
			addUser(reader, addr)
		case "deleteuser":
			resp, err := sendRequest("DELETE", addr, splitet[1])
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println(resp)
			}
		case "getuser":
			resp, err := sendRequest("GET", addr, splitet[1])
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println(resp)
			}
		default:
			fmt.Printf("Unknow Tastk: %s \n", splitet[0])
		}

	}
}
