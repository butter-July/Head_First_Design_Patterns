package main

import "fmt"

type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("search %s in %s", keyword, f.name)
	fmt.Println()
}
