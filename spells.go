package everquest

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Spell struct {
	Id                         int
	Name                       string
	Aeduration                 int
	Aerange                    int
	Affect_inanimate_object    int
	Ai_pt_bonus                int
	Ai_valid_targets           int
	Allow_spellscribe          int
	Anim_variation             int
	Attrib1                    int
	Attrib2                    int
	Attrib3                    int
	Attrib4                    int
	Attrib5                    int
	Attrib6                    int
	Attrib7                    int
	Attrib8                    int
	Attrib9                    int
	Attrib10                   int
	Attrib11                   int
	Attrib12                   int
	Attrib13                   int
	Attrib14                   int
	Attrib15                   int
	Attrib16                   int
	Attrib17                   int
	Attrib18                   int
	Attrib19                   int
	Attrib20                   int
	Attrib21                   int
	Attrib22                   int
	Attrib23                   int
	Attrib24                   int
	Attrib25                   int
	Attrib26                   int
	Attrib27                   int
	Attrib28                   int
	Attrib29                   int
	Attrib30                   int
	Attrib31                   int
	Attrib32                   int
	Attrib33                   int
	Attrib34                   int
	Attrib35                   int
	Attrib36                   int
	Attrib37                   int
	Attrib38                   int
	Attrib39                   int
	Attrib40                   int
	Attrib41                   int
	Autocast                   int
	Autocasttext               int
	Base1                      int
	Base2                      int
	Base10                     int
	Base11                     int
	Base12                     int
	Base13                     int
	Base14                     int
	Base15                     int
	Base16                     int
	Base17                     int
	Base19                     int
	Base20                     int
	Base21                     int
	Base22                     int
	Base24                     int
	Base26                     int
	Base27                     int
	Base28                     int
	Base29                     int
	Base2_1                    int
	Base2_10                   int
	Base2_11                   int
	Base2_12                   int
	Base2_13                   int
	Base2_14                   int
	Base2_15                   int
	Base2_16                   int
	Base2_17                   int
	Base2_18                   int
	Base2_19                   int
	Base2_2                    int
	Base2_20                   int
	Base2_21                   int
	Base2_22                   int
	Base2_23                   int
	Base2_24                   int
	Base2_25                   int
	Base2_26                   int
	Base2_28                   int
	Base2_29                   int
	Base2_3                    int
	Base2_30                   int
	Base2_31                   int
	Base2_32                   int
	Base2_33                   int
	Base2_35                   int
	Base2_36                   int
	Base2_37                   int
	Base2_38                   int
	Base2_39                   int
	Base2_4                    int
	Base2_40                   int
	Base2_41                   int
	Base2_5                    int
	Base2_6                    int
	Base2_7                    int
	Base2_8                    int
	Base2_9                    int
	Base3                      int
	Base4                      int
	Base5                      int
	Base6                      int
	Base7                      int
	Base8                      int
	Base9                      int
	Base18                     int
	Base23                     int
	Base25                     int
	Base2_27                   int
	Base2_34                   int
	Base30                     int
	Base31                     int
	Base32                     int
	Base33                     int
	Base34                     int
	Base35                     int
	Base36                     int
	Base37                     int
	Base38                     int
	Base39                     int
	Base40                     int
	Base41                     int
	Base_effects_focus_offset  int
	Base_effects_focus_slope   int
	Berlevel                   int
	Bonushate                  int
	Bookicon                   int
	Brdlevel                   int
	Bstlevel                   int
	Bypass_regen_check         int
	Calc1                      int
	Calc2                      int
	Calc3                      int
	Calc4                      int
	Calc5                      int
	Calc6                      int
	Calc7                      int
	Calc8                      int
	Calc9                      int
	Calc10                     int
	Calc11                     int
	Calc12                     int
	Calc13                     int
	Calc14                     int
	Calc15                     int
	Calc16                     int
	Calc17                     int
	Calc18                     int
	Calc19                     int
	Calc20                     int
	Calc21                     int
	Calc22                     int
	Calc23                     int
	Calc24                     int
	Calc25                     int
	Calc26                     int
	Calc27                     int
	Calc28                     int
	Calc29                     int
	Calc30                     int
	Calc31                     int
	Calc32                     int
	Calc33                     int
	Calc34                     int
	Calc35                     int
	Calc36                     int
	Calc37                     int
	Calc38                     int
	Calc39                     int
	Calc40                     int
	Calc41                     int
	Can_cast_in_combat         int
	Can_mgb                    int
	Cancelonsit                int
	Cast_not_standing          int
	Castinganim                int
	Castingtime                float64
	Castmsg1                   string
	Castmsg2                   string
	Castmsg3                   string
	Castmsg4                   string
	Castmsg5                   string
	Castrestriction            int
	Classes                    string
	Clrlevel                   int
	Cone_end_angle             int
	Cone_start_angle           int
	Deities                    int
	Deletable                  int
	Desc1                      string
	Desc2                      string
	Desc3                      string
	Desc4                      string
	Desc5                      string
	Desc6                      string
	Desc7                      string
	Desc8                      string
	Desc9                      string
	Desc10                     string
	Desc11                     string
	Desc12                     string
	Desc13                     string
	Desc14                     string
	Desc15                     string
	Desc16                     string
	Desc17                     string
	Desc18                     string
	Desc19                     string
	Desc20                     string
	Desc21                     string
	Desc22                     string
	Desc23                     string
	Desc24                     string
	Desc25                     string
	Desc26                     string
	Desc27                     string
	Desc28                     string
	Desc29                     string
	Desc30                     string
	Desc31                     string
	Desc32                     string
	Desc33                     string
	Desc34                     string
	Desc35                     string
	Desc36                     string
	Desc37                     string
	Desc38                     string
	Desc39                     string
	Desc40                     string
	Desc41                     string
	Descnum                    int
	Distance_mod_close_dist    int
	Distance_mod_close_mult    int
	Distance_mod_far_dist      int
	Distance_mod_far_mult      int
	Dotstackingexempt          int
	Drulevel                   int
	Duration                   int
	Duration_particle_effect   int
	Durationformula            int
	Durationfreeze             int
	Durationtext               string
	Enclevel                   int
	Endurance_cost             int
	Enduranceupkeep            int
	Environment                int
	Expansion                  string
	Extra                      int
	Feedbackable               int
	Fizzleadj                  int
	Fizzletime                 int
	Foci                       int
	Focus1                     int
	Focus2                     int
	Focus3                     int
	Focus4                     int
	Focusitems                 int
	Gemicon                    int
	Hateamount                 int
	Is_beta_only               int
	Is_skill                   int
	Lighttype                  int
	Location                   string
	Maglevel                   int
	Manacost                   int
	Max1                       int
	Max2                       int
	Max3                       int
	Max4                       int
	Max5                       int
	Max6                       int
	Max7                       int
	Max8                       int
	Max9                       int
	Max10                      int
	Max11                      int
	Max12                      int
	Max13                      int
	Max14                      int
	Max15                      int
	Max16                      int
	Max17                      int
	Max18                      int
	Max19                      int
	Max20                      int
	Max21                      int
	Max22                      int
	Max23                      int
	Max24                      int
	Max25                      int
	Max26                      int
	Max27                      int
	Max28                      int
	Max29                      int
	Max30                      int
	Max31                      int
	Max32                      int
	Max33                      int
	Max34                      int
	Max35                      int
	Max36                      int
	Max37                      int
	Max38                      int
	Max39                      int
	Max40                      int
	Max41                      int
	Max_hits_type              int
	Max_resist                 int
	Maxduration                int
	Maxtargets                 int
	Min_range                  int
	Min_resist                 int
	Minduration                int
	Minlevel                   int
	Mnklevel                   int
	Neclevel                   int
	No_buff_block              int
	No_detrimental_spell_aggro int
	No_heal_damage_item_mod    int
	No_npc_los                 int
	No_overwrite               int
	No_partial_save            int
	No_remove                  int
	No_resist                  int
	Nodispell                  int
	Not_focusable              int
	Not_shown_to_player        int
	Npc_category               int
	Npc_no_cast                int
	Npc_usefulness             int
	Numhits                    int
	Only_during_fast_regen     int
	Outofcombat                int
	Override_crit_chance       int
	Pallevel                   int
	Pcnpc_only_flag            int
	Persistdeath               int
	Primary_category           int
	Pushback                   float64
	Pushup                     int
	Pvp_duration               int
	Pvp_duration_cap           int
	Pvpresistbase              int
	Pvpresistcalc              int
	Pvpresistcap               int
	Range                      int
	Reagentcount1              int
	Reagentcount2              int
	Reagentcount3              int
	Reagentcount4              int
	Reagentid1                 int
	Reagentid2                 int
	Reagentid3                 int
	Reagentid4                 int
	Reagents                   int
	Recasttime                 float64
	Reflectable                int
	Resist                     string
	Resist_cap                 int
	Resist_per_level           int
	Resistadj                  int
	Rnglevel                   int
	Roglevel                   int
	Secondary_category         int
	Secondary_category2        int
	Shdlevel                   int
	Shmlevel                   int
	Shortbuff                  int
	Show_dot_message           int
	Show_wear_off_message      int
	Skill                      string
	Small_targets_only         int
	Sneak_attack               int
	Songcap                    int
	Source                     string
	Spaindex                   int
	Spell_class                int
	Spell_group_rank           int
	Spell_recourse_type        int
	Spell_subclass             int
	Spell_subgroup             int
	Spellanim                  int
	Spellgroup                 int
	Spellicon                  int
	Spelltype                  string
	Stacks_with_self           int
	Targetanim                 int
	Targetrestriction          int
	Targettype                 string
	Targname                   string
	Timeofday                  string
	Timer                      int
	Traveltype                 int
	Uninterruptable            int
	Unknown1                   int
	Unknown2                   int
	Updated                    time.Time
	Uses_persistant_particles  int
	Viral_range                int
	Viral_targets              int
	Viral_timer                int
	Warlevel                   int
	Wizlevel                   int
}

