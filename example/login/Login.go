package main

import (
	"errors"
	"math/rand"
	"time"

	sw "github.com/CalebJohnHunt/stacker"

	tea "github.com/charmbracelet/bubbletea"
)

type Login struct {
	authToken string
	err       error
}

type successfullyLoggedIn struct {
	authToken string
}

type failedLogin struct{}
type tryAgain struct{}

func (m *Login) Init() tea.Cmd {
	return sw.AddScene(&GetLineScene{prompt: "Enter username and password"})
}

func (m *Login) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var batch tea.Cmd
	switch msg := msg.(type) {
	case *GetLineScene:
		batch = tea.Batch(batch, func() tea.Msg {
			authToken := login(string(msg.line))
			if len(authToken) != 0 {
				return successfullyLoggedIn{authToken}
			}
			return failedLogin{}
		})
	case successfullyLoggedIn:
		m.authToken = msg.authToken
		return m, tea.Quit
	case failedLogin:
		m.err = errors.New("username or password wrong")
		batch = tea.Batch(batch, func() tea.Msg {
			time.Sleep(time.Second * 3)
			return tryAgain{}
		})
	case tryAgain:
		m.err = nil
		batch = tea.Batch(batch, sw.AddScene(&GetLineScene{prompt: "Try logging in again"}))
	}
	return m, batch
}

func login(s string) string {
	time.Sleep(time.Millisecond * 1500)
	if rand.New(rand.NewSource(time.Now().UnixMilli())).Intn(2) != 0 {
		return "Success!"
	}
	return ""
}

func (m *Login) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	return "logging in..."
}
