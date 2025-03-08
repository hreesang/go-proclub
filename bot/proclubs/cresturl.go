package proclubs

import "fmt"

var (
	crestBaseUrl = "https://eafc25.content.easports.com/fifa/fltOnlineAssets/24B23FDE-7835-41C2-87A2-F453DFDB2E82/2024/fcweb/crests/256x256/"
)

func TeamCrestURL(teamId int) string {
	finalUrl := fmt.Sprintf("%vl%v.png", crestBaseUrl, teamId)
	return finalUrl
}
