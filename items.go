package everquest

import (
	"bufio"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type ItemDB struct {
	items map[int]Item   // primary way to find an item is by it's ID #
	names map[string]int // Used to fast lookup ID by name (there may be duplicates, and is not recommended)
}

func (db *ItemDB) LoadFromFile(file string) {
	db.items = make(map[int]Item)
	db.names = make(map[string]int)
	psvfile, err := os.Open(file)
	if err != nil {
		log.Fatalln("Couldn't open the psv file", err)
	}
	defer psvfile.Close()

	r := bufio.NewReader(psvfile)

	// Iterate through the records
	headerSkipped := false
	var itemCount int
	for {
		// Read each record from csv
		line, tooLong, err := r.ReadLine()
		if !headerSkipped {
			headerSkipped = true
			// skip header line
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if tooLong {
			log.Printf("Item line too long!\n")
		}
		record := strings.Split(string(line), `|`)
		var item Item
		item.Load(record...)
		db.items[item.ID] = item
		db.names[strings.ToLower(item.Name)] = item.ID
		itemCount++
	}
	log.Printf("Loaded %d items\n", itemCount)
}

// FindIDByName does an item lookup by the item name, returns -1 if not found
func (db *ItemDB) FindIDByName(name string) int {
	lower := strings.ToLower(name)
	if val, ok := db.names[lower]; ok {
		return val
	}
	return -1
}

// GetItemByID returns an item given its ID, returns an empty struct if not found
func (db *ItemDB) GetItemByID(id int) Item {
	if val, ok := db.items[id]; ok {
		return val
	}
	return Item{}
}

// DownloadFile will download a itemdb given the url to the gz file and decompress it
// Expecting http://items.sodeq.org/downloads/items.txt.gz style db
func (db *ItemDB) Download(filepath, url string) error {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// Get the data
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		file, err := ioutil.TempFile("", "items.*.gz")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(file.Name())

		// Write the body to file
		_, err = io.Copy(file, resp.Body)
		err = db.decompress(file.Name(), filepath)
		return err
	}
	log.Printf("ItemDB already exists, skipping download\n")
	return nil
}

func (db *ItemDB) decompress(in, out string) error {
	// Open the gzip file.
	f, _ := os.Open(in)
	// close f on exit and check for its returned error
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	// Create new reader to decompress gzip.
	r, _ := gzip.NewReader(f)

	// open output file
	fo, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)
	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err = w.Flush(); err != nil {
		panic(err)
	}
	return err
}

