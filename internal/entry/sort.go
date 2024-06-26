package entry

import (
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type ByName []Entry

func (f ByName) Len() int {
	return len(f)
}

func (f ByName) Less(i, j int) bool {
	c := collate.New(language.Und, collate.IgnoreCase)
	return c.CompareString(f[i].Info.Name(), f[j].Info.Name()) > 0
}

func (f ByName) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type BySize []Entry

func (f BySize) Len() int {
	return len(f)
}

func (f BySize) Less(i, j int) bool {
	return f[i].Info.Size() > f[j].Info.Size()
}

func (f BySize) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type ByModTime []Entry

func (f ByModTime) Len() int {
	return len(f)
}

func (f ByModTime) Less(i, j int) bool {
	return f[i].Info.ModTime().After(f[j].Info.ModTime())
}

func (f ByModTime) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
