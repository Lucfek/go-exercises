package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getUser(id string) {
	resp, err := http.Get("http://127.0.0.1:8000/users/" + id + "/")
	if err != nil {
		log.Print(err)
	} else {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
		} else {
			fmt.Print(string(contents))
		}

	}
}

func addUser(reader *bufio.Reader) {
	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Surname: ")
	surname, _ := reader.ReadString('\n')
	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')

	name = name[:len(name)-1]
	surname = surname[:len(surname)-1]
	email = email[:len(email)-1]

	v := url.Values{}
	v.Set("name", name)
	v.Set("surname", surname)
	v.Set("email", email)

	resp, err := http.PostForm("http://127.0.0.1:8000/users/", v)
	if err != nil {
		log.Print(err)
	} else {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
		}
		fmt.Print(string(contents))
	}
}

func delUser(id string) {
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", "http://127.0.0.1:8000/users/"+id+"/", nil)
	if err != nil {
		log.Print(err)
		return
	}
	resp, fetchErr := client.Do(req)
	if err != nil {
		log.Print(fetchErr)
		return
	}

	contents, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(contents))

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nCommand: ")

		b, _, _ := reader.ReadLine()
		command := string(b)
		splitet := strings.Split(command, " ")

		switch splitet[0] {
		case "adduser":
			addUser(reader)
		case "deleteuser":
			if splitet[1] != "" {
				delUser(splitet[1])
			} else {
				log.Print("No ID specified")
			}
		case "getuser":
			if splitet[1] != "" {
				getUser(splitet[1])
			} else {
				log.Print("No ID specified")
			}
		}

	}
}
