package ui_textarea

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
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

	if m.canceled {
		return "", fmt.Errorf("user canceled")
	}

	return m.textarea.Value(), nil
}

type errMsg error

type model struct {
	textarea textarea.Model
	title    string
	err      error
	saved    bool
	canceled bool
}

func initialModel(title string) model {
	ti := textarea.New()
	ti.Focus()

	return model{
		textarea: ti,
		title:    title,
		err:      nil,
		saved:    false,
		canceled: false,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+s":
			m.saved = true
			return m, tea.Quit
		case "ctrl+c":
			m.canceled = true
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n%s\n%s",
		m.title,
		m.textarea.View(),
		"(ctrl+s to save)",
	) + "\n"
}

func main() {
	title := "Enter your text"
	text, err := ReadText(title)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if text == "" {
		fmt.Println("Canceled")
	} else {
		fmt.Printf("Saved text: %q\n", text)
	}
}
