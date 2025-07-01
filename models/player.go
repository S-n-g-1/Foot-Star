package models

type Player struct {
	ID       int
	Name     string
	Skills   map[Skill]int8
	National National
	Age      int8
	Height   int8
	Star     int8
	Foot     Foot
	Position Position
	Value    int32
	Overall  int8
	Club     *Club // relasi balik ke klub
}