// Load is for taking in csv data to generate an item
func (i *Item) Load(data ...string) {
	var x int
	var d int
	d, _ = strconv.Atoi(data[x])
	i.Itemclass = d
	x++
	i.Name = data[x]
	x++
	i.Lore = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	i.Idfile = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Idfileextra = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.ID = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Weight = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Norent = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Nodrop = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Attunable = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Size = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Slots = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Price = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Icon = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Benefitflag = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Tradeskills = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Cr = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Dr = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Pr = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Mr = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Fr = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Svcorruption = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Astr = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Asta = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Aagi = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Adex = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Acha = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Aint = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Awis = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Hp = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Mana = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Endurance = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Ac = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Regen = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Manaregen = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Enduranceregen = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Classes = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Races = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Deity = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Skillmodvalue = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Skillmodmax = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Skillmodtype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Skillmodextra = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Banedmgrace = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Banedmgbody = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Banedmgraceamt = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Banedmgamt = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Magic = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Foodduration = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Reqlevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Reclevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Reqskill = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bardtype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bardvalue = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN02 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN03 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN04 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Light = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Delay = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Elemdmgtype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Elemdmgamt = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Therange = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Damage = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Color = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Prestige = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN06 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN07 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN08 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Itemtype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Material = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Materialunk1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Elitematerial = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Heroforge1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Heroforge2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Sellrate = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Extradmgskill = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Extradmgamt = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Charmfileid = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Mounteffect = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Mountlevel2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Mounteffecttype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Mountlevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Mountunk1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Mountunk2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Mountunk3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Mountunk4 = d
	x++
	i.Charmfile = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augtype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augstricthidden = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augrestrict = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot1type = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot1visible = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot1unk = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot2type = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot2visible = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot2unk = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot3type = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot3visible = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot3unk = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot4type = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot4visible = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot4unk = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot5type = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot5visible = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot5unk = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot6type = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot6visible = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augslot6unk = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Pointtype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Ldontheme = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Ldonprice = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Ldonsellbackrate = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Ldonsold = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bagtype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bagslots = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bagsize = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bagwr = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Booktype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Booklang = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Filename = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Loregroup = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Artifactflag = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Summonedflag = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Favor = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Fvnodrop = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Attack = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Haste = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Guildfavor = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Augdistiller = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN09 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN10 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Nopet = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN11 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Stacksize = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Notransfer = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Expendablearrow = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN12 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN13 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Clickeffect = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Clicklevel2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Clicktype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Clicklevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Maxcharges = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Casttime = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Recastdelay = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Recasttype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Clickunk5 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Clickname = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Clickunk7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Proceffect = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Proclevel2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Proctype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Proclevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Prockunk1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Procunk2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Procunk3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Procunk4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Procrate = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Procname = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Procunk7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Worneffect = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Wornlevel2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Worntype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Wornlevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Wornunk1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Wornunk2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Wornunk3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Wornunk4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Wornunk5 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Wornname = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Wornunk7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Focuseffect = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Focuslevel2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Focustype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Focuslevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Focusunk1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Focusunk2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Focusunk3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Focusunk4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Focusunk5 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Focusname = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Focusunk7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Scrolleffect = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Scrolllevel2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Scrolleffecttype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Scrolllevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Scrollunk1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Scrollunk2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Scrollunk3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Scrollunk4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Scrollunk5 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Scrollname = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Scrollunk7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bardeffect = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bardlevel2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bardeffecttype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bardlevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bardunk1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bardunk2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bardunk3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bardunk4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bardunk5 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bardname = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Bardunk7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Mountunk5 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Blessingeffect = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Blessingname = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Mountname = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Mountunk7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Blessinglevel2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Blessingeffecttype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Blessinglevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Blessingunk1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Rightclickscriptid = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Questitemflag = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Powersourcecap = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Purity = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Epicitem = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Backstabdmg = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.HeroicStr = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.HeroicInt = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.HeroicWis = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.HeroicAgi = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.HeroicDex = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.HeroicSta = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.HeroicCha = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN29 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Healamt = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Spelldmg = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Clairvoyance = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN30 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN31 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN32 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN33 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN34 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN35 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN36 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN37 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Heirloom = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Placeablebitfield = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN38 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN39 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN40 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN41 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN42 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN43 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN44 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Placeablenpcname = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN46 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN47 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN48 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN49 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN50 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN51 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN52 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN53 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN54 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN55 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN56 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN57 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN58 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN59 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN60 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN61 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN62 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN63 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Collectible = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Nodestroy = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Nonpc = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Nozone = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN68 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN69 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN70 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN71 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Noground = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN73 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Marketplace = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Freestorage = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN76 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN77 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN78 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN79 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Blessingunk2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Blessingunk3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Blessingunk4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Blessingunk5 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Blessingunk7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Familiareffect = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Familiarlevel2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Familiareffecttype = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Familiarlevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Familiarunk1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Familiarunk2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Familiarunk3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Familiarunk4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Familiarunk5 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Familiarname = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Familiarunk7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.UNKNOWN80 = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Minluck = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Maxluck = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Loreequippedgroup = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Evoitem = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Evoid = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Evolvl = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Evomax = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Convertitem = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Convertid = d
	x++
	d, _ = strconv.Atoi(data[x])
	i.Convertname = d
	x++
	// 2020-04-19 22:57:39
	updated, _ := time.Parse("2006-01-02 15:04:05", data[x])
	i.Updated = updated
	x++
	created, _ := time.Parse("2006-01-02 15:04:05", data[x])
	i.Created = created
	x++
	i.Submitter = data[x]
	x++
	verified, _ := time.Parse("2006-01-02 15:04:05", data[x])
	i.Verified = verified
	x++
	i.Verifiedby = data[x]
	x++
	i.Collectversion = data[x]
	x++
}

