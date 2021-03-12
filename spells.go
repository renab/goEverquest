package everquest

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Spellbook contains a slice of all known spells
type Spellbook struct {
	Spells []Spell
}

// Spell defines a spellbook entry
type Spell struct {
	Level int
	Name  string
}

// LoadFromPath loads a standard everquest spell dump
func (s *Spellbook) LoadFromPath(path string) {
	// Open the file
	tsvfile, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open the tsv file", err)
	}

	// Parse the file
	r := csv.NewReader(tsvfile)
	r.Comma = '\t'
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		level, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Printf("Error converting level to int - Level: %s Name: %s\n", record[0], record[1])
			continue
		}
		spell := Spell{
			Level: level,
			Name:  record[1],
		}
		s.Spells = append(s.Spells, spell)
	}
}
