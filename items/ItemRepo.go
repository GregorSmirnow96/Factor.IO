package items

// A repository that holds metadata of all Factorio items.
type itemRepo struct {
	items map[string]Item
}

// The singleton instance of itemRepo.
var itemRepoInstance *itemRepo

// GetItemRepo is the getter for the singleton itemRepo
func GetItemRepo() itemRepo {
	if itemRepoInstance == nil {
		itemRepoInstance = &itemRepo{items: make(map[string]Item)}
	}

	return *itemRepoInstance
}

// AddItem adds an item to this repo's map of items. The key is
// the item's name.
func (repo itemRepo) AddItem(newItem Item) {
	repo.items[newItem.Name] = newItem
}

// Lookup finds an item in this repo's map of items, provided its
// name.
func (repo itemRepo) Lookup(itemName string) Item {
	return repo.items[itemName]
}
