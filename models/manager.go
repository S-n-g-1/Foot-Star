package models

type Manager struct {
	Name        string
	National	*National
	Age         int
	Club        *Club
	Balance     int64
}
