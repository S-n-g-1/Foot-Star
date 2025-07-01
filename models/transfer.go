package modles

type Transfer struct {
	Player     *Player
	FromClub   *Club
	ToClub     *Club
	Price      int32
	TransferDate time.Time
}
