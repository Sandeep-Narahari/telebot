package telebot

// ConvertBot.Handle("/start", func(c tele.Context) error {
// 	_, err := b.Send(c.Chat(), " Explore Autonomy:", &tele.ReplyMarkup{
// 		InlineKeyboard: [][]tele.InlineButton{
// 			{
// 				{Unique: "autonomy-app", Text: "Autonomy App", URL: "https://staging.app.autonomy.network/"},
// 			},
// 			{
// 				{Unique: "autonomy-explorer", Text: "Autonomy Explorer", URL: "https://staging.explorer.autonomy.network/"},
// 			},
// 			{
// 				{Unique: "autonomy-website", Text: "Autonomy Website", URL: "https://staging.autonomy.network/"},
// 			},
// 		},
// 	})
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return nil
// })

// func (b *Bot) Test() {
// 	fmt.Println("HELLOFOOF")
// 	b.telebot.Handle("/hello", func(c tele.Context) error {
// 		user := c.Sender()
// 		message := "Hello, " + user.FirstName + "!"
// 		_, err := b.telebot.Send(c.Recipient(), message)
// 		if err != nil {
// 			log.Fatal(err)
// 			return nil
// 		}
// 		return nil
// 	})
// }
