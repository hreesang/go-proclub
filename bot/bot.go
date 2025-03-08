package bot

import (
	dg "github.com/bwmarrin/discordgo"
	"github.com/hreesang/go-proclub/bot/config"
	"github.com/hreesang/go-proclub/bot/utils"
)

var (
	Session 	*dg.Session
)


func onConnected(_ *dg.Session, _ *dg.Ready) {
	utils.Log.Printf("Connected as %v#%v!\n", Session.State.User.Username, Session.State.User.Discriminator)
}

func Run() error {
	var err error
	
	utils.Log.Println("Authenticating the bot...")
	Session, err = dg.New("Bot " + config.BotToken)
	if err != nil {
		utils.Log.Fatalln("Failed to authenticate:", err)
		return err
	}

	AddEventHandler(onConnected)
	registerEventHandlers()

	utils.Log.Println("Connecting the bot...")
	err = Session.Open()
	if err != nil {
		utils.Log.Fatalln("Failed to connect:", err)
		return err
	}
	return err
}

func Stop() error {
	err := Session.Close()
	WaitEventHandlers()
	return err
}
