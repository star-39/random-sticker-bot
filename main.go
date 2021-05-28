package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v3"
)

var STICKER_SET string
var BOT_TOKEN string
var COMMAND string
var stickers []tb.Sticker

func main() {
	STICKER_SET = os.Getenv("STICKER_SET")
	BOT_TOKEN = os.Getenv("BOT_TOKEN")
	COMMAND = os.Getenv("COMMAND")

	if len(STICKER_SET) == 0 || len(BOT_TOKEN) == 0 || len(COMMAND) == 0 {
		print("STICKER_SET or GROUP_ID or COMMAND not set! exiting")
		os.Exit(1)
	}
	b, err := tb.NewBot(tb.Settings{
		Token:  BOT_TOKEN,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	initializeBot(b)

	b.Handle("/"+COMMAND, func(m tb.Context) error {
		sendSticker(m, b)
		return nil
	})

	b.Start()
}

func initializeBot(b *tb.Bot) {
	stickerSet, err := b.StickerSet(STICKER_SET)
	if err != nil {
		print("Sticker set not fount!")
		os.Exit(1)
	}
	stickers = stickerSet.Stickers
}

func sendSticker(m tb.Context, b *tb.Bot) {
	time.Sleep(10 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	sticker := stickers[rand.Intn(len(stickers))]
	sticker.Send(b, m.Chat(), &tb.SendOptions{})
}
