package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

type Persistence struct {
	lineSeparator string
}

// AddEntry return int
func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

// RemoveEntry return void
func (j *Journal) RemoveEntry(index int) {
	// ...
}

// separation on concerns

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

func main() {
	j := Journal{}
	j.AddEntry("Entry 1")
	j.AddEntry("Entry 2")
	fmt.Println(j.String())

	//
	SaveToFile(&j, "journal.txt")

	//
	p := Persistence{"\n "}
	p.SaveToFile(&j, "journal.txt")

}
