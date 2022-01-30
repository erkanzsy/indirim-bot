package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func notify(message string, imgUrl string) string {
	postBody, _ := json.Marshal(map[string]string{
		"chat_id" : "-720855660",
		"parse_mode": "HTML",
		"photo": imgUrl,
		"caption": message,
	})

	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(os.Getenv("NOTIFY_URL"), "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	return sb
}
