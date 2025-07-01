package models

type League struct {
	ID         int
	Name       string
	Country    string
	Tier       int8    // divisi: 1 (atas), 2 (kedua), dst
	Reputation int8    // 1–100
}
