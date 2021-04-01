package everquest

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
func GetClassesByRole(role string) []string {
	switch role {
	case "All":
		return []string{"Bard", "Beastlord", "Berserker", "Cleric", "Druid", "Enchanter", "Magician", "Monk", "Necromancer", "Paladin", "Ranger", "Rogue", "Shadow Knight", "Shaman", "Warrior", "Wizard"}
	case "Tank":
		return []string{"Paladin", "Shadow Knight", "Warrior"} // Ranger?
	case "Priest":
		return []string{"Cleric", "Druid", "Shaman"}
	case "DPS":
		return []string{"Beastlord", "Berserker", "Magician", "Monk", "Necromancer", "Ranger", "Rogue", "Wizard"}
	case "CC":
		return []string{"Bard", "Enchanter"}
	case "Lockpick":
		return []string{"Bard", "Rogue"}
	case "Undead":
		return []string{"Cleric", "Necromancer", "Paladin"}
	case "Puller":
		return []string{"Bard", "Monk", "Necromancer"}
	case "Deathtouch":
		return []string{"Ranger"}
	case "Cloth":
		return []string{"Enchanter", "Magician", "Necromancer", "Wizard"}
	case "Leather":
		return []string{"Beastlord", "Druid", "Monk"}
	case "Chain":
		return []string{"Berserker", "Ranger", "Rogue", "Shaman"}
	case "Plate":
		return []string{"Bard", "Cleric", "Paladin", "Shadow Knight", "Warrior"}
	}
	return []string{}
}
