package proclubs

type Club struct {
	ClubId             string   `json:"clubId"`
	Wins               string   `json:"wins"`
	Loses              string   `json:"loses"`
	Ties               string   `json:"ties"`
	GamesPlayed        string   `json:"gamesPlayed"`
	GamesPlayedPlayoff string   `json:"gamesPlayedPlayoff"`
	Goals              string   `json:"goals"`
	GoalsAgainst       string   `json:"goalsAgainst"`
	CleanSheets        string   `json:"cleanSheets"`
	Points             string   `json:"points"`
	ReputationTier     string   `json:"reputationTier"`
	ClubInfo           ClubInfo `json:"clubInfo"`
	Platform           string   `json:"platform"`
	ClubName           string   `json:"clubName"`
	CurrentDivision    string   `json:"currentDivision"`
}

type ClubInfo struct {
	Name           string 				`json:"name"`
	ClubId         int    				`json:"clubId"`
	RegionId       int    				`json:"regionId"`
	TeamId         int    				`json:"teamId"`
	CustomKit      ClubInfoCustomKit	`json:"customKit"`
}

type ClubInfoCustomKit struct {
	StadName        	string `json:"stadName"`
	KitId           	string `json:"kitId"`
	SeasonalTeamId  	string `json:"seasonalTeamId"`
	SeasonalKitId   	string `json:"seasonalKitId"`
	SelectedKitType 	string `json:"selectedKitType"`
	CustomKitId     	string `json:"customKitId"`
	CustomAwayKitId 	string `json:"customAwayKitId"`
	CustomThirdKitId 	string `json:"customThirdKitId"`
	CustomKeeperKitId 	string `json:"customKeeperKitId"`
	KitColor1       	string `json:"kitColor1"`
	KitColor2       	string `json:"kitColor2"`
	KitColor3       	string `json:"kitColor3"`
	KitColor4       	string `json:"kitColor4"`
	KitAColor1      	string `json:"kitAColor1"`
	KitAColor2      	string `json:"kitAColor2"`
	KitAColor3      	string `json:"kitAColor3"`
	KitAColor4      	string `json:"kitAColor4"`
	KitThrdColor1   	string `json:"kitThrdColor1"`
	KitThrdColor2   	string `json:"kitThrdColor2"`
	KitThrdColor3   	string `json:"kitThrdColor3"`
	KitThrdColor4   	string `json:"kitThrdColor4"`
	DCustomKit      	string `json:"dCustomKit"`
	CrestColor      	string `json:"crestColor"`
	CrestAssetId    	string `json:"crestAssetId"`
}

type MatchStats struct {
	MatchId   	string 							`json:"matchId"`
	MatchType	string
	Timestamp 	int64  							`json:"timestamp"`
	TimeAgo   	MatchStatsTimeAgo				`json:"timeAgo"`
	Clubs 		map[string]MatchStatsClub 		`json:"clubs"`
	Players 	map[string]MatchStatsPlayers	`json:"players"`
	Aggregate 	map[string]MatchStastAggregate	`json:"aggregate"`
}

type MatchStatsTimeAgo struct {
	Number	int    `json:"number"`
	Unit   	string `json:"unit"`
}

type MatchStatsClub struct {
	Date          string  `json:"date"`
	GameNumber    string  `json:"gameNumber"`
	Goals         string  `json:"goals"`
	GoalsAgainst  string  `json:"goalsAgainst"`
	Losses        string  `json:"losses"`
	MatchType     string  `json:"matchType"`
	Result        string  `json:"result"`
	Score         string  `json:"score"`
	SeasonId      string  `json:"season_id"`
	Team          string  `json:"TEAM"`
	Ties          string  `json:"ties"`
	WinnerByDnf   string  `json:"winnerByDnf"`
	Wins          string  `json:"wins"`
	Details       ClubInfo `json:"details"`
}

type MatchStatsPlayers map[string]MatchStatsPlayer

type MatchStatsPlayer struct {
	Assists          string `json:"assists"`
	CleanSheetsAny   string `json:"cleansheetsany"`
	CleanSheetsDef   string `json:"cleansheetsdef"`
	CleanSheetsGk    string `json:"cleansheetsgk"`
	Goals            string `json:"goals"`
	GoalsConceded    string `json:"goalsconceded"`
	Losses           string `json:"losses"`
	Mom              string `json:"mom"`
	Namespace        string `json:"namespace"`
	PassAttempts     string `json:"passattempts"`
	PassesMade       string `json:"passesmade"`
	Pos              string `json:"pos"`
	Rating           string `json:"rating"`
	RealTimeGame     string `json:"realtimegame"`
	RealTimeIdle     string `json:"realtimeidle"`
	RedCards         string `json:"redcards"`
	Saves            string `json:"saves"`
	Score            string `json:"SCORE"`
	Shots            string `json:"shots"`
	TackleAttempts   string `json:"tackleattempts"`
	TacklesMade      string `json:"tacklesmade"`
	VproAttr         string `json:"vproattr"`
	VproHackReason   string `json:"vprohackreason"`
	Wins             string `json:"wins"`
	PlayerName       string `json:"playername"`
} 

type MatchStastAggregate struct {
	Assists          int 		`json:"assists"`
	CleanSheetsAny   int 		`json:"cleansheetsany"`
	CleanSheetsDef   int 		`json:"cleansheetsdef"`
	CleanSheetsGk    int 		`json:"cleansheetsgk"`
	Goals            int 		`json:"goals"`
	GoalsConceded    int 		`json:"goalsconceded"`
	Losses           int 		`json:"losses"`
	Mom              int 		`json:"mom"`
	Namespace        int 		`json:"namespace"`
	PassAttempts     int 		`json:"passattempts"`
	PassesMade       int 		`json:"passesmade"`
	Pos              int 		`json:"pos"`
	Rating           float32 	`json:"rating"`
	RealTimeGame     int 		`json:"realtimegame"`
	RealTimeIdle     int 		`json:"realtimeidle"`
	RedCards         int 		`json:"redcards"`
	Saves            int 		`json:"saves"`
	Score            int 		`json:"SCORE"`
	Shots            int 		`json:"shots"`
	TackleAttempts   int 		`json:"tackleattempts"`
	TacklesMade      int 		`json:"tacklesmade"`
	VproAttr         int 		`json:"vproattr"`
	VproHackReason   int 		`json:"vprohackreason"`
	Wins             int 		`json:"wins"`
} 
