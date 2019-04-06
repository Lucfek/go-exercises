package main

import (
    "fmt"
    "strings"
    "io"
    "io/ioutil"
    "log"
    "flag"
)

func words(r io.Reader) (even []string, odd []string) {
    buf, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Fields(string(buf))
    
    var vavelCount int
    
    for _, word := range words {
        vavelCount = 0
        for _, letter := range word {
            switch letter {
                case 'a', 'e', 'i', 'o', 'u', 'y', 'A', 'E', 'I', 'O', 'U', 'Y', 'ą', 'ę', 'ó', 'Ą', 'Ę', 'Ó': vavelCount++
            }
        }
        if vavelCount % 2 == 0 {
            even = append(even, word)
        } else {
            odd = append(odd, word)
        }
    }
    
    return even, odd
}


func main() {
    filePtr := flag.String("file", "lorem.txt", "File you want to open")
    flag.Parse()
    
    b, err := ioutil.ReadFile(*filePtr)
    if err != nil {
        log.Fatal(err)
    }
    
    even, odd := words(strings.NewReader(string(b)))
    
    err = ioutil.WriteFile("even.txt", []byte(strings.Join(even, " ")), 0644)
    if err != nil {
        log.Fatal(err)
    }
    
    err = ioutil.WriteFile("odd.txt", []byte(strings.Join(odd, " ")), 0644)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(even, odd)
}