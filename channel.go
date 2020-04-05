package main

const MinLevel = 1
const MaxLevel = 7

type Channel struct {
	Enabled bool

	Channel int
	Bus 	int

	Level 	int
	Active  bool // Mute
}
