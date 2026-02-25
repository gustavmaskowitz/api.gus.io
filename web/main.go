package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"syscall/js"
	"time"
)

type IpifyResponse struct {
	Ip string `json:"ip"`
}

func main() {
	c := make(chan struct{}, 0)

	go func() {
		output := js.Global().Get("document").Call("getElementById", "json-output")
		if output.IsNull() {
			fmt.Println("Element with id 'json-output' not found")
			return
		}

		ip := "Unknown (fetch failed)"
		resp, err := http.Get("https://api.ipify.org?format=json")
		if err == nil {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				var ipify IpifyResponse
				if err := json.Unmarshal(body, &ipify); err == nil {
					ip = ipify.Ip
				}
			}
		}

		now := time.Now().UTC().Format("2006-01-02 15:04:05")

		result := map[string]interface{}{
			"api.gus.io": map[string]string{
				"now_utc":    now,
				"Your_IP":    ip,
				"built with": "golang syscall/js",
				"more info":  "http://www.gus.io/",
			},
		}

		jsonBytes, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			output.Set("textContent", "Error generating JSON")
			return
		}

		output.Set("textContent", string(jsonBytes))
		close(c)
	}()

	<-c
}
