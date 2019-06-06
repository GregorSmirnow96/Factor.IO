package items

// Item describes the information of an item in Factorio.
type Item struct {
	Name          string
	Recipe        map[string]int
	CraftSpeed    float32
	YieldPerCraft int
}
