package main

type PancakeHouseIterator struct {
	menu  *PancakeHouseMenu
	index int
}

func (p *PancakeHouseIterator) next() MenuItem {
	if p.HasNext() {
		item := p.menu.menuItem[p.index]
		p.index++
		return item
	}
	return MenuItem{}

}
func (it *PancakeHouseIterator) HasNext() bool {
	return it.index < len(it.menu.menuItem)
}
