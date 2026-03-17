package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	editor "github.com/ionut-t/goeditor/adapter-bubbletea"
	"github.com/ionut-t/goeditor/core"
)

type model struct {
	editor editor.Model
	result string
}

func (m model) Init() tea.Cmd {
	return m.editor.CursorBlink()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.editor.SetSize(msg.Width, msg.Height)
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	case editor.SaveMsg:
		m.result = strings.TrimSpace(msg.Content)
		return m, tea.Quit

	case editor.QuitMsg:
		return m, tea.Quit

	case editor.ErrorMsg:
		if msg.ID == core.ErrNoChangesToSaveId {
			m.result = strings.TrimSpace(m.editor.GetCurrentContent())
			return m, tea.Quit
		}

		return m, m.editor.DispatchError(msg.Error, 3*time.Second)
	}

	var cmd tea.Cmd
	m.editor, cmd = m.editor.Update(msg)
	return m, cmd
}

func (m model) View() tea.View {
	v := tea.NewView(m.editor.View())
	v.AltScreen = true
	return v
}

func run(prefill string) (string, error) {
	e := editor.New(80, 24)
	e.Focus()
	e.SetLanguage("bash", "catppuccin-mocha")
	e.SetCursorMode(editor.CursorBlink)
	e.ShowTildeIndicator(true)

	e.SetContent(prefill)

	if prefill != "" {
		_ = e.SetCursorPositionEnd()
	}

	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		return "", fmt.Errorf("could not open tty: %w", err)
	}
	defer func() { _ = tty.Close() }()

	p := tea.NewProgram(model{editor: e}, tea.WithInput(tty), tea.WithOutput(tty))
	final, err := p.Run()
	if err != nil {
		return "", err
	}

	return final.(model).result, nil
}
