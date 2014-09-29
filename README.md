GoSlackOff
==========

Send messages to Slack channels from Go

## Usage
First create an instance with your Slack company name and token as well as the username for the bot and a channel to post to:

```go
instance := goslackoff.New("SLACK-USERNAME", "SLACK-INCOMING-HOOK-TOKEN", "GoSlackOff", "#general")
```
#### Then to send a message:
```go
success, resp := instance.SendMessage(goslackoff.Message{
	Text:     "Hello from Go :)",
	Channel:  "#other-room",      //optional
	Username: "GoSlackOff Other", //optional
})
```
#### Catch the errors/http response with:
```go
fmt.Println("Message sent successfully?", success)

if !success {
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Code:", resp.Status)
	fmt.Println("Response Body:", body)
}
```

## Demo
Demo is included in test/test.go
