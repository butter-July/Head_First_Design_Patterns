package main

import "fmt"

type HomeTheaterFacade struct {
	amp *Amplifier
	//tuner *Tuner
	player  *StreamingPlayer
	project *Project
	screen  *Screen
	popper  *PopcornPopper
	lights  *TheaterLights
}

func (homefaced *HomeTheaterFacade) watchMovie() {
	fmt.Println("Get ready to watch a movie")
	homefaced.popper.on()
	homefaced.popper.pop()
	homefaced.lights.dim()
	homefaced.screen.down()
	homefaced.project.on()
	homefaced.project.wideScreenMode()
	homefaced.amp.on()
	homefaced.amp.setStreamingPlayer()
	homefaced.amp.setSurroundingSound()
	homefaced.amp.setVolume(5)
	homefaced.player.on()
	homefaced.player.play("Raiders of the  Lost Ark")
}
func (homeFaced *HomeTheaterFacade) endMovie() {
	fmt.Println("Shutting movie theater down...")
	homeFaced.popper.off()
	homeFaced.lights.on()
	homeFaced.screen.up()
	homeFaced.project.off()
	homeFaced.amp.off()
	homeFaced.player.stop("Raiders of the  Lost Ark")
	homeFaced.player.off()
}