func (s *Spell) Load(data ...string) {
	var x int
	var d int
	d, _ = strconv.Atoi(data[x])
	s.Id = d
	x++
	s.Name = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	s.Aeduration = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Aerange = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Affect_inanimate_object = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Ai_pt_bonus = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Ai_valid_targets = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Allow_spellscribe = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Anim_variation = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib5 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib6 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib8 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib9 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib10 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib11 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib12 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib13 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib14 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib15 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib16 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib17 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib18 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib19 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib20 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib21 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib22 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib23 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib24 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib25 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib26 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib27 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib28 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib29 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib30 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib31 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib32 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib33 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib34 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib35 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib36 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib37 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib38 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib39 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib40 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Attrib41 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Autocast = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Autocasttext = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base10 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base11 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base12 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base13 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base14 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base15 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base16 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base17 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base19 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base20 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base21 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base22 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base24 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base26 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base27 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base28 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base29 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_10 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_11 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_12 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_13 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_14 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_15 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_16 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_17 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_18 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_19 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_20 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_21 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_22 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_23 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_24 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_25 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_26 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_28 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_29 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_30 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_31 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_32 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_33 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_35 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_36 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_37 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_38 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_39 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_40 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_41 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_5 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_6 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_8 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_9 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base5 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base6 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base8 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base9 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base18 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base23 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base25 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_27 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base2_34 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base30 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base31 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base32 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base33 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base34 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base35 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base36 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base37 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base38 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base39 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base40 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base41 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base_effects_focus_offset = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Base_effects_focus_slope = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Berlevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Bonushate = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Bookicon = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Brdlevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Bstlevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Bypass_regen_check = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc5 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc6 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc8 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc9 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc10 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc11 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc12 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc13 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc14 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc15 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc16 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc17 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc18 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc19 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc20 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc21 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc22 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc23 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc24 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc25 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc26 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc27 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc28 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc29 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc30 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc31 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc32 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc33 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc34 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc35 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc36 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc37 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc38 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc39 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc40 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Calc41 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Can_cast_in_combat = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Can_mgb = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Cancelonsit = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Cast_not_standing = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Castinganim = d
	x++
	flt, _ := strconv.ParseFloat(data[x], 64)
	s.Castingtime = flt
	x++
	s.Castmsg1 = data[x]
	x++
	s.Castmsg2 = data[x]
	x++
	s.Castmsg3 = data[x]
	x++
	s.Castmsg4 = data[x]
	x++
	s.Castmsg5 = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	s.Castrestriction = d
	x++
	s.Classes = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	s.Clrlevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Cone_end_angle = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Cone_start_angle = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Deities = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Deletable = d
	x++
	s.Desc1 = data[x]
	x++
	s.Desc2 = data[x]
	x++
	s.Desc3 = data[x]
	x++
	s.Desc4 = data[x]
	x++
	s.Desc5 = data[x]
	x++
	s.Desc6 = data[x]
	x++
	s.Desc7 = data[x]
	x++
	s.Desc8 = data[x]
	x++
	s.Desc9 = data[x]
	x++
	s.Desc10 = data[x]
	x++
	s.Desc11 = data[x]
	x++
	s.Desc12 = data[x]
	x++
	s.Desc13 = data[x]
	x++
	s.Desc14 = data[x]
	x++
	s.Desc15 = data[x]
	x++
	s.Desc16 = data[x]
	x++
	s.Desc17 = data[x]
	x++
	s.Desc18 = data[x]
	x++
	s.Desc19 = data[x]
	x++
	s.Desc20 = data[x]
	x++
	s.Desc21 = data[x]
	x++
	s.Desc22 = data[x]
	x++
	s.Desc23 = data[x]
	x++
	s.Desc24 = data[x]
	x++
	s.Desc25 = data[x]
	x++
	s.Desc26 = data[x]
	x++
	s.Desc27 = data[x]
	x++
	s.Desc28 = data[x]
	x++
	s.Desc29 = data[x]
	x++
	s.Desc30 = data[x]
	x++
	s.Desc31 = data[x]
	x++
	s.Desc32 = data[x]
	x++
	s.Desc33 = data[x]
	x++
	s.Desc34 = data[x]
	x++
	s.Desc35 = data[x]
	x++
	s.Desc36 = data[x]
	x++
	s.Desc37 = data[x]
	x++
	s.Desc38 = data[x]
	x++
	s.Desc39 = data[x]
	x++
	s.Desc40 = data[x]
	x++
	s.Desc41 = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	s.Descnum = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Distance_mod_close_dist = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Distance_mod_close_mult = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Distance_mod_far_dist = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Distance_mod_far_mult = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Dotstackingexempt = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Drulevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Duration = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Duration_particle_effect = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Durationformula = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Durationfreeze = d
	x++
	s.Durationtext = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	s.Enclevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Endurance_cost = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Enduranceupkeep = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Environment = d
	x++
	s.Expansion = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	s.Extra = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Feedbackable = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Fizzleadj = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Fizzletime = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Foci = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Focus1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Focus2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Focus3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Focus4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Focusitems = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Gemicon = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Hateamount = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Is_beta_only = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Is_skill = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Lighttype = d
	x++
	s.Location = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	s.Maglevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Manacost = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max5 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max6 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max7 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max8 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max9 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max10 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max11 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max12 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max13 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max14 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max15 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max16 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max17 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max18 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max19 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max20 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max21 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max22 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max23 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max24 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max25 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max26 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max27 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max28 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max29 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max30 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max31 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max32 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max33 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max34 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max35 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max36 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max37 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max38 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max39 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max40 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max41 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max_hits_type = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Max_resist = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Maxduration = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Maxtargets = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Min_range = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Min_resist = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Minduration = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Minlevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Mnklevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Neclevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.No_buff_block = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.No_detrimental_spell_aggro = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.No_heal_damage_item_mod = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.No_npc_los = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.No_overwrite = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.No_partial_save = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.No_remove = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.No_resist = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Nodispell = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Not_focusable = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Not_shown_to_player = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Npc_category = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Npc_no_cast = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Npc_usefulness = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Numhits = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Only_during_fast_regen = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Outofcombat = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Override_crit_chance = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Pallevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Pcnpc_only_flag = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Persistdeath = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Primary_category = d
	x++
	flt, _ = strconv.ParseFloat(data[x], 64)
	s.Pushback = flt
	x++
	d, _ = strconv.Atoi(data[x])
	s.Pushup = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Pvp_duration = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Pvp_duration_cap = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Pvpresistbase = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Pvpresistcalc = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Pvpresistcap = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Range = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Reagentcount1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Reagentcount2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Reagentcount3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Reagentcount4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Reagentid1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Reagentid2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Reagentid3 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Reagentid4 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Reagents = d
	x++
	flt, _ = strconv.ParseFloat(data[x], 64)
	s.Recasttime = flt
	x++
	d, _ = strconv.Atoi(data[x])
	s.Reflectable = d
	x++
	s.Resist = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	s.Resist_cap = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Resist_per_level = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Resistadj = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Rnglevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Roglevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Secondary_category = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Secondary_category2 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Shdlevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Shmlevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Shortbuff = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Show_dot_message = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Show_wear_off_message = d
	x++
	s.Skill = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	s.Small_targets_only = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Sneak_attack = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Songcap = d
	x++
	s.Source = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	s.Spaindex = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Spell_class = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Spell_group_rank = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Spell_recourse_type = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Spell_subclass = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Spell_subgroup = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Spellanim = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Spellgroup = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Spellicon = d
	x++
	s.Spelltype = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	s.Stacks_with_self = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Targetanim = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Targetrestriction = d
	x++
	s.Targettype = data[x]
	x++
	s.Targname = data[x]
	x++
	s.Timeofday = data[x]
	x++
	d, _ = strconv.Atoi(data[x])
	s.Timer = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Traveltype = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Uninterruptable = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Unknown1 = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Unknown2 = d
	x++
	updated, _ := time.Parse("2006-01-02 15:04:05", data[x])
	s.Updated = updated
	x++
	d, _ = strconv.Atoi(data[x])
	s.Uses_persistant_particles = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Viral_range = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Viral_targets = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Viral_timer = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Warlevel = d
	x++
	d, _ = strconv.Atoi(data[x])
	s.Wizlevel = d
	x++
}

