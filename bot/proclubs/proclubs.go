package proclubs

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var (
	client	= &http.Client{
		Transport: &http.Transport{
			ForceAttemptHTTP2: false,
		},
	}
	baseUrl = "https://proclubs.ea.com/api/fc/"
	method	= "GET"
	headers  = http.Header{
		"User-Agent": []string{"Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1"},
		"Accept": []string{
			"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		},
		"Accept-Encoding": []string{"gzip", "deflate", "br", "zstd"},
		"Accept-Language": []string{"en-US,en;q=0.9"},
		"Cache-Control":   []string{"max-age=0"},
		"Upgrade-Insecure-Requests": []string{"1"},
		"Sec-Fetch-Dest":  []string{"document"},
		"Sec-Fetch-Mode":  []string{"navigate"},
		"Sec-Fetch-Site":  []string{"none"},
		"Sec-Fetch-User":  []string{"?1"},
	}
)

type route int

const (
	routeClubSearch		route = iota
	routeMatchesStats
	routeClubInfo
	routeClubOverallStats
)

var routeEndpoints = map[route]string{
	routeClubSearch:		"allTimeLeaderboard/search",
	routeMatchesStats:		"clubs/matches",
	routeClubInfo:			"clubs/info",
	routeClubOverallStats:	"clubs/overallStats",
}

func getEndpointResponseData(r route, searchParams *url.Values, target any) (error) {
	finalUrl := baseUrl + routeEndpoints[r] + "?" + searchParams.Encode()
	request, err := http.NewRequest(method, finalUrl, nil)
	if err != nil {
		return err
	}
	request.Header = headers
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var reader io.Reader
	if strings.Contains(response.Header.Get("Content-Encoding"), "gzip") {
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return err
		}
		defer reader.(*gzip.Reader).Close()
	} else {
		reader = response.Body
	}	

	return json.NewDecoder(reader).Decode(target)
}

func SearchClub(clubName string, platform Platform) ([]*Club, error) {
	var clubs []*Club

	searchParams := url.Values{
		"clubName": []string{clubName},
		"platform": []string{PlatformString(platform)},
	}

	if err := getEndpointResponseData(routeClubSearch, &searchParams, &clubs); err != nil {
		return clubs, err
	}
	
	return clubs, nil
}

func GetMatchesStatsFromClubId(clubId string, platform Platform, matchType MatchType, count int) ([]*MatchStats, error) {
	var matchesStats []*MatchStats
	
	searchParams := url.Values{
		"clubIds": []string{clubId},
		"platform": []string{PlatformString(platform)},
		"matchType": []string{MatchTypeString(matchType)},
		"maxResultCount": []string{strconv.Itoa(count)},
	}

	if err := getEndpointResponseData(routeMatchesStats, &searchParams, &matchesStats); err != nil {
		return matchesStats, err
	}

	for _, matchStats := range matchesStats {
		matchStats.MatchType = MatchTypeString(matchType)
	}
	
	return matchesStats, nil
}

func GetClubInfo(clubId string, platform Platform) (*ClubInfo, error) {
	var clubInfo *ClubInfo
	
	searchParams := url.Values{
		"clubIds": []string{clubId},
		"platform": []string{PlatformString(platform)},
	}

	if err := getEndpointResponseData(routeClubInfo, &searchParams, &clubInfo); err != nil {
		return clubInfo, err
	}
	
	return clubInfo, nil
}

func GetClubOverallStats(clubId string, platform Platform) (*ClubOverallStats, error) {
	var clubsOverallStats []*ClubOverallStats
	
	searchParams := url.Values{
		"clubIds": []string{clubId},
		"platform": []string{PlatformString(platform)},
	}

	if err := getEndpointResponseData(routeClubOverallStats, &searchParams, &clubsOverallStats); err != nil {
		return clubsOverallStats[0], err
	}
	
	return clubsOverallStats[0], nil
}