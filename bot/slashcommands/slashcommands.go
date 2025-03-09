package slashcommands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/hreesang/go-proclub/bot"
	"github.com/hreesang/go-proclub/bot/utils"
)

type AddedSlashCommand struct {
	data	*discordgo.ApplicationCommand
	execute	func (*discordgo.Session, *discordgo.InteractionCreate) error
}

var addedSlashCommands = make(map[string]AddedSlashCommand)
var registeredGuildsCommands = make(map[string][]*discordgo.ApplicationCommand)


func AddSlashCommand(cmd *discordgo.ApplicationCommand, fn func (*discordgo.Session, *discordgo.InteractionCreate) error) error {
	if _, exists := addedSlashCommands[cmd.Name]; exists {
		return fmt.Errorf("command already exists: %v", cmd.Name)
	}
	
	slashCommand := AddedSlashCommand{
		data: cmd,
		execute: fn,
	}

	addedSlashCommands[cmd.Name] = slashCommand
	return nil
}

func registerSlashCommandsForGuild(guildId string) (int, error) {
	var totalErr error
	var totalCommands int

	for _, addedSlashCommand := range addedSlashCommands {
		session := bot.Session
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, guildId, addedSlashCommand.data)
		if err != nil {
		if totalErr != nil {
				totalErr = err
			} else {
				totalErr = fmt.Errorf("%w\n%w", totalErr, err)
			}
			continue
		}

		guildCommands := registeredGuildsCommands[guildId];
		guildCommands = append(guildCommands, cmd)
		registeredGuildsCommands[guildId] = guildCommands

		totalCommands++
	}

	return totalCommands, totalErr
}

func unregisterSlashCommandsForGuild(guildId string, withApplicationCommand bool) (int, error) {
	var totalErr error
	var totalCommands int

	if withApplicationCommand {
		for _, cmd := range registeredGuildsCommands[guildId] {
			session := bot.Session
			if err := session.ApplicationCommandDelete(session.State.User.ID, guildId, cmd.ID); err != nil {
				if totalErr != nil {
					totalErr = err
				} else {
					totalErr = fmt.Errorf("%w\n%w", totalErr, err)
				}
				continue
			}

			totalCommands++
		}
	} else {
		totalCommands = len(registeredGuildsCommands[guildId])
	}

	delete(registeredGuildsCommands, guildId)
	return totalCommands, totalErr
}


// -
// Internal
// -

func init() {
	bot.AddEventHandler(onInteractionCreate)
	bot.AddEventHandler(onGuildCreate)
	bot.AddEventHandler(onGuildDelete)
	bot.AddEventHandler(onDisconnect)
}

func onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	appCmdData := i.ApplicationCommandData()
	if cmd, exists := addedSlashCommands[appCmdData.Name]; exists {
		var params string
		for _, option := range appCmdData.Options {
			if params == "" {
				params = fmt.Sprintf(" %v", option.Value)
			} else {
				params = fmt.Sprintf("%v %v", params, option.Value)
			}
		}
		utils.Log.Printf("'%v' executed a command: /%v%v", i.Member.User.Username, cmd.data.Name, params)

		if err := cmd.execute(s, i); err != nil {
			utils.Log.Printf("An error occured while executing slash command '%v':\n", cmd.data.Name)
			utils.Log.Println(err)
		}
	}
}

func onDisconnect(s *discordgo.Session, _ *discordgo.Disconnect) {
	utils.Log.Println("Retrieving guilds to unregister slash commands...")
	guilds, err := s.UserGuilds(200, "", "", false)
	if err != nil {
		utils.Log.Println("Failed to retrieve the guilds:", err)
		return
	}

	for _, guild := range guilds {
		utils.Log.Printf("Unregistering commands from guild '%v (ID: %v)' ...\n", guild.Name, guild.ID)
		totalCommands, err := unregisterSlashCommandsForGuild(guild.ID, true)
		if err != nil {
			utils.Log.Printf("Unregistered %v commands from guild '%v (ID: %v)' but with an error:\n", totalCommands, guild.Name, guild.ID)
			utils.Log.Println(err)
		} else {
			utils.Log.Printf("Successfully unregistered %v commands from guild '%v (ID: %v)'", totalCommands, guild.Name, guild.ID)
		}
	}
}

func onGuildCreate(_ *discordgo.Session, guild *discordgo.GuildCreate) {
	utils.Log.Printf("Registering commands to guild '%v (ID: %v)' ...\n", guild.Name, guild.ID)
	totalCommands, err := registerSlashCommandsForGuild(guild.ID)
	if err != nil {
		utils.Log.Printf("Registered %v commands to guild '%v (ID: %v)' but with an error:\n", totalCommands, guild.Name, guild.ID)
		utils.Log.Println(err)
	} else {
		utils.Log.Printf("Successfully registered %v commands to guild '%v (ID: %v)'", totalCommands, guild.Name, guild.ID)
	}
}

func onGuildDelete(_ *discordgo.Session, guild *discordgo.GuildDelete) {
	utils.Log.Printf("Unregistering commands from recently left guild '%v (ID: %v)' ...\n", guild.Name, guild.ID)
	totalCommands, err := unregisterSlashCommandsForGuild(guild.ID, false)
	if err != nil {
		utils.Log.Printf("Unregistered %v commands from recently left guild 'ID: %v' but with an error:\n", totalCommands, guild.ID)
		utils.Log.Println(err)
	} else {
		utils.Log.Printf("Successfully unregistered %v commands from recently left guild 'ID: %v'", totalCommands, guild.ID)
	}
}
