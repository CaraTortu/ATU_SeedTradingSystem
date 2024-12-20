package currency

// Item will store the name and cost of the item
type Item struct {
	Cost int
	Name string
}

// Storage will store all the items that the user has bought
type Storage struct {
	Items   map[Item]int
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
func NewStorage(balance int) Storage {
	return Storage{
		Items:   make(map[Item]int),
		Balance: balance,
	}
}

// AddItem will add an item to the storage
func (s *Storage) AddItem(item Item) {
	if _, ok := s.Items[item]; ok {
		s.Items[item]++
		return
	}
	s.Items[item] = 1
}

// RemoveItem will remove an item from the storage
func (s *Storage) RemoveItem(item *Item) (deleted bool) {
	if _, ok := s.Items[*item]; ok {
		s.Items[*item]--
		if s.Items[*item] == 0 {
			delete(s.Items, *item)
		}

		return true
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
func (s *Storage) HasItem(item Item) (has_item bool) {
	_, ok := s.Items[item]
	return ok
}
