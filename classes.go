package everquest

import "errors"

/*
Bard
Beastlord
Berserker
Cleric
Druid
Enchanter
Magician
Monk
Necromancer
Paladin
Ranger
Rogue
Shadow Knight
Shaman
Warrior
Wizard
*/

// GetClassesByRole will take a role like Tank and return Warrior/Paladin/Shadow Knight
func GetClassesByRole(role string) ([]string, error) {
	switch role {
	case "All":
		return []string{"Bard", "Beastlord", "Berserker", "Cleric", "Druid", "Enchanter", "Magician", "Monk", "Necromancer", "Paladin", "Ranger", "Rogue", "Shadow Knight", "Shaman", "Warrior", "Wizard"}, nil
	case "Tank":
		return []string{"Paladin", "Shadow Knight", "Warrior"}, nil // Ranger?
	case "Priest":
		return []string{"Cleric", "Druid", "Shaman"}, nil
	case "DPS":
		return []string{"Beastlord", "Berserker", "Magician", "Monk", "Necromancer", "Ranger", "Rogue", "Wizard"}, nil
	case "CC":
		return []string{"Bard", "Enchanter"}, nil
	case "Lockpick":
		return []string{"Bard", "Rogue"}, nil
	case "Undead":
		return []string{"Cleric", "Necromancer", "Paladin"}, nil
	case "Puller":
		return []string{"Bard", "Monk", "Necromancer"}, nil
	case "Deathtouch":
		return []string{"Ranger"}, nil
	case "Cloth":
		return []string{"Enchanter", "Magician", "Necromancer", "Wizard"}, nil
	case "Leather":
		return []string{"Beastlord", "Druid", "Monk"}, nil
	case "Chain":
		return []string{"Berserker", "Ranger", "Rogue", "Shaman"}, nil
	case "Plate":
		return []string{"Bard", "Cleric", "Paladin", "Shadow Knight", "Warrior"}, nil
	case "Pet":
		return []string{"Beastlord", "Magician", "Necromancer"}, nil
	}
	return []string{}, errors.New("unknown role " + role)
}

func ShortClassNameToFull(sName string) (string, error) {
	switch sName {
	case "BRD":
		return "Bard", nil
	case "BST":
		return "Beastlord", nil
	case "BER":
		return "Berserker", nil
	case "CLR":
		return "Cleric", nil
	case "DRU":
		return "Druid", nil
	case "ENC":
		return "Enchanter", nil
	case "MAG":
		return "Magician", nil
	case "MNK":
		return "Monk", nil
	case "NEC":
		return "Necromancer", nil
	case "PAL":
		return "Paladin", nil
	case "RNG":
		return "Ranger", nil
	case "ROG":
		return "Rogue", nil
	case "SHD":
		return "Shadow Knight", nil
	case "SHM":
		return "Shaman", nil
	case "WAR":
		return "Warrior", nil
	case "WIZ":
		return "Wizard", nil
	default:
		return "Unknown", errors.New("unknown class " + sName)
	}
}
