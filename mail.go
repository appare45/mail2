package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/appare45/mail2/config"
	"github.com/appare45/mail2/message"
	"github.com/appare45/mail2/smtp/client"
	"github.com/appare45/mail2/smtp/commands"
	"github.com/appare45/mail2/smtp/entity"
	"github.com/appare45/mail2/smtp/server"
)

// クライアントとサーバを通信させてテストするコード
func main() {
	configPath := "config.toml"
	slog.SetLogLoggerLevel(slog.LevelDebug)
	config, err := config.Init(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	server := server.NewServer(fmt.Sprintf(":%d", config.Smtp.Port), entity.NewDomain("example.local"))
	go server.Start()
	time.Sleep(1 * time.Second)
	ip, err := config.Smtp.Ipaddr()
	if err != nil {
		fmt.Println(err)
		return
	}
	smtpClint, err := client.NewSmtpClient(*ip)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer smtpClint.Close()
	defer smtpClint.Command(commands.NewQuit())
	smtpClint.Command(commands.NewEhlo(entity.NewDomain("example.com")))
	smtpClint.Command(commands.NewMailFrom(entity.NewEmail("test", entity.NewDomain("example.com"))))
	smtpClint.Command(commands.NewRcptTo(entity.NewEmail("test", entity.NewDomain("example.local"))))
	message := message.NewMessage(
		*message.NewHeader(
			"test@example.com",
			message.NewDateTime(time.Now()),
		),
		message.NewBody("test"))
	smtpClint.Command(commands.NewData(message.Data_stream()))
}
