package main

import (
	"sort"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type SortMode int

const (
	SortName SortMode = iota
	SortSize
	SortTime
)

func (m *model) sortByName() {
	sort.Slice(m.files, func(i, j int) bool {
		c := collate.New(language.Und, collate.IgnoreCase)
		return c.CompareString(m.files[i].Name(), m.files[j].Name()) == 1
	})
}

func (m *model) sortBySize() {
	sort.Slice(m.files, func(i, j int) bool {
		return m.files[i].Size() < m.files[j].Size()
	})
}

func (m *model) sortByModTime() {
	sort.Slice(m.files, func(i, j int) bool {
		return m.files[i].ModTime().After(m.files[j].ModTime())
	})
}
