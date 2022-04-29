package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
	"github.com/joho/godotenv"
)

func env(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error loading .env file")
	}

	return os.Getenv(key)
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent)  {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}	
}

func main()  {
	botToken := env("SLACK_BOT_TOKEN")
	appToken := env("SLACK_APP_TOKEN")

	// fmt.Printf("botToken: %s = %s\n", "SLACK_BOT_TOKEN", botToken)
	// fmt.Printf("botToken: %s = %s\n", "SLACK_APP_TOKEN", appToken)

	bot := slacker.NewClient(botToken, appToken)

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
		Description: "year of birth calculator",
		Example: "my year of birth is 2020",
		Handler: func (botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter)  {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := time.Now().Year() - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	} 
}