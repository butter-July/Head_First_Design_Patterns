package main

func main() {
	remote := SimpleRemoteControl{}
	light := Light{}
	LigthOn := LightOnCommand{&light}
	Lightoff := LightOffCommand{&light}

	remote.SetCommand(1, &LigthOn, &Lightoff)
	remote.OnButtonWasPressed(1)
	
}
