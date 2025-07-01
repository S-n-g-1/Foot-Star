package modles

type MatchEvent struct {
	Minute  int
	Player  *Player
	Event   string // "goal", "yellow_card", "red_card", "substitution", dll
	Team    *Club
}
