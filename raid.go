package everquest

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Raid contains all members of a raid dump
type Raid struct {
	Members []RaidMember
}

// RaidMember is a struct containing all raid dump info
type RaidMember struct {
	Group  int    // group number
	Player string // player number
	Level  int    // player level
	Class  string // player class
	Role   string // raid role - group leader, raid leader
	Unk1   string // unknown
	Unk2   string // unknown
	Unk3   string // "Yes"??????
}

// LoadFromPath takes a standard everquest raid dump and loads it into a struct
func (raid *Raid) LoadFromPath(path string) {
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
		group, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Printf("Error converting group to int - Level: %s Name: %s\n", record[0], record[1])
			continue
		}
		level, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Printf("Error converting level to int - Level: %s Name: %s\n", record[0], record[1])
			continue
		}
		raidMember := RaidMember{
			Group:  group,
			Player: record[1],
			Level:  level,
			Class:  record[3],
			Role:   record[4],
			Unk1:   record[5],
			Unk2:   record[6],
			Unk3:   record[7],
		}
		raid.Members = append(raid.Members, raidMember)
	}
}

func getRecentRaidDump(path string) string {
	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(filepath.Base(path), "RaidRoster") {
			files = append(files, filepath.Base(path))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files[len(files)-1] // return last file - should be latest
}

func NewRaidMembers(master, new Raid) []RaidMember {
	masterList := make(map[string]interface{}, len(master.Members))
	var results []RaidMember
	for _, masterMember := range master.Members {
		masterList[masterMember.Player] = nil
	}
	for _, newMember := range new.Members {
		if _, ok := masterList[newMember.Player]; ok {
			results = append(results, newMember)
		}
	}
	return results
}

func MissingRaidMembers(master, new Raid) []RaidMember {
	return NewRaidMembers(new, master)
}
