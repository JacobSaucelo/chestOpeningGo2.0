package main

type Chest struct {
	items []Item
}

func (c *Chest) NewChest(i []Item) Chest {
	return Chest{
		items: i,
	}
}
