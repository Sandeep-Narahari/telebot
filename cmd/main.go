package main

import (
	"log"

	"github.com/Sandeep-Narahari/telegram_bot/telebot"
)

func main() {
	bot, err := telebot.NewBot()
	if err != nil {
		log.Fatal(err)
	}
	bot.TestHandlers()
	stop := make(chan struct{})
	defer close(stop)

	go func() {
		bot.Telebot.Start()
	}()

	<-stop
	// fmt.Println("dfdf")
	// go bot.Test()
	// fmt.Println("dfdfdfdf")
	// for {
	// 	time.Sleep(time.Second)
	// } // Newtelebot.Handle("/start", startHandler)
	// telebot.StartHandler(Newtelebot)
	// envPath := "./cmd/.env"
	// err := os.Setenv("DOTENV_PATH", envPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Load the environment variables from the .env file
	// err = godotenv.Load(os.Getenv("DOTENV_PATH"))
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// pref := tele.Settings{
	// 	Token: os.Getenv("token"),

	// 	Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	// }
	// fmt.Println("PREF", os.Getenv("token"))

	// b, err := tele.NewBot(pref)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// // selector := &tele.InlineButton{Unique: "hello button", Text: "hello botn", Data: "CLICK"}

	// b.Handle("/hello", func(c tele.Context) error {
	// 	user := c.Sender()
	// 	message := "Hello, " + user.FirstName + "!"
	// 	_, err := b.Send(c.Recipient(), message)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		return nil
	// 	}
	// 	return nil
	// })
	// b.Handle(tele.OnUserJoined, func(ctx tele.Context) error {
	// 	commands := []string{
	// 		"/start - start the bot",
	// 		"/help - show this help message",
	// 		"/hello - say hello to the bot",
	// 	}
	// 	helpText := "Available commands:\n\n" + strings.Join(commands, "\n")
	// 	_, err := b.Send(ctx.Chat(), helpText)

	// 	return err
	// })

	// b.Start()
}

// func (m telebot.Bot) StartHandler() {
// 	m
// }
