package restapi

import (
	"fmt"
	"net/http"
)

func SpotifyConnect(name *string) {

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", nil)

	req.Header.Set("Authorization", "...")

	resp, err := client.Do(req)

	fmt.Println(resp)
	fmt.Println(err)
}
