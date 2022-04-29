# slack-age-bot
A Slack bot built with Go

# Note 
make sure you have go installed on your system and it is available in your path

# Setup 
Go to [slack-api](https://api.slack.com) to create a slack application and add it to a workspace before continuing

You will need a slack bot token and a slack app token to make your application work. Once you have these two keys clone the repo, copy the contents of the env.example file and paste them into a .env file you have created in your project directory and add the keys to the file then run the following

```go
go build
```

```go
go run main.go
```
