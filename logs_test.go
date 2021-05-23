package everquest

import (
	"regexp"
	"testing"
	"time"
)

func TestGetChannel(t *testing.T) {
	chanMSG := `Ravnor tells Von_parses:5, 'Lord Vyemm in 485s, 249k AH | Kaijin 22042 AH | Voltha 18485 AH | Patchouli 17664 AH | Silvaefar 17092 AH | Bunzz 17062 AH | Milliardo 15491 AH | Scylla 14543 AH | Vinadru 13840 AH | Porrt 13074 AH | Blossom 12927 AH | Impulse 12661 AH | Clearwater 10482 AH | Stony 9390 AH | Sacristan 8511 AH | Banis 7346 AH'`
	sysMSG := `Patchouli winces.`
	guildMSG := `Zobac tells the guild, 'Gratz Banis and Guzz!! :P'`
	tellMSG := `Zortax tells you, 'no idea , havent spoken to either of them in a while .  I assume holidays , but you never know'`
	raidMSG := `Ryze tells the raid, 'anyone else i missed let me know'`
	sayMSG := `Bunzz says, 'Hail, Charybdis'`
	groupMSG := `Glooping tells the group, 'hey talen can you run a melee haste on this fight possibly? or do you not go in with Haye?'`
	aucMSG := `Scylla auctions, 'Complete HEAL <<< Krotax >>>'`

	if getChannel(aucMSG) != "auction" {
		t.Fatalf("Error reading auction channel")
	}
	if getChannel(groupMSG) != "group" {
		t.Fatalf("Error reading group channel\n")
	}
	if getChannel(sayMSG) != "say" {
		t.Fatalf("Error reading say channel\n")
	}
	if getChannel(raidMSG) != "raid" {
		t.Fatalf("Error reading raid channel\n")
	}
	if getChannel(tellMSG) != "tell" {
		t.Fatalf("Error reading tell channel\n")
	}
	if getChannel(guildMSG) != "guild" {
		t.Fatalf("Error reading guild channel\n")
	}
	if getChannel(sysMSG) != "system" {
		t.Fatalf("Error reading system channel\n")
	}
	if getChannel(chanMSG) != "Von_parses" {
		t.Fatalf("Error reading Von_parses channel\n%s\nshows as\n%s", chanMSG, getChannel(chanMSG))
	}
}

func TestReadLogLine(t *testing.T) {
	testLog := `[Sat Jan 02 20:44:08 2021] Destrod tells the guild, 'takings bids on Shawl of Perception pst bids close in 2mins'`
	r, _ := regexp.Compile(EQBaseLogLine)
	results := r.FindAllStringSubmatch(testLog, -1)
	eqlog := readLogLine(results)
	testTime := time.Date(2021, time.January, 2, 20, 44, 8, 0, time.Local)
	if eqlog.Channel != "guild" {
		t.Fatalf("Error parsing channel")
	}
	if eqlog.Msg != "Destrod tells the guild, 'takings bids on Shawl of Perception pst bids close in 2mins'" {
		t.Fatalf("Error parsing msg")
	}
	if eqlog.Source != "Destrod" {
		t.Fatalf("Error parsing source")
	}
	if !eqlog.T.Equal(testTime) {
		t.Fatalf("Error parsing time")
	}
}

func TestEQTimeConv(t *testing.T) {
	testTime := "Sat May 23 20:44:08 2021"
	conv := eqTimeConv(testTime)
	passTime := time.Date(2021, time.May, 23, 20, 44, 8, 0, time.Local)
	if !conv.Equal(passTime) {
		t.Fatalf("Error parsing eq time to time.Time: %s vs %s", conv.String(), passTime.String())
	}
}

func TestGetSource(t *testing.T) {
	testMSG := "Destrod tells the guild, 'takings bids on Shawl of Perception pst bids close in 2mins'"
	if getSource(testMSG) != "Destrod" {
		t.Fatalf("Error getting message source\n")
	}
}

func TestGetLogPath(t *testing.T) {
	player := "Mortimus"
	server := "aradune"
	basePath := "C:/Everquest"
	success := "C:/Everquest/Logs/eqlog_Mortimus_aradune.txt"
	if GetLogPath(player, server, basePath) != success {
		t.Fatalf("Error getting log path of player")
	}
}
