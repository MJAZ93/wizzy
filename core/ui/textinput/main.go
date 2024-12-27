package ui_textinput

import (
	"fmt"
	"github.com/fatih/color"
	"regexp"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func ReadText(title string, regexPattern string) (string, error) {
	var compiledRegex *regexp.Regexp
	var err error

	if regexPattern != "" {
		compiledRegex, err = regexp.Compile(regexPattern)
		if err != nil {
			return "", fmt.Errorf("invalid regex pattern: %v", err)
		}
	}

	for {
		p := tea.NewProgram(initialModel(title))
		result, err := p.Run()
		if err != nil {
			return "", err
		}

		m, ok := result.(model)
		if !ok {
			return "", fmt.Errorf("unexpected model type")
		}

		input := m.textInput.Value()
		if m.cancelled {
			return "", fmt.Errorf("input cancelled")
		}
		if compiledRegex == nil || compiledRegex.MatchString(input) {
			return input, nil
		} else {
			s := fmt.Sprintf("Input does not match the required pattern. Please try again. Valid regex: %s", compiledRegex)
			color.Red(s)
		}
	}
}

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
	title     string
	cancelled bool
}

func initialModel(title string) model {
	ti := textinput.New()
	ti.Focus()
	ti.Width = 100

	return model{
		textInput: ti,
		err:       nil,
		title:     title,
		cancelled: false,
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
		case tea.KeyEnter:
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			m.cancelled = true
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n%s\n",
		m.title,
		m.textInput.View(),
	) + "\n"
}