type SpellDB struct {
	byID   map[int]Spell  // primary way to find an spell is by it's ID #
	byName map[string]int // Used to fast lookup ID by name (there may be duplicates, and is not recommended)
}

func (db *SpellDB) LoadFromFile(file string, Info *log.Logger) error {
	db.byID = make(map[int]Spell)
	db.byName = make(map[string]int)
	psvfile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer psvfile.Close()

	// r := bufio.NewReader(psvfile)
	r := csv.NewReader(bufio.NewReader(psvfile))
	r.LazyQuotes = true

	// Iterate through the records
	headerSkipped := false
	var spellCount int
	// Iterate through the records
	for {
		spellCount++
		if headerSkipped && spellCount == 1 {
			continue
		}
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		var spell Spell
		spell.Load(record...)
		db.byID[spell.Id] = spell
		db.byName[strings.ToLower(spell.Name)] = spell.Id
	}
	Info.Printf("Loaded %d spells\n", spellCount)
	return nil
}

// FindIDByName does an spell lookup by the spell name, returns -1 if not found
func (db *SpellDB) FindIDByName(name string) (int, error) {
	lower := strings.ToLower(name)
	if val, ok := db.byName[lower]; ok {
		return val, nil
	}
	return -1, errors.New("cannot find spell id with name: " + name)
}

