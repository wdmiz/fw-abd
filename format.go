package main

import (
	"fmt"
	"io/fs"

	"github.com/charmbracelet/lipgloss"
)

var (
	styleHeader    = lipgloss.NewStyle().Background(lipgloss.Color("15")).Foreground(lipgloss.Color("0"))
	styleDefault   = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	styleDirectory = lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Bold(true)
	styleSymlink   = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	styleAux       = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
)

func (m *model) asList() (out string) {
	out = m.path + "\n"

	labels := " "

	if m.showSize {
		labels += " SIZE   "
	}

	labels += " NAME "

	out += styleHeader.Render(labels) + "\n"

	if m.showFrom > 0 {
		out += styleAux.Render(fmt.Sprintf("(%d more above)", m.showFrom)) + "\n"
	} else {
		out += styleAux.Render("(Beginning)") + "\n"
	}

	for i := m.showFrom; i-m.showFrom <= m.capacity && i < len(m.files); i++ {
		f := m.files[i]

		var text string
		var style lipgloss.Style

		if f.IsDir() {
			text = "D "
			style = styleDirectory
		} else if f.Mode()&fs.ModeSymlink != 0 {
			text = "S "
			style = styleSymlink
		} else {
			text = "- "
			style = styleDefault
		}

		if m.showSize {
			if !f.IsDir() && f.Mode()&fs.ModeSymlink == 0 {
				text += formatSize(f.Size()) + " "
			} else if f.Mode()&fs.ModeSymlink != 0 {
				text += "SYMLINK "
			} else {
				text += "――――――― "
			}

		}

		text += f.Name()

		if m.selected == i {
			style = style.Copy().Reverse(true)
		}

		out += style.Render(text) + "\n"
	}

	if n := len(m.files) - 1 - m.showFrom; n > m.capacity {
		out += styleAux.Render(fmt.Sprintf("(%d more below)", n-m.capacity)) + "\n"
	} else {
		out += styleAux.Render("(End)") + "\n"
	}

	return
}

func formatSize(byteSize int64) string {
	if byteSize <= 1000 {
		return fmt.Sprintf("%7s", fmt.Sprintf("%dB", byteSize))
	}

	power := 1
	s := byteSize
	for s > 1000000 {
		power++
		s = s / 1000
	}

	size := float64(s) / 1000.0
	out := fmt.Sprintf("%5s", fmt.Sprintf("%.4g", size))

	switch power {
	case 1:
		out += "kB"
	case 2:
		out += "MB"
	case 3:
		out += "GB"
	case 4:
		out += "TB"
	case 5:
		out += "PB"
	case 6:
		out += "EB"
	case 7:
		out += "ZB"
	case 8:
		out += "YB"
	case 9:
		out += "RB"
	case 10:
		out += "QB"
	}

	return out
}
