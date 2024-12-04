package currency

// Item will store the name and cost of the item
type Item struct {
	Cost int
	Name string
}

// Storage will store all the items that the user has bought
type Storage struct {
	Items   []Item
	Balance int
}

// Global items will store all the possible items that can be bought
var Items = []Item{
	{
		Cost: 10,
		Name: "Apple",
	},
}

var ItemsAmount = len(Items)

// NewStorage will create a new storage
func NewStorage() Storage {
	return Storage{
		Items:   []Item{},
		Balance: 0,
	}
}

// AddItem will add an item to the storage
func (s *Storage) AddItem(item Item) {
	s.Items = append(s.Items, item)
}

// RemoveItem will remove an item from the storage
func (s *Storage) RemoveItem(item *Item) (deleted bool) {
	for i, v := range s.Items {
		if v == *item {
			s.Items = append(s.Items[:i], s.Items[i+1:]...)
			return true
		}
	}

	return false
}

// BuyItem will buy an item if the user has enough balance
func (s *Storage) BuyItem(item Item) (success bool) {
	if s.Balance >= item.Cost {
		s.Balance -= item.Cost
		s.AddItem(item)
		return true
	}

	return false
}

// HasItem will check if the storage has the item
func (s *Storage) HasItem(item Item) (index int) {
	for i, v := range s.Items {
		if v == item {
			return i
		}
	}

	return -1
}
