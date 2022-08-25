package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Names []string

func (n Names) print() {
	for _, name := range n {
		fmt.Println("Hello "+name+"... Welcome to GoLang")
	}
}

func (n Names) toString(sepeartor string) string{
	return strings.Join([]string(n), sepeartor)
}

func (n Names) saveToFile(filename string, sep string) error{
	return ioutil.WriteFile(filename, []byte(n.toString(sep)), 0666)
}

func readNamesFromFile(filename string) Names{
	names, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println("Error --> ", err)
		os.Exit(1)
	}

	name := strings.Split(string(names), " ")

	return Names(name)

}

func (n Names) shuffle() Names{
	source := rand.NewSource(time.Now().UnixNano())
	ra := rand.New(source)

	for i , _ := range n {
		num := ra.Intn(len(n)-1)
		n[i], n[num] = n[num], n[i]
		// temp := n[num]
		// n[num] = n[i]
		// n[i] = temp
	}

	return n
} 