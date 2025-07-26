package main

import "fmt"

type Folder struct {
	name  string
	small []Component
}

func (f *Folder) search(keyword string) {
	fmt.Printf("searching %s,in %s", keyword, f.name)
	fmt.Println()
	for _, thing := range f.small {
		thing.search(keyword) //递归叭.看他里面有没有其他的"小盒子"
	}
}
func (f *Folder) add(folder Component) {
	f.small = append(f.small, folder)
}
