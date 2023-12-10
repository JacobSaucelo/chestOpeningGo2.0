package main

import (
	"fmt"
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

//sort the chest by rarity on create chest

func (c *Chest) OpenChest() {
	// percentage / sum of items
	totalPercentage := 0.0

	for _, item := range c.items {
		// fmt.Println("item: ", item)
		totalPercentage += float64(item.Rarity)
	}
	fmt.Println("totalPercentage: ", totalPercentage)

	randomNum := rand.Float64() * totalPercentage

	cumulativePerc := 0.0
	// chestStat := (cumulativePerc / totalPercentage) * 100

	for _, stat := range c.items {
		cumulativePerc += float64(stat.Rarity)

		// fmt.Println("cumulativePerc: ", cumulativePerc)
		// fmt.Println("chestStat: ", chestStat)
		if randomNum <= cumulativePerc {
			fmt.Println("------------------------------------")
			fmt.Println(stat)
			// fmt.Println("rando: ", randomNum)
			// fmt.Println("stats: ", cumulativePerc)
			break

		}
	}

}
