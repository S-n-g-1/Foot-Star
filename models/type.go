package models

type Foot string

const (
	RightFoot Foot = "right"
	LeftFoot  Foot = "left"
)

type Position string

const (
	Defender   Position = "defender"
	Attacker   Position = "attacker"
	Midfielder Position = "midfielder"
	Goalkeeper Position = "goalkeeper"
)

type National string

const (
	Indonesia National = "Indonesia"
	Belanda   National = "Belanda"
	Portugal  National = "Portugal"
	Jerman    National = "Jerman"
	Brazil    National = "Brazil"
	Inggris   National = "Inggris"
	Italia    National = "Italia"
)

type Skill string

const (
	Save           Skill = "Save"
	Reflexes       Skill = "Reflexes"
	Anticipation   Skill = "Anticipation"
	AerialAbility  Skill = "Aerial ability"
	Positioning    Skill = "Positioning"
	Tackling       Skill = "Tackling"
	Marking        Skill = "Marking"
	BallRecovery   Skill = "Ball recovery"
	Shoot          Skill = "Shoot"
	Heading        Skill = "Heading"
	Strength       Skill = "Strength"
	Physique       Skill = "Physique"
	Speed          Skill = "Speed"
	Agility        Skill = "Agility"
	Creativity     Skill = "Creativity"
	Passes         Skill = "Passes"
	Endurance      Skill = "Endurance"
	FreeKicks      Skill = "Free kicks"
	Dribbling      Skill = "Dribbling"
	Technique      Skill = "Technique"
	Leadership     Skill = "Leadership"
	Aggressiveness Skill = "Aggressiveness"
	Intelligence   Skill = "Intelligence"
	Control        Skill = "Control"
	Mobility       Skill = "Mobility"
)
