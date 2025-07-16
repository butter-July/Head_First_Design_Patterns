package main

import "fmt"

type Project struct {
}

func (project *Project) on() {
	fmt.Println("Project on")
}
func (project *Project) wideScreenMode() {
	fmt.Println("Project in  widescreen mode (16 x 9 aspect ratio)")
}
func (projecter *Project) off() {
	fmt.Println("Projector off!")
}
