package main

func main() {
	s := InitScreen()
	EventLoop(s)
	s.Fini()
}
