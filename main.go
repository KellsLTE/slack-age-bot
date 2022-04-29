package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

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
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-2643950721141-3443620972999-y4WOc43IZ8LKXaLLIVIcJTRl")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03D8622MRC-3469483265793-18beeff27830475e4648dda31828e282b9e7a35c6d3b241d7e3a51b5d72a68fd")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

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