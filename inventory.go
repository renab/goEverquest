package everquest

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// /outputfile Inventory defaults to base eq folder Charactername_server-Inventory.txt
// We should pull inventory when log see's it was dumped

// Inventory defines the equippable portion of a players inventory
type Inventory struct {
	Charm          int
	Ear1           int
	Ear1Slot1      int
	Head           int
	HeadSlot1      int
	HeadSlot2      int
	Face           int
	FaceSlot1      int
	Ear2           int
	Ear2Slot1      int
	Neck           int
	NeckSlot1      int
	Shoulders      int
	ShouldersSlot1 int
	Arms           int
	ArmsSlot1      int
	ArmsSlot2      int
	Back           int
	BackSlot1      int
	Wrist1         int
	Wrist1Slot1    int
	Wrist1Slot2    int
	Wrist2         int
	Wrist2Slot1    int
	Wrist2Slot2    int
	Range          int
	RangeSlot1     int
	Hands          int
	HandsSlot1     int
	HandsSlot2     int
	Primary        int
	PrimarySlot1   int
	PrimarySlot2   int
	Secondary      int
	SecondarySlot1 int
	Fingers1       int
	Fingers1Slot1  int
	Fingers2       int
	Fingers2Slot1  int
	Chest          int
	ChestSlot1     int
	ChestSlot2     int
	Legs           int
	LegsSlot1      int
	LegsSlot2      int
	Feet           int
	FeetSlot1      int
	FeetSlot2      int
	Waist          int
	WaistSlot1     int
	PowerSource    int
	Ammo           int
}

// LoadFromPath loads a player inventory from the standard eq export format for inventory
func (p *Inventory) LoadFromPath(path string) {
	// Open the file
	csvfile, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open the tsv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	r.Comma = '\t'
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	var i int
	var bEarSeen bool
	var bEarSlot1Seen bool
	var bWristSeen bool
	var bWristSlot1Seen bool
	var bWristSlot2Seen bool
	var bFingerSeen bool
	var bFingerSlot1Seen bool
	for {
		i++
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] == "Location" && i == 1 { // Header Row
			continue
		}
		item, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Printf("Error converting item to int - Item: %s ID: %s\n", record[1], record[2])
			continue
		}
		switch record[0] {
		case "Charm":
			p.Charm = item
		case "Ear": // Output file does not distinguish between each ear
			if !bEarSeen {
				p.Ear1 = item
				bEarSeen = true
			} else {
				p.Ear2 = item
			}
		case "Ear-Slot1": // Output file does not distinguish between each ear
			if !bEarSlot1Seen {
				p.Ear1Slot1 = item
				bEarSlot1Seen = true
			} else {
				p.Ear2Slot1 = item
			}
		case "Head":
			p.Head = item
		case "Head-Slot1":
			p.HeadSlot1 = item
		case "Head-Slot2":
			p.HeadSlot2 = item
		case "Face":
			p.Face = item
		case "Face-Slot1":
			p.FaceSlot1 = item
		case "Neck":
			p.Neck = item
		case "Neck-Slot1":
			p.NeckSlot1 = item
		case "Shoulders":
			p.Shoulders = item
		case "Shoulders-Slots1":
			p.ShouldersSlot1 = item
		case "Arms":
			p.Arms = item
		case "Arms-Slot1":
			p.ArmsSlot1 = item
		case "Arms-Slot2":
			p.ArmsSlot2 = item
		case "Back":
			p.Back = item
		case "Back-Slot1":
			p.BackSlot1 = item
		case "Wrist":
			if !bWristSeen {
				p.Wrist1 = item
				bWristSeen = true
			} else {
				p.Wrist2 = item
			}
		case "Wrist-Slot1":
			if !bWristSlot1Seen {
				p.Wrist1Slot1 = item
				bWristSlot1Seen = true
			} else {
				p.Wrist2Slot1 = item
			}
		case "Wrist-Slot2":
			if !bWristSlot2Seen {
				p.Wrist1Slot2 = item
				bWristSlot2Seen = true
			} else {
				p.Wrist2Slot2 = item
			}
		case "Range":
			p.Range = item
		case "Range-Slot1":
			p.RangeSlot1 = item
		case "Hands":
			p.Hands = item
		case "Hands-Slot1":
			p.HandsSlot1 = item
		case "Hands-Slot2":
			p.HandsSlot2 = item
		case "Primary":
			p.Primary = item
		case "Primary-Slot1":
			p.PrimarySlot1 = item
		case "Primary-Slot2":
			p.PrimarySlot2 = item
		case "Secondary":
			p.Secondary = item
		case "Secondary-Slot1":
			p.SecondarySlot1 = item
		case "Fingers":
			if !bFingerSeen {
				p.Fingers1 = item
				bFingerSeen = true
			} else {
				p.Fingers2 = item
			}
		case "Fingers-Slot1":
			if !bFingerSlot1Seen {
				p.Fingers1Slot1 = item
				bFingerSlot1Seen = true
			} else {
				p.Fingers2Slot1 = item
			}
		case "Chest":
			p.Chest = item
		case "Chest-Slot1":
			p.ChestSlot1 = item
		case "Chest-Slot2":
			p.ChestSlot2 = item
		case "Legs":
			p.Legs = item
		case "Legs-Slot1":
			p.LegsSlot1 = item
		case "Legs-Slot2":
			p.LegsSlot2 = item
		case "Feet":
			p.Feet = item
		case "Feet-Slot1":
			p.FeetSlot1 = item
		case "Feet-Slot2":
			p.FeetSlot2 = item
		case "Waist":
			p.Waist = item
		case "Waist-Slot1":
			p.WaistSlot1 = item
		case "Power Source":
			p.PowerSource = item
		case "Ammo":
			p.Ammo = item
		}
		// fmt.Printf("Location: %s Name: %s ID: %s Count: %s Slots: %s\n", record[0], record[1], record[2], record[3], record[4])
	}
}
