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

type User struct {
	ID      uint64
	Name    string
	Surname string
	Email   string
}

type response struct {
	Succ bool
	Msg  string
	User User
}

func printer(data response) {
	if data.Succ {
		fmt.Println("----------")
		fmt.Printf("ID: %v\nName: %v\nSurname: %v\nEmail: %v\n----------\n%v\n",
			data.User.ID, data.User.Name, data.User.Surname, data.User.Email, data.Msg)
		fmt.Println("----------")
	} else {
		fmt.Println("----------")
		fmt.Printf("Error: %v\n", data.Msg)
		fmt.Println("----------")
	}
}

func getValues(reader *bufio.Reader) (url.Values, error) {
	fmt.Print("Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		return url.Values{}, err
	}
	fmt.Print("Surname: ")
	surname, err := reader.ReadString('\n')
	if err != nil {
		return url.Values{}, err
	}
	fmt.Print("Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		return url.Values{}, err
	}

	name = name[:len(name)-1]
	surname = surname[:len(surname)-1]
	email = email[:len(email)-1]

	v := url.Values{
		"name":    {name},
		"surname": {surname},
		"email":   {email},
	}

	return v, nil
}

func sendRequest(reqType string, addr string, id string) (response, error) {
	var data response
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

func addUser(reader *bufio.Reader, addr string) (response, error) {
	var data response
	v, err := getValues(reader)
	if err != nil {
		return data, err
	}
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

		b, _, err := reader.ReadLine()
		if err != nil {
			log.Println(err)
			return
		}
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
