package models

type Match struct {
	ID          int
	HomeClub    *Club
	AwayClub    *Club
	HomeScore   int
	AwayScore   int
	Date        time.Time
	MatchEvents []MatchEvent // Opsional
}
