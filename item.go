package main

type Rarity float64

const (
	Yellow Rarity = 0.25575
	Red    Rarity = 0.63940
	Pink   Rarity = 3.19693
	Purple Rarity = 15.98465
	Blue   Rarity = 79.92327
)

type Hat struct {
	Name   string
	Rarity Rarity
}
