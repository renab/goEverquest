package everquest

import "testing"

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
