package everquest

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// id,name,aeduration,aerange,affect_inanimate_object,ai_pt_bonus,ai_valid_targets,allow_spellscribe
// ,anim_variation,attrib1,attrib2,attrib3,attrib4,attrib5,attrib6,attrib7,attrib8,attrib9,attrib10
// ,attrib11,attrib12,attrib13,attrib14,attrib15,attrib16,attrib17,attrib18,attrib19,attrib20,attrib21
// ,attrib22,attrib23,attrib24,attrib25,attrib26,attrib27,attrib28,attrib29,attrib30,attrib31,attrib32
// ,attrib33,attrib34,attrib35,attrib36,attrib37,attrib38,attrib39,attrib40,attrib41,autocast,autocasttext
// ,base1,base2,base10,base11,base12,base13,base14,base15,base16,base17,base19,base20,base21,base22,base24
// ,base26,base27,base28,base29,base2_1,base2_10,base2_11,base2_12,base2_13,base2_14,base2_15,base2_16,base2_17
// ,base2_18,base2_19,base2_2,base2_20,base2_21,base2_22,base2_23,base2_24,base2_25,base2_26,base2_28,base2_29
// ,base2_3,base2_30,base2_31,base2_32,base2_33,base2_35,base2_36,base2_37,base2_38,base2_39,base2_4,base2_40,base2_41
// ,base2_5,base2_6,base2_7,base2_8,base2_9,base3,base4,base5,base6,base7,base8,base9,base18,base23,base25,base2_27
// ,base2_34,base30,base31,base32,base33,base34,base35,base36,base37,base38,base39,base40,base41,base_effects_focus_offset
// ,base_effects_focus_slope,berlevel,bonushate,bookicon,brdlevel,bstlevel,bypass_regen_check,calc1,calc2
// ,calc3,calc4,calc5,calc6,calc7,calc8,calc9,calc10,calc11,calc12,calc13,calc14,calc15,calc16,calc17
// ,calc18,calc19,calc20,calc21,calc22,calc23,calc24,calc25,calc26,calc27,calc28,calc29,calc30,calc31
// ,calc32,calc33,calc34,calc35,calc36,calc37,calc38,calc39,calc40,calc41,can_cast_in_combat,can_mgb,cancelonsit
// ,cast_not_standing,castinganim,castingtime,castmsg1,castmsg2,castmsg3,castmsg4,castmsg5,castrestriction
// ,classes,clrlevel,cone_end_angle,cone_start_angle,deities,deletable,desc1,desc2,desc3,desc4,desc5,desc6
// ,desc7,desc8,desc9,desc10,desc11,desc12,desc13,desc14,desc15,desc16,desc17,desc18,desc19,desc20,desc21
// ,desc22,desc23,desc24,desc25,desc26,desc27,desc28,desc29,desc30,desc31,desc32,desc33,desc34,desc35,desc36
// ,desc37,desc38,desc39,desc40,desc41,descnum,distance_mod_close_dist,distance_mod_close_mult,distance_mod_far_dist
// ,distance_mod_far_mult,dotstackingexempt,drulevel,duration,duration_particle_effect,durationformula,durationfreeze
// ,durationtext,enclevel,endurance_cost,enduranceupkeep,environment,expansion,extra,feedbackable,fizzleadj,fizzletime
// ,foci,focus1,focus2,focus3,focus4,focusitems,gemicon,hateamount,is_beta_only,is_skill,lighttype,location,maglevel
// ,manacost,max1,max2,max3,max4,max5,max6,max7,max8,max9,max10,max11,max12,max13,max14,max15,max16,max17,max18,max19
// ,max20,max21,max22,max23,max24,max25,max26,max27,max28,max29,max30,max31,max32,max33,max34,max35,max36,max37,max38
// ,max39,max40,max41,max_hits_type,max_resist,maxduration,maxtargets,min_range,min_resist,minduration,minlevel,mnklevel
// ,neclevel,no_buff_block,no_detrimental_spell_aggro,no_heal_damage_item_mod,no_npc_los,no_overwrite,no_partial_save
// ,no_remove,no_resist,nodispell,not_focusable,not_shown_to_player,npc_category,npc_no_cast,npc_usefulness,numhits
// ,only_during_fast_regen,outofcombat,override_crit_chance,pallevel,pcnpc_only_flag,persistdeath,primary_category
// ,pushback,pushup,pvp_duration,pvp_duration_cap,pvpresistbase,pvpresistcalc,pvpresistcap,range,reagentcount1,reagentcount2
// ,reagentcount3,reagentcount4,reagentid1,reagentid2,reagentid3,reagentid4,reagents,recasttime,reflectable,resist,resist_cap
// ,resist_per_level,resistadj,rnglevel,roglevel,secondary_category,secondary_category2,shdlevel,shmlevel,shortbuff
// ,show_dot_message,show_wear_off_message,skill,small_targets_only,sneak_attack,songcap,source,spaindex,spell_class
// ,spell_group_rank,spell_recourse_type,spell_subclass,spell_subgroup,spellanim,spellgroup,spellicon,spelltype,stacks_with_self
// ,targetanim,targetrestriction,targettype,targname,timeofday,timer,traveltype,uninterruptable,unknown1,unknown2,updated
// ,uses_persistant_particles,viral_range,viral_targets,viral_timer,warlevel,wizlevel

// /outputfile Spellbook defaults to base eq folder Charactername_server-Spellbook.txt
// we should pull spellbook when we see it was dumped

// Spellbook contains a slice of all known spells
type Spellbook struct {
	Spells []SpellBookEntry
}

// Spell defines a spellbook entry
type SpellBookEntry struct {
	Level int
	Name  string
}

// LoadFromPath loads a standard everquest spell dump
func (s *Spellbook) LoadFromPath(path string, Err *log.Logger) error {
	// Open the file
	tsvfile, err := os.Open(path)
	if err != nil {
		return err
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
			Err.Printf("Error converting level to int - Level: %s Name: %s\n", record[0], record[1])
			continue
		}
		spell := SpellBookEntry{
			Level: level,
			Name:  record[1],
		}
		s.Spells = append(s.Spells, spell)
	}
	return nil
}
