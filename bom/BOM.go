package bom

import (
	"../items"
)

// The {B}ill {o}f {M}aterials of an item.
type bom struct {
	quantities map[string]int
}

// CalculateBOM finds the total number of materials required
// to craft the specified item.
func CalculateBOM(itemName string) map[string]int {
	theBOM := bom{quantities: make(map[string]int)}
	theBOM.findBOMRecursively(itemName)
	return theBOM.quantities
}

// A recursive function to find the total number of each item
// used in crafting an item.
func (theBOM *bom) findBOMRecursively(itemName string) {
	itemRepo := items.GetItemRepo()
	itemOfInterest := itemRepo.Lookup(itemName)

	for recipeItemName, itemQuantity := range itemOfInterest.Recipe {
		recipeItem := itemRepo.Lookup(recipeItemName)

		for i := 0; i < itemQuantity; i++ {
			theBOM.findBOMRecursively(recipeItem.Name)
		}
	}
	theBOM.addItem(itemName)
}

// Increments the specified item's amount by 1.
func (theBOM *bom) addItem(itemName string) {
	if value, ok := theBOM.quantities[itemName]; ok {
		newValue := value + 1
		theBOM.quantities[itemName] = newValue
	} else {
		theBOM.quantities[itemName] = 1
	}
}
