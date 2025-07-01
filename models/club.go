package models

type Club struct {
	ID        int
	Name      string
	Stadium   string
	Founded   int
	League    *League   // relasi ke League (struct dari league.go)
	Budget    int32
	Players   []*Player // relasi ke Player (struct dari player.go)
	CoachName string
	Reputation int8
}
