package everquest

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func getRecentRosterDump(path string, guildName string) string {
	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(filepath.Base(path), guildName) {
			files = append(files, filepath.Base(path))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files[len(files)-1] // return last file - should be latest
}

type Guild struct {
	Members []GuildMember
}

type GuildMember struct {
	Name                string    // Character Name
	Level               int       // Character Level
	Class               string    // Player class ex: Necromancer
	Rank                string    // Guild Rank Name
	Alt                 bool      // Is this character flagged as an alt
	LastOnline          time.Time // Last time this character was online
	Zone                string    // Zone this character is currently in
	PublicNote          string    // Public Note
	PersonalNote        string    // Personal Note - assumed
	TributeStatus       bool      // Tribute status on or off
	TrophyTributeStatus bool      // Trophy Tribute Status on or off
	Donations           int       // total donations
	LastDonation        time.Time // Last date of donation
	PublicNote2         string    // Seems to be the public note again not sure why
	PersonalNote2       string    // Probably Personal Note again
}

// LoadFromPath takes a standard everquest raid dump and loads it into a struct
func (guild *Guild) LoadFromPath(path string) {
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
		level, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("Error converting level to int - Level: %s Name: %s\n", record[1], record[0])
			continue
		}
		var alt bool
		if record[4] == "A" {
			alt = true
		}
		format := "01/02/06"
		lastOnline, err := time.Parse(format, record[5])
		if err != nil {
			log.Printf("Error converting last_online to time - Time: %s\n", record[5])
			continue
		}
		var tributeStatus bool
		if record[9] == "on" {
			tributeStatus = true
		}
		var trophyTributeStatus bool
		if record[10] == "on" {
			trophyTributeStatus = true
		}
		donations, err := strconv.Atoi(record[11])
		if err != nil {
			log.Printf("Error converting donations to int - Donation: %s Name: %s\n", record[11], record[0])
			continue
		}
		lastDonation, err := time.Parse(format, record[12])
		if err != nil {
			log.Printf("Error converting last_donation to time - Time: %s\n", record[12])
			continue
		}

		guildMember := GuildMember{
			Name:                record[0],
			Level:               level,
			Class:               record[2],
			Rank:                record[3],
			Alt:                 alt,
			LastOnline:          lastOnline,
			Zone:                record[6],
			PublicNote:          record[7],
			PersonalNote:        record[8],
			TributeStatus:       tributeStatus,
			TrophyTributeStatus: trophyTributeStatus,
			Donations:           donations,
			LastDonation:        lastDonation,
			PublicNote2:         record[13],
			PersonalNote2:       record[14],
		}
		guild.Members = append(guild.Members, guildMember)
	}
}

func NewGuildMembers(master, new Guild) []GuildMember {
	masterList := make(map[string]interface{}, len(master.Members))
	var results []GuildMember
	for _, masterMember := range master.Members {
		masterList[masterMember.Name] = nil
	}
	for _, newMember := range new.Members {
		if _, ok := masterList[newMember.Name]; ok {
			results = append(results, newMember)
		}
	}
	return results
}

func MissingGuildMembers(master, new Guild) []GuildMember {
	return NewGuildMembers(new, master)
}
