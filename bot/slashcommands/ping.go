package slashcommands

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	AddSlashCommand(
		&discordgo.ApplicationCommand{
			Name: 			"ping",
			Description:	"Replies to you with a Pong.",
		},
		ping,
	)
}

func ping(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type:	discordgo.InteractionResponseChannelMessageWithSource,
		Data:	&discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
	return nil
}
