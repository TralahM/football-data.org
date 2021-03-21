package football_data

import (
	"time"
)

type CompetitionsResponse struct {
	Count        int           `json:"count"`
	Filters      Filters       `json:"filters"`
	Competitions []Competition `json:"competitions"`
}
type Area struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Countrycode string `json:"countryCode"`
	Ensignurl   string `json:"ensignUrl"`
}

type Competition struct {
	ID                       int       `json:"id"`
	Area                     Area      `json:"area"`
	Name                     string    `json:"name"`
	Code                     string    `json:"code"`
	Emblemurl                string    `json:"emblemUrl"`
	Plan                     string    `json:"plan"`
	Currentseason            Season    `json:"currentSeason"`
	Numberofavailableseasons int       `json:"numberOfAvailableSeasons"`
	Lastupdated              time.Time `json:"lastUpdated"`
}

type GameResponse struct {
	Head2Head Head2Head `json:"head2head"`
	Match     Match     `json:"match"`
}
type HometeamH2H struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Wins   int    `json:"wins"`
	Draws  int    `json:"draws"`
	Losses int    `json:"losses"`
}
type AwayteamH2H struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Wins   int    `json:"wins"`
	Draws  int    `json:"draws"`
	Losses int    `json:"losses"`
}
type Head2Head struct {
	Numberofmatches int         `json:"numberOfMatches"`
	Totalgoals      int         `json:"totalGoals"`
	Hometeam        HometeamH2H `json:"homeTeam"`
	Awayteam        AwayteamH2H `json:"awayTeam"`
}
type Odds struct {
	Homewin interface{} `json:"homeWin"`
	Draw    interface{} `json:"draw"`
	Awaywin interface{} `json:"awayWin"`
}
type Coach struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Countryofbirth string `json:"countryOfBirth"`
	Nationality    string `json:"nationality"`
}
type Captain struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Shirtnumber int    `json:"shirtNumber"`
}
type Player struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Firstname      string    `json:"firstName"`
	Lastname       string    `json:"lastName"`
	Dateofbirth    string    `json:"dateOfBirth"`
	Countryofbirth string    `json:"countryOfBirth"`
	Nationality    string    `json:"nationality"`
	Position       string    `json:"position"`
	Lastupdated    time.Time `json:"lastUpdated"`
}
type Lineup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Position    string `json:"position"`
	Shirtnumber int    `json:"shirtNumber"`
}
type Bench struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Position    string `json:"position"`
	Shirtnumber int    `json:"shirtNumber"`
}
type Scorer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Assist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Goals struct {
	Minute    int         `json:"minute"`
	Extratime interface{} `json:"extraTime"`
	Type      string      `json:"type"`
	Team      Team        `json:"team"`
	Scorer    Scorer      `json:"scorer"`
	Assist    Assist      `json:"assist"`
}
type Bookings struct {
	Minute int    `json:"minute"`
	Team   Team   `json:"team"`
	Player Player `json:"player"`
	Card   string `json:"card"`
}
type Substitutions struct {
	Minute    int    `json:"minute"`
	Team      Team   `json:"team"`
	Playerout Player `json:"playerOut"`
	Playerin  Player `json:"playerIn"`
}
type Match struct {
	ID            int             `json:"id"`
	Competition   Competition     `json:"competition"`
	Season        Season          `json:"season"`
	Utcdate       time.Time       `json:"utcDate"`
	Status        string          `json:"status"`
	Minute        string          `json:"minute"`
	Attendance    int             `json:"attendance"`
	Venue         string          `json:"venue"`
	Matchday      int             `json:"matchday"`
	Stage         string          `json:"stage"`
	Group         string          `json:"group"`
	Lastupdated   time.Time       `json:"lastUpdated"`
	Odds          Odds            `json:"odds"`
	Score         Score           `json:"score"`
	Hometeam      Hometeam        `json:"homeTeam"`
	Awayteam      Awayteam        `json:"awayTeam"`
	Goals         []Goals         `json:"goals"`
	Bookings      []Bookings      `json:"bookings"`
	Substitutions []Substitutions `json:"substitutions"`
	Referees      []Referee       `json:"referees"`
}

type TeamsResponse struct {
	Count   int     `json:"count"`
	Filters Filters `json:"filters"`
	Teams   []Team  `json:"teams"`
}
type Filters struct {
	Areas      []int  `json:"areas"`
	Permission string `json:"permission"`
}
type Team struct {
	ID          int       `json:"id"`
	Area        Area      `json:"area"`
	Name        string    `json:"name"`
	Shortname   string    `json:"shortName"`
	Tla         string    `json:"tla"`
	Cresturl    string    `json:"crestUrl"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	Website     string    `json:"website"`
	Email       string    `json:"email"`
	Founded     int       `json:"founded"`
	Clubcolors  string    `json:"clubColors"`
	Venue       string    `json:"venue"`
	Lastupdated time.Time `json:"lastUpdated"`
}
type StandingsResponse struct {
	Filters     Filters     `json:"filters"`
	Competition Competition `json:"competition"`
	Season      Season      `json:"season"`
	Standings   []Standing  `json:"standings"`
}
type Season struct {
	ID              int         `json:"id"`
	Startdate       string      `json:"startDate"`
	Enddate         string      `json:"endDate"`
	Currentmatchday int         `json:"currentMatchday"`
	Winner          interface{} `json:"winner"`
	Availablestages []string    `json:"availableStages"`
}
type Table struct {
	Position       int    `json:"position"`
	Team           Team   `json:"team"`
	Playedgames    int    `json:"playedGames"`
	Form           string `json:"form"`
	Won            int    `json:"won"`
	Draw           int    `json:"draw"`
	Lost           int    `json:"lost"`
	Points         int    `json:"points"`
	Goalsfor       int    `json:"goalsFor"`
	Goalsagainst   int    `json:"goalsAgainst"`
	Goaldifference int    `json:"goalDifference"`
}
type Standing struct {
	Stage string      `json:"stage"`
	Type  string      `json:"type"`
	Group interface{} `json:"group"`
	Table []Table     `json:"table"`
}

type MatchesResponse struct {
	Count       int         `json:"count"`
	Filters     Filters     `json:"filters"`
	Competition Competition `json:"competition"`
	Matches     []Match     `json:"matches"`
}
type Fulltime struct {
	Hometeam int `json:"homeTeam"`
	Awayteam int `json:"awayTeam"`
}
type Halftime struct {
	Hometeam int `json:"homeTeam"`
	Awayteam int `json:"awayTeam"`
}
type Extratime struct {
	Hometeam int `json:"homeTeam"`
	Awayteam int `json:"awayTeam"`
}
type Penalties struct {
	Hometeam interface{} `json:"homeTeam"`
	Awayteam interface{} `json:"awayTeam"`
}
type Score struct {
	Winner    string    `json:"winner"`
	Duration  string    `json:"duration"`
	Fulltime  Fulltime  `json:"fullTime"`
	Halftime  Halftime  `json:"halfTime"`
	Extratime Extratime `json:"extraTime"`
	Penalties Penalties `json:"penalties"`
}
type Hometeam struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Coach   Coach    `json:"coach"`
	Captain Captain  `json:"captain"`
	Lineup  []Lineup `json:"lineup"`
	Bench   []Bench  `json:"bench"`
}
type Awayteam struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Coach   Coach    `json:"coach"`
	Captain Captain  `json:"captain"`
	Lineup  []Lineup `json:"lineup"`
	Bench   []Bench  `json:"bench"`
}
type Referee struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Nationality string `json:"nationality"`
}
