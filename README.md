[![Build Status](https://travis-ci.com/TralahM/football-data.org.svg?branch=master)](https://travis-ci.com/TralahM/football-data.org)
[![License: GPLv3](https://img.shields.io/badge/License-GPLV2-green.svg)](https://opensource.org/licenses/GPLV2)
[![Organization](https://img.shields.io/badge/Org-TralahTek-blue.svg)](https://github.com/TralahTek)
[![PRs Welcome](https://img.shields.io/badge/PRs-Welcome-brightgreen.svg?style=flat-square)](https://github.com/TralahM/football-data.org/pull/)
[![GitHub pull-requests](https://img.shields.io/badge/Issues-pr-red.svg?style=flat-square)](https://github.com/TralahM/football-data.org/pull/)
[![Language](https://img.shields.io/badge/Language-go-00ADD8.svg)](https://github.com/TralahM)
<img title="Watching" src="https://img.shields.io/github/watchers/TralahM/football-data.org?label=Watchers&color=blue&style=flat-square">
<img title="Stars" src="https://img.shields.io/github/stars/TralahM/football-data.org?color=red&style=flat-square">
<img title="Forks" src="https://img.shields.io/github/forks/TralahM/football-data.org?color=green&style=flat-square">
<noscript><a href="https://liberapay.com/TralahM/donate"><img alt="Donate using Liberapay" src="https://liberapay.com/assets/widgets/donate.svg"></a></noscript>

# football-data.org.
A Go Client for the [Football-data.org](football-data.org) API.


# Installation
```console
# In terminal do:
$ go get -u github.com/tralahm/football-data.org
```

## Usage
```go
package main

import (
	"fmt"

	football_data "github.com/tralahm/football-data.org"
)

func main() {

	XAuthToken := "your football-data api key"

	apiClient := football_data.NewAPIClient(XAuthToken)

	leaguesResponse := apiClient.GetLeagues()
	fmt.Println(leaguesResponse.Competitions[0])
	LeagueID := "2018"
	competition := apiClient.GetLeague(LeagueID)
	fmt.Println(competition)
	gamesResponse := apiClient.GetGames(LeagueID)
	fmt.Println(gamesResponse.Matches[0])
	teamsResponse := apiClient.GetTeams(LeagueID)
	fmt.Println(teamsResponse.Teams[0])
	GameID := "239019"
	gameR := apiClient.GetGameDetails(GameID)
	fmt.Println(gameR.Head2Head, gameR.Match)
	standingsResponse := apiClient.GetLeague(LeagueID)
	fmt.Println(standingsResponse.Standings[0])
}

```


# LICENCE

[Read the license here](LICENSE)


# Self-Promotion

[![](https://img.shields.io/badge/Github-TralahM-green?style=for-the-badge&logo=github)](https://github.com/TralahM)
[![](https://img.shields.io/badge/Twitter-%40tralahtek-blue?style=for-the-badge&logo=twitter)](https://twitter.com/TralahM)
[![TralahM](https://img.shields.io/badge/Kaggle-TralahM-purple.svg?style=for-the-badge&logo=kaggle)](https://kaggle.com/TralahM)
[![TralahM](https://img.shields.io/badge/LinkedIn-TralahM-white.svg?style=for-the-badge&logo=linkedin)](https://linkedin.com/in/TralahM)


[![Blog](https://img.shields.io/badge/Blog-tralahm.tralahtek.com-blue.svg?style=for-the-badge&logo=rss)](https://tralahm.tralahtek.com)

[![TralahTek](https://img.shields.io/badge/Organization-TralahTek-cyan.svg?style=for-the-badge)](https://org.tralahtek.com)


