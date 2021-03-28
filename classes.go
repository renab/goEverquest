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
	}
	return []string{}
}
