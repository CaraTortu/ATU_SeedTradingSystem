package currency

import "testing"

// Testing if the storage can be created
func TestNewStorage(t *testing.T) {
	storage := NewStorage(0)
	if storage.Balance != 0 || len(storage.Items) != 0 {
		t.Errorf("Storage was not created correctly. It should have no items and a balance of 0")
	}
}

// Testing if an item can be added to the storage
func TestAddItem(t *testing.T) {
	storage := NewStorage(0)
	storage.AddItem(Items[ItemsAmount-1])

	// Make sure the item is in the storage
	has_item := storage.HasItem(Items[ItemsAmount-1])

	if !has_item {
		t.Errorf("Item was not found in the storage and it should have been")
	}
}

// Testing if an item can be removed from the storage
func TestRemoveItem(t *testing.T) {
	itemIdx := ItemsAmount - 1
	storage := NewStorage(0)
	storage.AddItem(Items[itemIdx])

	// Make sure the item is in the storage
	if len(storage.Items) != 1 {
		t.Errorf("Item was not added to the storage and it should have been")
	}

	if !storage.RemoveItem(&Items[itemIdx]) || len(storage.Items) != 0 {
		t.Errorf("Item was not removed from the storage and it should have been")
	}
}

// Testing if an item can be bought
func TestBuyItem(t *testing.T) {
	storage := NewStorage(10)

	expensiveItem := Item{
		Cost: 100,
		Name: "Expensive Item",
	}

	if storage.BuyItem(expensiveItem) {
		t.Errorf("Item was bought even though the balance was not enough. It should not have been bought")
	}

	storage.Balance = 100
	if !storage.BuyItem(expensiveItem) {
		t.Errorf("Item was not bought even though the balance was enough. It should have been bought")
	}
}
