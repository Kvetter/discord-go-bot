package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Start api call")
	b := make([]byte, 20)
	req, err := http.NewRequest("GET", "https://discordapp.com/api/v6", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*req)
}
