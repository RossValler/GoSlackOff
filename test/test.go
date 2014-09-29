package main

import (
	"fmt"
	"github.com/rossvaller/goslackoff"
	"io/ioutil"
)

func main() {
	//-----------------------------------------
	// Create an instance of GSO
	//-----------------------------------------
	instance := goslackoff.New("SLACK-USERNAME", "SLACK-INCOMING-HOOK-TOKEN", "GoSlackOff", "#general")

	//-----------------------------------------
	// Send the message
	//-----------------------------------------
	success, resp := instance.SendMessage(goslackoff.Message{
		Text:     "Hello from Go :)",
		Channel:  "#other-room",      //optional
		Username: "GoSlackOff Other", //optional
	})

	fmt.Println("Message sent successfully?", success)

	if !success {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Response Code:", resp.Status)
		fmt.Println("Response Body:", body)
	}
}
