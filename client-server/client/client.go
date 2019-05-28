package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func printer(data map[string]interface{}) {
	fmt.Println("----------")
	for k, v := range data {
		fmt.Printf("%s: %v \n", k, v)
	}
	fmt.Println("----------")
}

func sendRequest(reqType string, addr string, id string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if id == "" {
		return data, errors.New("No specified Id")
	}
	client := &http.Client{}
	req, err := http.NewRequest(reqType, "http://"+addr+"/users/"+id+"/", nil)
	if err != nil {
		return data, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return data, err
	}

	json.NewDecoder(resp.Body).Decode(&data)
	return data, nil
}

func addUser(reader *bufio.Reader, addr string) (map[string]interface{}, error) {
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

	var data map[string]interface{}
	resp, err := http.PostForm("http://"+addr+"/users/", v)
	if err != nil {
		return data, err
	}
	json.NewDecoder(resp.Body).Decode(&data)
	return data, nil
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
			resp, err := addUser(reader, addr)
			if err != nil {
				log.Println(err)
			} else {
				printer(resp)
			}
		case "deleteuser":
			resp, err := sendRequest("DELETE", addr, splitet[1])
			if err != nil {
				log.Println(err)
			} else {
				printer(resp)
			}
		case "getuser":
			resp, err := sendRequest("GET", addr, splitet[1])
			if err != nil {
				log.Println(err)
			} else {
				printer(resp)
			}
		case "exit":
			return
		default:
			fmt.Printf("Unknow Tastk: %s \n", splitet[0])
		}

	}
}
