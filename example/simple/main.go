package main

import (
	sw "github.com/CalebJohnHunt/stacker"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	tea.NewProgram(sw.NewStacker(&model1{name: "a"})).Run()
}
