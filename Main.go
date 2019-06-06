package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"./bom"
	"./items"
	"./ratios"
)

func main() {
	craftedItemPtr := flag.String("Item", "Iron Gear", "The item you want to craft.")
	yieldPerMinPtr := flag.String("YieldPerMin", "-1", "The number of this item you want per minute.")
	flag.Parse()
	craftedItemName := *craftedItemPtr
	yieldPerMinString := *yieldPerMinPtr
	ratios.RELAX = flag.Args()

	// Load items into hashmap from json file.
	// Put these tuples in a hashmap of (name:item).

	jsonFile, err := os.Open("./Items.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var itemStructs []items.Item
	json.Unmarshal(byteValue, &itemStructs)

	itemRepo := items.GetItemRepo()
	for _, item := range itemStructs {
		itemRepo.AddItem(item)
	}

	// Find BOM of a selected item,
	// then calculate ratio of assemblers per item:
	// #_of_item / (craft_speed * yield_per_craft)
	itemBOM := bom.CalculateBOM(craftedItemName)
	floatRatios := ratios.FindFloatAssemblerRatios(itemBOM)

	yieldPerMin, _ := strconv.ParseInt(yieldPerMinString, 10, 64)
	integerRatios := ratios.ConvertRatiosToIntegers(
		floatRatios,
		int(yieldPerMin),
		craftedItemName)

	for itemName, assemblerQuantity := range integerRatios {
		fmt.Println(itemName+":", assemblerQuantity)
	}

	// Create GUI to select an item.
}
