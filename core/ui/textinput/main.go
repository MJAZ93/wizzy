package ui_textinput

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func ReadText(title string) (string, error) {
	p := tea.NewProgram(initialModel(title))
	result, err := p.Run()
	if err != nil {
		return "", err
	}

	m, ok := result.(model)
	if !ok {
		return "", fmt.Errorf("unexpected model type")
	}

	return m.textInput.Value(), nil
}

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
	title     string
}

func initialModel(title string) model {
	ti := textinput.New()
	ti.Focus()
	ti.Width = 100

	return model{
		textInput: ti,
		err:       nil,
		title:     title,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"\n\n%s\n%s\n%s",
		m.title,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
