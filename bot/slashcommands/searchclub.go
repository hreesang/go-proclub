package slashcommands

import (
	"fmt"
	"sort"

	dg "github.com/bwmarrin/discordgo"
	"github.com/hreesang/go-proclub/bot/proclubs"
)

func init() {
	platformChoices := make([]*dg.ApplicationCommandOptionChoice, 0)
	for i := range proclubs.MaxPlatforms {
		platformChoices = append(platformChoices,
			&dg.ApplicationCommandOptionChoice{
				Name: 	proclubs.PlatformName(i),
				Value: 	i,
			},
		)
	}

	AddSlashCommand(&dg.ApplicationCommand{
		Name: 			"searchclub",
		Description: 	"Searches the specified club and its platform.",
		Options: 		[]*dg.ApplicationCommandOption{
			{
				Required: 		true,
				Name: 			"name",
				Description:	"The name of the club to search for",
				Type: 			dg.ApplicationCommandOptionString,
			},
			{
				Required: 		true,
				Name: 			"platform",
				Type: 			dg.ApplicationCommandOptionInteger,
				Choices: 		platformChoices,
				Description:	"The platform on which the club is located",
			},
		},
	}, searchclub)
}

func searchclub(s *dg.Session, i *dg.InteractionCreate) error {
	s.InteractionRespond(i.Interaction, &dg.InteractionResponse{
		Type: dg.InteractionResponseDeferredChannelMessageWithSource,
	})

	options := i.ApplicationCommandData().Options

	clubName := options[0].StringValue()
	platform := proclubs.Platform(options[1].IntValue())

	clubs, err := proclubs.SearchClub(clubName, platform)
	if err != nil || len(clubs) == 0 {
		errorMessage := fmt.Sprintf("There is no club with the name '%v' on the %v platform.", clubName, proclubs.PlatformName(platform)) 
		if _, interactionErr := s.InteractionResponseEdit(i.Interaction, &dg.WebhookEdit{
			Content: &errorMessage,
		}); interactionErr != nil {
			return interactionErr
		}
		return err
	}
	club := clubs[0]

	clubOverallStats, err := proclubs.GetClubOverallStats(club.ClubId, platform); if err != nil {
		errorMessage := "An error occured while retrieving the club information."
		if _, interactionErr := s.InteractionResponseEdit(i.Interaction, &dg.WebhookEdit{
			Content: &errorMessage,
		}); interactionErr != nil {
			return interactionErr
		}
		return err
	}
	
	leagueMatchesStats, err := proclubs.GetMatchesStatsFromClubId(club.ClubId, platform, proclubs.MatchTypeLeague, 3)
	if err != nil {
		errorMessasge := "An error occured, unable to retrieve the club's last matches."
		if _, interactionErr := s.InteractionResponseEdit(i.Interaction, &dg.WebhookEdit{
			Content: &errorMessasge,
		}); interactionErr != nil {
			return interactionErr
		}
		return err
	}
	
	playoffMatchesStats, err := proclubs.GetMatchesStatsFromClubId(club.ClubId, platform, proclubs.MatchTypePlayoff, 3)
	if err != nil {
		errorMessasge := "An error occured, unable to retrieve the club's last matches."
		if _, interactionErr := s.InteractionResponseEdit(i.Interaction, &dg.WebhookEdit{
			Content: &errorMessasge,
		}); interactionErr != nil {
			return interactionErr
		}
		return err
	}

	matchesStats := append(playoffMatchesStats, leagueMatchesStats...)
	sort.Slice(matchesStats, func (i, j int) bool {
		return matchesStats[i].Timestamp < matchesStats[j].Timestamp
	})
	matchesStats = matchesStats[:3]

	if _, err := s.InteractionResponseEdit(i.Interaction, &dg.WebhookEdit{
		Embeds: &[]*dg.MessageEmbed{
			clubMessageEmbed(club, clubOverallStats, matchesStats),
		},
	}); err != nil {
		return err
	}

	return nil
}

func clubMessageEmbed(c *proclubs.Club, ovr *proclubs.ClubOverallStats, ms []*proclubs.MatchStats) *dg.MessageEmbed {
	var platformName string
	if platform, err := proclubs.StringToPlatform(c.Platform); err != nil {
		platformName = "Unknown"
	} else {
		platformName = proclubs.PlatformName(platform)
	}
	
	var lastMatchesContent string
	for _, m := range ms {
		var opponentId string

		for clubIds := range m.Clubs {
			if clubIds == c.ClubId {
				continue
			}

			opponentId = clubIds
			break
		}

		club := m.Clubs[c.ClubId]
		opponentClub := m.Clubs[opponentId]
		
		lastMatchesContent = fmt.Sprintf(
			"%v(%v) %v - %v vs. %v\n",
			lastMatchesContent,
			proclubs.MatchTypeName(proclubs.StringToMatchType(m.MatchType)),
			club.Score,
			opponentClub.Score,
			opponentClub.Details.Name,
		)
	}
	
	return &dg.MessageEmbed{
		Color: 		0x07f468,
		Author: 	&dg.MessageEmbedAuthor{
			Name: 		"Club Info",
		},
		Thumbnail: 	&dg.MessageEmbedThumbnail{
			URL: 		proclubs.TeamCrestURL(c.ClubInfo.TeamId),
		},
		Title: 			fmt.Sprintf("%v (ID: %v)", c.ClubName, c.ClubId),
		Description: 	fmt.Sprintf("%v is an EA Sports FC Pro Club hosted in %v and currently playing in Division %v.", c.ClubName, c.ClubInfo.CustomKit.StadName, c.CurrentDivision),
		Fields: 	[]*dg.MessageEmbedField{
			{
				Name:		"Platform",
				Value:		platformName,
				Inline:		true,
			},
			{
				Name:		"Stadium",
				Value:		c.ClubInfo.CustomKit.StadName,
				Inline:		true,
			},
			{
				Name:		"Best Division",
				Value: 		"Divison " + ovr.BestDivision,
				Inline: 	true,
			},
			{
				Name:		"Games",
				Value: 		fmt.Sprintf("%v Matches Played\n%vW %vL %vD\n%v Unbeaten Streak", ovr.GamesPlayed, ovr.Wins, ovr.Losses, ovr.Ties, ovr.UnbeatenStreak),
				Inline: 	true,
			},
			{
				Name:		"Club Stats",
				Value: 		fmt.Sprintf("%v Goals Scored\n%v Goals Against\n%v Cleansheets", ovr.Goals, ovr.GoalsAgainst, c.CleanSheets),
				Inline: 	true,
			},
			{
				Name: 		"Last 3 Matches",
				Value:		lastMatchesContent,
				Inline: 	true,
			},
		},
	}
}
