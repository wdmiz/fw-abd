package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	path  string
	files []fs.FileInfo

	selected int
	showSize bool

	width, height int
	capacity      int

	showFrom int

	showLog bool

	log []string
}

func main() {
	pathFlag := flag.String("path", ".", "path")
	flag.Parse()

	p := tea.NewProgram(initModel(*pathFlag), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Failed to run: %v\n", err)
		os.Exit(1)
	}
}

func initModel(path string) (m model) {
	m = model{}

	err := m.loadDir(path)

	for _, e := range err {
		m.log = append(m.log, e.Error())
	}

	return
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyDown:
			if m.selected >= len(m.files)-1 {
				break
			}

			m.selected++

			if m.selected > m.showFrom+(m.height-6) {
				m.showFrom = m.selected - (m.height - 6)
			}

		case tea.KeyUp:
			if m.selected <= 0 {
				break
			}

			m.selected--

			if m.selected < m.showFrom {
				m.showFrom = m.selected
			}

		case tea.KeyCtrlD:
			m.loadDir(filepath.Dir(m.path))

		case tea.KeyCtrlI:
			m.showSize = !m.showSize

		case tea.KeyEnter:
			if m.selected > len(m.files) || m.selected < 0 {
				break
			}

			if m.files[m.selected].IsDir() {
				m.loadDir(filepath.Join(m.path, m.files[m.selected].Name()))
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.capacity = m.height - 6
	}
	return m, nil
}

func (m model) View() (v string) {
	if m.capacity < 0 {
		return "Terminal too small, need >6 rows"

	}
	return m.asList()
}

func (m *model) loadDir(path string) (errors []error) {
	path, err := filepath.Abs(path)

	if err != nil {
		errors = append(errors, err)
		return
	}

	entries, err := os.ReadDir(path)

	if err != nil {
		errors = append(errors, err)
		return
	}

	m.files = nil
	m.path = path
	m.selected = 0
	m.showFrom = 0

	for _, entry := range entries {
		fi, err := entry.Info()

		if err != nil {
			errors = append(errors, err)
			continue
		}

		m.files = append(m.files, fi)
	}

	return
}