// Item is an Everquest Item
type Item struct {
	Itemclass          int       `db:"itemclass"`
	Name               string    `db:"name"`
	Lore               string    `db:"lore"`
	Idfile             int       `db:"idfile"`
	Idfileextra        int       `db:"idfileextra"`
	ID                 int       `db:"id"`
	Weight             int       `db:"weight"`
	Norent             int       `db:"norent"`
	Nodrop             int       `db:"nodrop"`
	Attunable          int       `db:"attunable"`
	Size               int       `db:"size"`
	Slots              int       `db:"slots"`
	Price              int       `db:"price"`
	Icon               int       `db:"icon"`
	Benefitflag        int       `db:"benefitflag"`
	Tradeskills        int       `db:"tradeskills"`
	Cr                 int       `db:"cr"`
	Dr                 int       `db:"dr"`
	Pr                 int       `db:"pr"`
	Mr                 int       `db:"mr"`
	Fr                 int       `db:"fr"`
	Svcorruption       int       `db:"svcorruption"`
	Astr               int       `db:"astr"`
	Asta               int       `db:"asta"`
	Aagi               int       `db:"aagi"`
	Adex               int       `db:"adex"`
	Acha               int       `db:"acha"`
	Aint               int       `db:"aint"`
	Awis               int       `db:"awis"`
	Hp                 int       `db:"hp"`
	Mana               int       `db:"mana"`
	Endurance          int       `db:"endurance"`
	Ac                 int       `db:"ac"`
	Regen              int       `db:"regen"`
	Manaregen          int       `db:"manaregen"`
	Enduranceregen     int       `db:"enduranceregen"`
	Classes            int       `db:"classes"`
	Races              int       `db:"races"`
	Deity              int       `db:"deity"`
	Skillmodvalue      int       `db:"skillmodvalue"`
	Skillmodmax        int       `db:"skillmodmax"`
	Skillmodtype       int       `db:"skillmodtype"`
	Skillmodextra      int       `db:"skillmodextra"`
	Banedmgrace        int       `db:"banedmgrace"`
	Banedmgbody        int       `db:"banedmgbody"`
	Banedmgraceamt     int       `db:"banedmgraceamt"`
	Banedmgamt         int       `db:"banedmgamt"`
	Magic              int       `db:"magic"`
	Foodduration       int       `db:"foodduration"`
	Reqlevel           int       `db:"reqlevel"`
	Reclevel           int       `db:"reclevel"`
	Reqskill           int       `db:"reqskill"`
	Bardtype           int       `db:"bardtype"`
	Bardvalue          int       `db:"bardvalue"`
	UNKNOWN02          int       `db:"UNKNOWN02"`
	UNKNOWN03          int       `db:"UNKNOWN03"`
	UNKNOWN04          int       `db:"UNKNOWN04"`
	Light              int       `db:"light"`
	Delay              int       `db:"delay"`
	Elemdmgtype        int       `db:"elemdmgtype"`
	Elemdmgamt         int       `db:"elemdmgamt"`
	Therange           int       `db:"therange"`
	Damage             int       `db:"damage"`
	Color              int       `db:"color"`
	Prestige           int       `db:"prestige"`
	UNKNOWN06          int       `db:"UNKNOWN06"`
	UNKNOWN07          int       `db:"UNKNOWN07"`
	UNKNOWN08          int       `db:"UNKNOWN08"`
	Itemtype           int       `db:"itemtype"`
	Material           int       `db:"material"`
	Materialunk1       int       `db:"materialunk1"`
	Elitematerial      int       `db:"elitematerial"`
	Heroforge1         int       `db:"heroforge1"`
	Heroforge2         int       `db:"heroforge2"`
	Sellrate           int       `db:"sellrate"`
	Extradmgskill      int       `db:"extradmgskill"`
	Extradmgamt        int       `db:"extradmgamt"`
	Charmfileid        int       `db:"charmfileid"`
	Mounteffect        int       `db:"mounteffect"`
	Mountlevel2        int       `db:"mountlevel2"`
	Mounteffecttype    int       `db:"mounteffecttype"`
	Mountlevel         int       `db:"mountlevel"`
	Mountunk1          int       `db:"mountunk1"`
	Mountunk2          int       `db:"mountunk2"`
	Mountunk3          int       `db:"mountunk3"`
	Mountunk4          int       `db:"mountunk4"`
	Charmfile          string    `db:"charmfile"`
	Augtype            int       `db:"augtype"`
	Augstricthidden    int       `db:"augstricthidden"`
	Augrestrict        int       `db:"augrestrict"`
	Augslot1type       int       `db:"augslot1type"`
	Augslot1visible    int       `db:"augslot1visible"`
	Augslot1unk        int       `db:"augslot1unk"`
	Augslot2type       int       `db:"augslot2type"`
	Augslot2visible    int       `db:"augslot2visible"`
	Augslot2unk        int       `db:"augslot2unk"`
	Augslot3type       int       `db:"augslot3type"`
	Augslot3visible    int       `db:"augslot3visible"`
	Augslot3unk        int       `db:"augslot3unk"`
	Augslot4type       int       `db:"augslot4type"`
	Augslot4visible    int       `db:"augslot4visible"`
	Augslot4unk        int       `db:"augslot4unk"`
	Augslot5type       int       `db:"augslot5type"`
	Augslot5visible    int       `db:"augslot5visible"`
	Augslot5unk        int       `db:"augslot5unk"`
	Augslot6type       int       `db:"augslot6type"`
	Augslot6visible    int       `db:"augslot6visible"`
	Augslot6unk        int       `db:"augslot6unk"`
	Pointtype          int       `db:"pointtype"`
	Ldontheme          int       `db:"ldontheme"`
	Ldonprice          int       `db:"ldonprice"`
	Ldonsellbackrate   int       `db:"ldonsellbackrate"`
	Ldonsold           int       `db:"ldonsold"`
	Bagtype            int       `db:"bagtype"`
	Bagslots           int       `db:"bagslots"`
	Bagsize            int       `db:"bagsize"`
	Bagwr              int       `db:"bagwr"`
	Booktype           int       `db:"booktype"`
	Booklang           int       `db:"booklang"`
	Filename           int       `db:"filename"`
	Loregroup          int       `db:"loregroup"`
	Artifactflag       int       `db:"artifactflag"`
	Summonedflag       int       `db:"summonedflag"`
	Favor              int       `db:"favor"`
	Fvnodrop           int       `db:"fvnodrop"`
	Attack             int       `db:"attack"`
	Haste              int       `db:"haste"`
	Guildfavor         int       `db:"guildfavor"`
	Augdistiller       int       `db:"augdistiller"`
	UNKNOWN09          int       `db:"UNKNOWN09"`
	UNKNOWN10          int       `db:"UNKNOWN10"`
	Nopet              int       `db:"nopet"`
	UNKNOWN11          int       `db:"UNKNOWN11"`
	Stacksize          int       `db:"stacksize"`
	Notransfer         int       `db:"notransfer"`
	Expendablearrow    int       `db:"expendablearrow"`
	UNKNOWN12          int       `db:"UNKNOWN12"`
	UNKNOWN13          int       `db:"UNKNOWN13"`
	Clickeffect        int       `db:"clickeffect"`
	Clicklevel2        int       `db:"clicklevel2"`
	Clicktype          int       `db:"clicktype"`
	Clicklevel         int       `db:"clicklevel"`
	Maxcharges         int       `db:"maxcharges"`
	Casttime           int       `db:"casttime"`
	Recastdelay        int       `db:"recastdelay"`
	Recasttype         int       `db:"recasttype"`
	Clickunk5          int       `db:"clickunk5"`
	Clickname          int       `db:"clickname"`
	Clickunk7          int       `db:"clickunk7"`
	Proceffect         int       `db:"proceffect"`
	Proclevel2         int       `db:"proclevel2"`
	Proctype           int       `db:"proctype"`
	Proclevel          int       `db:"proclevel"`
	Prockunk1          int       `db:"prockunk1"`
	Procunk2           int       `db:"procunk2"`
	Procunk3           int       `db:"procunk3"`
	Procunk4           int       `db:"procunk4"`
	Procrate           int       `db:"procrate"`
	Procname           int       `db:"procname"`
	Procunk7           int       `db:"procunk7"`
	Worneffect         int       `db:"worneffect"`
	Wornlevel2         int       `db:"wornlevel2"`
	Worntype           int       `db:"worntype"`
	Wornlevel          int       `db:"wornlevel"`
	Wornunk1           int       `db:"wornunk1"`
	Wornunk2           int       `db:"wornunk2"`
	Wornunk3           int       `db:"wornunk3"`
	Wornunk4           int       `db:"wornunk4"`
	Wornunk5           int       `db:"wornunk5"`
	Wornname           int       `db:"wornname"`
	Wornunk7           int       `db:"wornunk7"`
	Focuseffect        int       `db:"focuseffect"`
	Focuslevel2        int       `db:"focuslevel2"`
	Focustype          int       `db:"focustype"`
	Focuslevel         int       `db:"focuslevel"`
	Focusunk1          int       `db:"focusunk1"`
	Focusunk2          int       `db:"focusunk2"`
	Focusunk3          int       `db:"focusunk3"`
	Focusunk4          int       `db:"focusunk4"`
	Focusunk5          int       `db:"focusunk5"`
	Focusname          int       `db:"focusname"`
	Focusunk7          int       `db:"focusunk7"`
	Scrolleffect       int       `db:"scrolleffect"`
	Scrolllevel2       int       `db:"scrolllevel2"`
	Scrolleffecttype   int       `db:"scrolleffecttype"`
	Scrolllevel        int       `db:"scrolllevel"`
	Scrollunk1         int       `db:"scrollunk1"`
	Scrollunk2         int       `db:"scrollunk2"`
	Scrollunk3         int       `db:"scrollunk3"`
	Scrollunk4         int       `db:"scrollunk4"`
	Scrollunk5         int       `db:"scrollunk5"`
	Scrollname         int       `db:"scrollname"`
	Scrollunk7         int       `db:"scrollunk7"`
	Bardeffect         int       `db:"bardeffect"`
	Bardlevel2         int       `db:"bardlevel2"`
	Bardeffecttype     int       `db:"bardeffecttype"`
	Bardlevel          int       `db:"bardlevel"`
	Bardunk1           int       `db:"bardunk1"`
	Bardunk2           int       `db:"bardunk2"`
	Bardunk3           int       `db:"bardunk3"`
	Bardunk4           int       `db:"bardunk4"`
	Bardunk5           int       `db:"bardunk5"`
	Bardname           int       `db:"bardname"`
	Bardunk7           int       `db:"bardunk7"`
	Mountunk5          int       `db:"mountunk5"`
	Blessingeffect     int       `db:"blessingeffect"`
	Blessingname       int       `db:"blessingname"`
	Mountname          int       `db:"mountname"`
	Mountunk7          int       `db:"mountunk7"`
	Blessinglevel2     int       `db:"blessinglevel2"`
	Blessingeffecttype int       `db:"blessingeffecttype"`
	Blessinglevel      int       `db:"blessinglevel"`
	Blessingunk1       int       `db:"blessingunk1"`
	Rightclickscriptid int       `db:"rightclickscriptid"`
	Questitemflag      int       `db:"questitemflag"`
	Powersourcecap     int       `db:"powersourcecap"`
	Purity             int       `db:"purity"`
	Epicitem           int       `db:"epicitem"`
	Backstabdmg        int       `db:"backstabdmg"`
	HeroicStr          int       `db:"heroic_str"`
	HeroicInt          int       `db:"heroic_int"`
	HeroicWis          int       `db:"heroic_wis"`
	HeroicAgi          int       `db:"heroic_agi"`
	HeroicDex          int       `db:"heroic_dex"`
	HeroicSta          int       `db:"heroic_sta"`
	HeroicCha          int       `db:"heroic_cha"`
	UNKNOWN29          int       `db:"UNKNOWN29"`
	Healamt            int       `db:"healamt"`
	Spelldmg           int       `db:"spelldmg"`
	Clairvoyance       int       `db:"clairvoyance"`
	UNKNOWN30          int       `db:"UNKNOWN30"`
	UNKNOWN31          int       `db:"UNKNOWN31"`
	UNKNOWN32          int       `db:"UNKNOWN32"`
	UNKNOWN33          int       `db:"UNKNOWN33"`
	UNKNOWN34          int       `db:"UNKNOWN34"`
	UNKNOWN35          int       `db:"UNKNOWN35"`
	UNKNOWN36          int       `db:"UNKNOWN36"`
	UNKNOWN37          int       `db:"UNKNOWN37"`
	Heirloom           int       `db:"heirloom"`
	Placeablebitfield  int       `db:"placeablebitfield"`
	UNKNOWN38          int       `db:"UNKNOWN38"`
	UNKNOWN39          int       `db:"UNKNOWN39"`
	UNKNOWN40          int       `db:"UNKNOWN40"`
	UNKNOWN41          int       `db:"UNKNOWN41"`
	UNKNOWN42          int       `db:"UNKNOWN42"`
	UNKNOWN43          int       `db:"UNKNOWN43"`
	UNKNOWN44          int       `db:"UNKNOWN44"`
	Placeablenpcname   int       `db:"placeablenpcname"`
	UNKNOWN46          int       `db:"UNKNOWN46"`
	UNKNOWN47          int       `db:"UNKNOWN47"`
	UNKNOWN48          int       `db:"UNKNOWN48"`
	UNKNOWN49          int       `db:"UNKNOWN49"`
	UNKNOWN50          int       `db:"UNKNOWN50"`
	UNKNOWN51          int       `db:"UNKNOWN51"`
	UNKNOWN52          int       `db:"UNKNOWN52"`
	UNKNOWN53          int       `db:"UNKNOWN53"`
	UNKNOWN54          int       `db:"UNKNOWN54"`
	UNKNOWN55          int       `db:"UNKNOWN55"`
	UNKNOWN56          int       `db:"UNKNOWN56"`
	UNKNOWN57          int       `db:"UNKNOWN57"`
	UNKNOWN58          int       `db:"UNKNOWN58"`
	UNKNOWN59          int       `db:"UNKNOWN59"`
	UNKNOWN60          int       `db:"UNKNOWN60"`
	UNKNOWN61          int       `db:"UNKNOWN61"`
	UNKNOWN62          int       `db:"UNKNOWN62"`
	UNKNOWN63          int       `db:"UNKNOWN63"`
	Collectible        int       `db:"collectible"`
	Nodestroy          int       `db:"nodestroy"`
	Nonpc              int       `db:"nonpc"`
	Nozone             int       `db:"nozone"`
	UNKNOWN68          int       `db:"UNKNOWN68"`
	UNKNOWN69          int       `db:"UNKNOWN69"`
	UNKNOWN70          int       `db:"UNKNOWN70"`
	UNKNOWN71          int       `db:"UNKNOWN71"`
	Noground           int       `db:"noground"`
	UNKNOWN73          int       `db:"UNKNOWN73"`
	Marketplace        int       `db:"marketplace"`
	Freestorage        int       `db:"freestorage"`
	UNKNOWN76          int       `db:"UNKNOWN76"`
	UNKNOWN77          int       `db:"UNKNOWN77"`
	UNKNOWN78          int       `db:"UNKNOWN78"`
	UNKNOWN79          int       `db:"UNKNOWN79"`
	Blessingunk2       int       `db:"blessingunk2"`
	Blessingunk3       int       `db:"blessingunk3"`
	Blessingunk4       int       `db:"blessingunk4"`
	Blessingunk5       int       `db:"blessingunk5"`
	Blessingunk7       int       `db:"blessingunk7"`
	Familiareffect     int       `db:"familiareffect"`
	Familiarlevel2     int       `db:"familiarlevel2"`
	Familiareffecttype int       `db:"familiareffecttype"`
	Familiarlevel      int       `db:"familiarlevel"`
	Familiarunk1       int       `db:"familiarunk1"`
	Familiarunk2       int       `db:"familiarunk2"`
	Familiarunk3       int       `db:"familiarunk3"`
	Familiarunk4       int       `db:"familiarunk4"`
	Familiarunk5       int       `db:"familiarunk5"`
	Familiarname       int       `db:"familiarname"`
	Familiarunk7       int       `db:"familiarunk7"`
	UNKNOWN80          int       `db:"UNKNOWN80"`
	Minluck            int       `db:"minluck"`
	Maxluck            int       `db:"maxluck"`
	Loreequippedgroup  int       `db:"loreequippedgroup"`
	Evoitem            int       `db:"evoitem"`
	Evoid              int       `db:"evoid"`
	Evolvl             int       `db:"evolvl"`
	Evomax             int       `db:"evomax"`
	Convertitem        int       `db:"convertitem"`
	Convertid          int       `db:"convertid"`
	Convertname        int       `db:"convertname"`
	Updated            time.Time `db:"updated"`
	Created            time.Time `db:"created"`
	Submitter          string    `db:"submitter"`
	Verified           time.Time `db:"verified"`
	Verifiedby         string    `db:"verifiedby"`
	Collectversion     string    `db:"collectversion"`
}
