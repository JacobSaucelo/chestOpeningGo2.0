package main

import (
	"math/rand"
)

type Chest struct {
	items []Item
}

func NewChest(i []Item) Chest {
	return Chest{
		items: i,
	}
}

func (c *Chest) AddItem() {

}

func (c *Chest) OpenChest() Item { //? this requires sorted percentage
	totalPercentage := 0.0
	for _, item := range c.items {
		totalPercentage += float64(item.Rarity)
	}

	randomNum := rand.Float64() * 100
	cumulativePerc := 0.0
	item := Item{}

	for _, stat := range c.items {
		cumulativePerc += (float64(stat.Rarity) / totalPercentage) * 100
		if randomNum <= cumulativePerc {
			// fmt.Println("rando: ", randomNum)
			// fmt.Println("stats: ", cumulativePerc)
			item = stat
			break
		}
	}

	return item
}
