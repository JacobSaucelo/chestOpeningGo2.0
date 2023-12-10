package main

import "fmt"

var caseItem = []Item{
	{
		Name:   "----------------------------M4A4 | Eye of Horus",
		Rarity: Red,
	},
	{
		Name:   "FAMAS | Waters of Nephthys",
		Rarity: Pink,
	},
	{
		Name:   "P250 | Apep's Curse",
		Rarity: Pink,
	},
	{
		Name:   "Glock-18 | Ramese's Reach",
		Rarity: Pink,
	},
	{
		Name:   "Nova | Sobek's Bite",
		Rarity: Pink,
	},
	{
		Name:   "AK-47 | Steel Delta",
		Rarity: Pink,
	},
	{
		Name:   "AWP | Black Nile",
		Rarity: Pink,
	},
	{
		Name:   "P90 | ScaraB Rush",
		Rarity: Pink,
	},
	{
		Name:   "M4A1-S | Mud-Spec",
		Rarity: Pink,
	},
	{
		Name:   "MAG-7 | Copper Coated",
		Rarity: Pink,
	},
	{
		Name:   "Tec-9 | Mummy's Rot",
		Rarity: Purple,
	},
	{
		Name:   "MAC-10 | Echoing Sands",
		Rarity: Purple,
	},
	{
		Name:   "SSG 08 | Azure Glyph",
		Rarity: Purple,
	},
	{
		Name:   "USP-S | Desert Tactical",
		Rarity: Purple,
	},
	{
		Name:   "AUG | Snake Pit",
		Rarity: Blue,
	},
	{
		Name:   "MP7 | Sunbaked",
		Rarity: Blue,
	},
	{
		Name:   "XM1014 | Hieroglyph",
		Rarity: Blue,
	},
	{
		Name:   "M249 | Submerged",
		Rarity: Blue,
	},
	{
		Name:   "R8 Revolver | Inlay",
		Rarity: Blue,
	},
}

func main() {

	anubisCase := NewChest(caseItem)
	fmt.Println(anubisCase.OpenChest())
	// temp := 10.0
	// tries := 0

	// for temp > 0.63940 {
	// 	item := anubisCase.OpenChest()
	// 	temp = float64(item.Rarity)
	// 	fmt.Println("tries: ", tries, item)
	// 	tries += 1
	// }
}
