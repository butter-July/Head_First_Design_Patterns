package main

type DinerMenuIterator struct {
	items    []*MenuItem
	position int
}

func (diner *DinerMenuIterator) HasNext() bool {
	if diner.position >= len(diner.items) {
		return false
	}
	return true
}
func (diner *DinerMenuIterator) next() MenuItem {
	menuItem := diner.items[diner.position]
	diner.position++
	return *menuItem
}