// GetSpellByID returns an spell given its ID, returns an empty struct if not found
func (db *SpellDB) GetSpellByID(id int) (Spell, error) {
	if val, ok := db.byID[id]; ok {
		return val, nil
	}
	return Spell{}, errors.New("cannot find spell by id")
}

// SearchSpellsByName will do a long search to find spells containing the input value
func (db *SpellDB) SearchSpellsByName(name string) []Spell {
	var results []Spell
	for _, spell := range db.byID {
		if strings.Contains(spell.Name, name) {
			results = append(results, spell)
		}
	}
	return results
}

func (s *Spell) GetClasses() []string {
	classLevels := strings.Split(s.Classes, " ")
	var classes []string
	for _, class := range classLevels {
		if len(class) >= 3 {
			fullClass, err := ShortClassNameToFull(class[:3])
			if err != nil {
				continue
			}
			classes = append(classes, fullClass)
		}
	}
	return classes
}

func (s *Spell) ClassCanUse(class string) bool {
	classes := s.GetClasses()
	for _, usable := range classes {
		if class == usable {
			return true
		}
	}
	return false
}

func (db *SpellDB) GetClassSpells(class string) []Spell {
	var results []Spell
	for _, spell := range db.byID {
		if spell.ClassCanUse(class) {
			results = append(results, spell)
		}
	}
	return results
}
