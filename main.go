package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	"strings"
)

type AWS struct {
	Status   string `csv:"status_code"`
	Method   string `csv:"http_method"`
	Uri      string `csv:"absolute_uri"`
	Relative string
}

func main() {
	in, err := os.Open("api.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	clients := []*AWS{}

	if err := gocsv.UnmarshalFile(in, &clients); err != nil {
		panic(err)
	}

	var store = map[string]int{}

	for _, client := range clients {
		relative := strings.Split(client.Uri, "?")
		client.Relative = relative[0]
		fmt.Printf("%+v\n", client.Relative)
		store[client.Relative] += 1
	}

	fmt.Printf("%v\n", store)

}
