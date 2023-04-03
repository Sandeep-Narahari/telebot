package telebot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	Telebot *tele.Bot
}

func NewBot() (*Bot, error) {
	botCh := make(chan *tele.Bot)

	go func() {
		envPath := "./cmd/.env"
		err := os.Setenv("DOTENV_PATH", envPath)
		if err != nil {
			log.Fatal(err)
		}
		// Load the environment variables from the .env file
		err = godotenv.Load(os.Getenv("DOTENV_PATH"))
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		pref := tele.Settings{
			Token:  os.Getenv("token"),
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		}
		b, err := tele.NewBot(pref)
		if err != nil {
			log.Fatal(err)
		}
		// Send the bot instance through the channel
		botCh <- b

		// Start the bot in another goroutine to prevent blocking
		// go func() {
		// 	b.Start()
		// }()
	}()

	// Wait for the instance to be sent through the channel
	teleBot := <-botCh

	fmt.Println("BOT CONNECTED ")

	return &Bot{
		Telebot: teleBot,
	}, nil
}
func (d Bot) TestHandlers() {
	d.Telebot.Handle("/hello", func(c tele.Context) error {
		user := c.Sender()
		message := "Hello, " + user.FirstName + "!"
		_, err := d.Telebot.Send(c.Recipient(), message)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return nil
	})
}

// func NewConvertBot() Bot {
// 	botCh := make(chan *tele.Bot)
// 	go ConnectTeleBot(botCh)
// 	fmt.Println("dfdf")
// 	convertBot := &Bot{
// 		telebot: <-botCh,
// 	}
// 	fmt.Println("CON", convertBot)
// 	return *convertBot
// }
// func ConnectTeleBot(ch chan<- *tele.Bot) (*tele.Bot, error) {

// 	envPath := "./cmd/.env"
// 	err := os.Setenv("DOTENV_PATH", envPath)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	// Load the environment variables from the .env file
// 	err = godotenv.Load(os.Getenv("DOTENV_PATH"))
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 		return nil, err
// 	}
// 	pref := tele.Settings{
// 		Token:  os.Getenv("token"),
// 		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
// 	}
// 	fmt.Println("PdfdfREF", os.Getenv("token"))

// 	b, err := tele.NewBot(pref)
// 	if err != nil {
// 		return nil, err
// 	}
// 	b.Start()
// 	ch <- b

// 	return b, nil
// }
