package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("error loading .env file")
	}
	return os.Getenv(key)
}

func randomAC(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	bot := slacker.NewClient(goDotEnvVariable("API_TOKEN"), slacker.WithDebug(true))

	// daily := &slacker.Event

	definition := &slacker.CommandDefinition{
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			rand.Seed(time.Now().UnixNano())
			weekday := time.Now().Weekday()
			ac := strconv.Itoa(randomAC(1, 20))
			attachement := []slack.Block{}
			attachement = append(attachement, slack.NewContextBlock("1",
				slack.NewTextBlockObject("mrkdwn", weekday.String()+"'s AC: "+ac, false, false)),
			)
			response.Reply("", slacker.WithBlocks(attachement))
		},
	}

	bot.Command("test", definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		panic(err)
	}
}
