package main

type PancakeHouseIterator struct {
	menu  []MenuItem
	index int
}

func (p *PancakeHouseIterator) next() *MenuItem {
	if p.HasNext() {
		item := &p.menu[p.index]
		p.index++
		return item
	}
	return nil

}
func (it *PancakeHouseIterator) HasNext() bool {
	return it.index < len(it.menu)
}
