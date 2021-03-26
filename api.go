package football_data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type APIClient struct {
	client Client
}

func DoGet(addr string, apiClient Client) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, addr, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	res, err := apiClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer res.Body.Close()
	log.Println("statusCode: ", res.Status)
	if res.Status == 429 {
		sleep := res.Header.Get("X-RequestCounter-Reset").(int)
		log.Fatalf("X-RequestCounter-Reset: %v", sleep)
		time.Sleep(time.Second * sleep)
		return DoGet(addr, apiClient)
	}
	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading HTTP response: %s", err)
	}
	return responseBytes, err

}
func (c APIClient) GetLeagues() CompetitionsResponse {
	endpoint := "/v2/competitions"
	addr := baseUrl + endpoint
	responseBytes, err := DoGet(addr, c.client)
	if err != nil {
		log.Fatalf("Error geting Leagues: %s", err)
	}
	// log.Println(string(responseBytes))
	var competitions CompetitionsResponse
	if err := json.Unmarshal(responseBytes, &competitions); err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n\n", competitions.Competitions[0])
	// return string(responseBytes)
	return competitions
}

func (c APIClient) GetLeague(leagueCode string) Competition {
	endpoint := "/v2/competitions/" + leagueCode
	addr := baseUrl + endpoint
	responseBytes, err := DoGet(addr, c.client)
	if err != nil {
		log.Fatalf("Error geting League %s: %s", leagueCode, err)
	}
	// log.Println(string(responseBytes))
	var competition Competition
	if err := json.Unmarshal(responseBytes, &competition); err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n\n", competition)
	// return string(responseBytes)
	return competition
}

func (c APIClient) GetTeams(leagueCode string) TeamsResponse {
	endpoint := "/v2/competitions/" + leagueCode + "/teams"
	addr := baseUrl + endpoint
	responseBytes, err := DoGet(addr, c.client)
	if err != nil {
		log.Fatalf("Error getting %s Teams: %s", leagueCode, err)
	}
	// log.Println(string(responseBytes))
	var teams TeamsResponse
	if err := json.Unmarshal(responseBytes, &teams); err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n\n", teams.Teams[0])
	return teams
}

func (c APIClient) GetGames(leagueCode string) MatchesResponse {
	endpoint := "/v2/competitions/" + leagueCode + "/matches"
	addr := baseUrl + endpoint
	responseBytes, err := DoGet(addr, c.client)
	if err != nil {
		log.Fatalf("Error getting %s Matches: %s", leagueCode, err)
	}
	// log.Println(string(responseBytes))
	var matches MatchesResponse
	if err := json.Unmarshal(responseBytes, &matches); err != nil {
		log.Fatal(err)
	}
	return matches
}

func (c APIClient) GetGameDetails(gameId string) GameResponse {
	endpoint := "/v2/matches/" + gameId
	addr := baseUrl + endpoint
	responseBytes, err := DoGet(addr, c.client)
	if err != nil {
		log.Fatalf("Error geting Game %s: %s", gameId, err)
	}
	// log.Println(string(responseBytes))
	var match GameResponse
	if err := json.Unmarshal(responseBytes, &match); err != nil {
		log.Fatal(err)
	}
	return match
}

func (c APIClient) GetLeagueStandings(leagueId string) StandingsResponse {
	endpoint := "/v2/competitions/" + leagueId + "/standings"
	addr := baseUrl + endpoint
	var standings StandingsResponse
	responseBytes, err := DoGet(addr, c.client)
	if err != nil {
		log.Fatalf("Error geting Standings %s: %s", leagueId, err)
	}
	if err := json.Unmarshal(responseBytes, &standings); err != nil {
		log.Fatal(err)
	}
	return standings
}
