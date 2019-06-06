package ratios

import (
	"../items"
)

// PROOF OF CONCEPT: Make the map a custom made object that
// implements retrieval ignoring specified keys.
// For now, use a static list and hard code it in FindSmallestDecimalComponent().
var RELAX = []string{}

// FindFloatAssemblerRatios finds the perfect ratio of assemblers
// needed to craft an item provided its BOM.
func FindFloatAssemblerRatios(
	itemBOM map[string]int) map[string]float32 {

	itemRepo := items.GetItemRepo()

	assemblerRatios := make(map[string]float32)
	for itemName, itemQuantity := range itemBOM {
		itemInfo := itemRepo.Lookup(itemName)

		craftSpeed := itemInfo.CraftSpeed
		yieldPerCraft := itemInfo.YieldPerCraft
		finalRatio := float32(itemQuantity) * craftSpeed /
			(float32(yieldPerCraft))

		assemblerRatios[itemName] = finalRatio
	}

	return assemblerRatios
}

// ConvertRatiosToIntegers takes a map of assembler ratios and finds
// the smallest ~equivilent integer ratios.
func ConvertRatiosToIntegers(
	floatRatios map[string]float32,
	yieldPerMin int,
	finalItem string) map[string]int {

	for {
		smallestNonZeroDecimal := findSmallestDecimalComponent(floatRatios)
		levelOfErrorIsInsignificant := smallestNonZeroDecimal < 0.01

		if levelOfErrorIsInsignificant {
			break
		}

		nextScalar := findNextScaler(smallestNonZeroDecimal)
		floatRatios = scaleRatios(floatRatios, nextScalar)
	}

	if yieldPerMin > 0 {
		floatRatios = scaleToMatchDesiredYield(
			floatRatios,
			yieldPerMin,
			finalItem)
	}

	integerRatios := make(map[string]int)
	for itemName, floatRatio := range floatRatios {
		integerRatios[itemName] = int(floatRatio)
		if floatRatio-float32(int(floatRatio)) > 0 {
			integerRatios[itemName] = integerRatios[itemName] + 1
		}
		if integerRatios[itemName] == 0 {
			integerRatios[itemName] = 1
		}
	}

	return integerRatios
}

func findSmallestDecimalComponent(
	floatRatios map[string]float32) float32 {

	var smallestNonZeroDecimal float32
	smallestNonZeroDecimal = 1.0
	for itemName, ratio := range floatRatios {
		// TEMP FOR HARD CODED CONSTRAINT RELAXER.
		if contains(RELAX, itemName) {
			continue
		}

		decimalComponent := ratio - float32(int(ratio))
		if decimalComponent > 0 &&
			decimalComponent < smallestNonZeroDecimal {
			smallestNonZeroDecimal = decimalComponent
		}
	}

	if smallestNonZeroDecimal == 1 {
		return 0
	}
	return smallestNonZeroDecimal
}

func findNextScaler(smallestDecimalComponent float32) float32 {
	return 1 / smallestDecimalComponent
}

func scaleRatios(
	floatRatios map[string]float32,
	scalar float32) map[string]float32 {

	scaledRatios := make(map[string]float32)
	for itemName, ratio := range floatRatios {
		scaledRatios[itemName] = ratio * scalar
	}

	return scaledRatios
}

func scaleToMatchDesiredYield(
	floatRatios map[string]float32,
	yieldPerMin int,
	finalItem string) map[string]float32 {

	itemRepo := items.GetItemRepo()
	craftedItem := itemRepo.Lookup(finalItem)
	assemblersNeeded := float32(yieldPerMin) /
		(float32(craftedItem.YieldPerCraft) * 60 / craftedItem.CraftSpeed)

	scalar := assemblersNeeded / floatRatios[finalItem]

	return scaleRatios(
		floatRatios,
		scalar)
}

// TEMP FOR HARD CODED CONSTRAINT RELAXER.
func contains(
	elements []string,
	theString string) bool {
	for _, element := range elements {
		if element == theString {
			return true
		}
	}
	return false
}
