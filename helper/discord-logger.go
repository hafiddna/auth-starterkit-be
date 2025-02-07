package helper

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/hafiddna/auth-starterkit-be/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	logType, title, message, service string
	detail                           map[string]interface{}
	color                            int
)

func DiscordLogger(logTypeParam, titleParam, messageParam, serviceParam string, detailParam map[string]interface{}) {
	logType = logTypeParam
	title = titleParam
	message = messageParam
	service = serviceParam
	detail = detailParam

	discord, err := discordgo.New("Bot " + config.Config.App.Discord.Token)
	if err != nil {
		log.Println("error creating Discord session,", err)
	}
	defer func(discord *discordgo.Session) {
		err := discord.Close()
		if err != nil {
			log.Println("error closing Discord session,", err)
		}
	}(discord)

	discord.AddHandler(sendReadyMessageToChannel)
	bot := discord.Open()
	if bot != nil {
		log.Println("error opening connection,", bot)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func sendReadyMessageToChannel(s *discordgo.Session, m *discordgo.Ready) {
	var field []*discordgo.MessageEmbedField

	for key, value := range detail {
		var stringValue string

		switch value.(type) {
		case string, int, int32, int64, float64, float32, complex64, complex128:
			stringValue = fmt.Sprintf("%v", value)
		case []interface{}, map[string]interface{}:
			marshal, err := json.Marshal(value)
			if err != nil {
				log.Println("error marshalling value,", err)
			}
			stringValue = string(marshal)
		case bool:
			if value.(bool) {
				stringValue = "true"
			} else {
				stringValue = "false"
			}
		case nil:
			stringValue = "null"
		}

		field = append(field, &discordgo.MessageEmbedField{
			Name:  key,
			Value: stringValue,
		})
	}

	switch logType {
	case "info":
		color = 0x00ff00
	case "warning":
		color = 0xffff00
	case "error":
		color = 0xff0000
	}

	embed := &discordgo.MessageEmbed{
		Type:        "article",
		Title:       title,
		Description: message,
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       color,
		Footer: &discordgo.MessageEmbedFooter{
			Text: service,
		},
		Fields: field,
	}

	_, _ = s.ChannelMessageSendEmbed("1191312140343189615", embed)
}
