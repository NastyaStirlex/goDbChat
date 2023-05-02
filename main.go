package main

import (
	db2 "awesomeProject/db"
	"awesomeProject/models"
	"fmt"
	"github.com/containrrr/shoutrrr"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

const (
	token   = "6117072333:AAHAFOzN_P6Z09fLQFcL8cDn_YCiyW8osUU"
	botName = "@test_channel_goland"
)

var countLikes = 0
var countUnlikes = 0

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Like("+string(countLikes)+")", "clickLike"),
		tgbotapi.NewInlineKeyboardButtonData("Unlike("+string(countUnlikes)+")", "clickUnlike"),
	),
)

func main() {

	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}

	db2.Init()
	db := db2.GetDb()

	a, err := db.DB()
	if err != nil {
		fmt.Println(err)
	}

	n, err := migrate.Exec(a, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)

	// telegram
	url := "telegram://" + token + "@telegram/?chats=" + botName

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates, _ := bot.GetUpdatesChan(updateConfig)

	var mem *models.Mem
	db.Model(models.Mem{}).Order("random()").First(&mem)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {

				db.FirstOrCreate(&models.User{
					UserId: uint(update.Message.From.ID),
					Name:   update.Message.From.UserName,
				})

				db.FirstOrCreate(&models.Result{
					MemId:  mem.Id,
					UserId: uint(update.Message.From.ID),
					IsLike: models.Undefined,
				}, models.Result{UserId: uint(update.Message.From.ID), MemId: mem.Id})

				msg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, mem.Link)

				msg.Caption = "Оцени мем:"
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ReplyMarkup = numericKeyboard

				if _, err = bot.Send(msg); err != nil {
					panic(err)
				}

				err = shoutrrr.Send(url, update.Message.Text)
				if err != nil {
					panic(err)
					return
				}
			}

		} else if update.CallbackQuery != nil {
			//callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			//if _, err := bot.MakeRequest(callback); err != nil {
			//	panic(err)
			//}
			if update.CallbackQuery.Data == "clickLike" {
				var result *models.Result
				db.Model(models.Result{}).Where("mem_id = ? AND user_id = ?", mem.Id, update.Message.From.ID).First(&result)
				result.IsLike = models.Like
				db.Save(&result)
				//db.Model(models.Result{}).Where("mem_id = ? AND user_id = ?", mem.Id, update.Message.From.ID).Update("is_like", 1)

			} else if update.CallbackQuery.Data == "clickUnlike" {
				//sqlStateDislike := "UPDATE mems_db.public.results SET islike=false WHERE mem_id=$1 AND user_id=$2"
				//_, err = db.Exec(sqlStateDislike, fileId, update.Message.From.ID)
				if err != nil {
					panic(err)
				}
			}
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}

		}
	}

	// teams
	//webhook_url := "https://m365x012372.webhook.office.com"

}
