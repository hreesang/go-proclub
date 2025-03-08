package proclubs

type MatchType = int

const (
	MatchTypeLeague 	MatchType = iota
	MatchTypePlayoff
)

var matchStrings = map[MatchType]string {
	MatchTypeLeague:	"leagueMatch",
	MatchTypePlayoff:	"playoffMatch",
}

var matchNames = map[MatchType]string {
	MatchTypeLeague:	"League",
	MatchTypePlayoff:	"Playoff",
}

func MatchTypeString(m MatchType) string {
	return matchStrings[m]
}

func StringToMatchType(str string) MatchType {
	for i, matchString := range matchStrings {
		if str == matchString {
			return i
		}
	}
	return MatchTypeLeague
}

func MatchTypeName(m MatchType) string {
	return matchNames[m]
}
