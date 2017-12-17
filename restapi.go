package restapi

import (
	"bytes"
	"fmt"
	"net/http"
)

func ApiCall() {
	req, err := http.NewRequest("GET", "https://discordapp.com/api/v6", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(req)
}
