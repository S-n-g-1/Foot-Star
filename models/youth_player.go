package modles

type YouthPlayer struct {
	PlayerData Player
	Ready      bool // jika true, bisa dipromosikan
	TrainingProgress int
}
