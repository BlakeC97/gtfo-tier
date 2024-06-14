package tiering

type Weapon string
type Situation string

const (
	ARBALIST            Weapon = "Arbalist Machine Gun"
	BURST_CANNON        Weapon = "Burst Cannon"
	CHOKE_MOD           Weapon = "Choke Mod Shotgun"
	COMBAT_SHOTGUN      Weapon = "Combat Shotgun"
	HEAVY_ASSAULT_RIFLE Weapon = "Heavy Assault Rifle"
	HEL_GUN             Weapon = "HEL Gun"
	HEL_RIFLE           Weapon = "HEL Rifle"
	HIGH_CAL            Weapon = "High Caliber Pistol"
	PRECISION_RIFLE     Weapon = "Precision Rifle"
	REVOLVER            Weapon = "Revolver"
	SCATTERGUN          Weapon = "Scattergun"
	SHORT_RIFLE         Weapon = "Short Rifle"
	SHOTGUN             Weapon = "Shotgun"
	SNIPER              Weapon = "Sniper"
	VERUTA              Weapon = "Veruta Machine Gun"
)

const (
	BABY           Situation = "Baby Waves"
	BIG_CHARGERS   Situation = "Big Chargers"
	CHARGERS       Situation = "Chargers"
	CHARGER_SCOUTS Situation = "Charger Scouts"
	CQC            Situation = "Close Quarter Waveclear"
	FLYERS         Situation = "Flyers"
	GENERIC        Situation = "Generic Waveclear"
	GIANTS         Situation = "Giants/Hybrids"
	KRAKEN         Situation = "Kraken"
	MOTHERS        Situation = "Mothers"
	NIGHTMARES     Situation = "Nightmares"
	P_MOTHER       Situation = "PMother"
	REACTOR        Situation = "Reactor Holds"
	SCOUTS         Situation = "Scouts/Shadow Scouts"
	SNATCHERS      Situation = "Snatchers"
	TANK           Situation = "Tanks"
)

var (
	AllWeapons = []Weapon{
		ARBALIST,
		BURST_CANNON,
		CHOKE_MOD,
		COMBAT_SHOTGUN,
		HEAVY_ASSAULT_RIFLE,
		HEL_GUN,
		HEL_RIFLE,
		HIGH_CAL,
		PRECISION_RIFLE,
		REVOLVER,
		SCATTERGUN,
		SHORT_RIFLE,
		SHOTGUN,
		SNIPER,
		VERUTA,
	}
	AllSituations = []Situation{
		BABY,
		BIG_CHARGERS,
		CHARGERS,
		CHARGER_SCOUTS,
		CQC,
		FLYERS,
		GENERIC,
		GIANTS,
		KRAKEN,
		MOTHERS,
		NIGHTMARES,
		P_MOTHER,
		REACTOR,
		SCOUTS,
		SNATCHERS,
		TANK,
	}
)

type TierList struct {
	Best   []Weapon
	Good   []Weapon
	Decent []Weapon
	Niche  []Weapon
	DoNot  []Weapon
	Notes  []string
}

func TierLists() map[Situation]TierList {
	return map[Situation]TierList{
		BABY: {
			Best:   []Weapon{HEAVY_ASSAULT_RIFLE, VERUTA, SHOTGUN},
			Good:   []Weapon{COMBAT_SHOTGUN, HEL_GUN, ARBALIST},
			Decent: []Weapon{SHORT_RIFLE, PRECISION_RIFLE},
			Niche:  []Weapon{HEL_RIFLE, SCATTERGUN},
			DoNot:  []Weapon{REVOLVER, HIGH_CAL, BURST_CANNON, CHOKE_MOD, SNIPER},
			Notes: []string{
				"CHRIS: MENTION THAT THEY ALSO RANKED MAIN WEAPONS, DUMBASS",
			},
		},
		BIG_CHARGERS: {
			Best:   []Weapon{BURST_CANNON, SCATTERGUN, HEL_RIFLE},
			Good:   []Weapon{CHOKE_MOD, HEL_GUN},
			Decent: []Weapon{HIGH_CAL, SHOTGUN},
			Niche:  []Weapon{COMBAT_SHOTGUN, VERUTA, ARBALIST, HEAVY_ASSAULT_RIFLE},
			DoNot:  []Weapon{SHORT_RIFLE, REVOLVER, PRECISION_RIFLE, SNIPER},
			Notes:  []string{},
		},
		CHARGERS: {
			Best:   []Weapon{HEL_RIFLE, HEL_GUN},
			Good:   []Weapon{CHOKE_MOD, HIGH_CAL},
			Decent: []Weapon{SHOTGUN, COMBAT_SHOTGUN, BURST_CANNON, SCATTERGUN, VERUTA},
			Niche:  []Weapon{SHORT_RIFLE, ARBALIST, HEAVY_ASSAULT_RIFLE},
			DoNot:  []Weapon{PRECISION_RIFLE, REVOLVER, SNIPER},
			Notes:  []string{},
		},
		CHARGER_SCOUTS: {
			Best:   []Weapon{BURST_CANNON, SCATTERGUN, CHOKE_MOD},
			Good:   []Weapon{SNIPER, HEL_RIFLE},
			Decent: []Weapon{HIGH_CAL, SHOTGUN},
			Niche:  []Weapon{HEL_GUN, REVOLVER, COMBAT_SHOTGUN},
			DoNot:  []Weapon{SHORT_RIFLE, ARBALIST, HEAVY_ASSAULT_RIFLE, VERUTA, PRECISION_RIFLE},
			Notes:  []string{},
		},
		CQC: {
			Best:   []Weapon{COMBAT_SHOTGUN, VERUTA, SHOTGUN},
			Good:   []Weapon{HIGH_CAL, HEAVY_ASSAULT_RIFLE, HEL_GUN, ARBALIST, REVOLVER},
			Decent: []Weapon{SCATTERGUN, SHORT_RIFLE, HEL_RIFLE, CHOKE_MOD, BURST_CANNON},
			Niche:  []Weapon{},
			DoNot:  []Weapon{SNIPER, PRECISION_RIFLE},
			Notes: []string{
				"Sniper and Precision Rifle were not officially assigned a tier, but they dissed them both",
			},
		},
		FLYERS: {
			Best:   []Weapon{HEL_GUN},
			Good:   []Weapon{HEL_RIFLE, CHOKE_MOD, HIGH_CAL},
			Decent: []Weapon{PRECISION_RIFLE, REVOLVER, SHOTGUN},
			Niche:  []Weapon{HEAVY_ASSAULT_RIFLE, BURST_CANNON, SHORT_RIFLE},
			DoNot:  []Weapon{COMBAT_SHOTGUN, SCATTERGUN, ARBALIST, VERUTA, SNIPER},
			Notes:  []string{},
		},
		GENERIC: {
			Best:   []Weapon{HEL_GUN},
			Good:   []Weapon{HIGH_CAL, REVOLVER, SHOTGUN, COMBAT_SHOTGUN, HEAVY_ASSAULT_RIFLE},
			Decent: []Weapon{ARBALIST, VERUTA, PRECISION_RIFLE, CHOKE_MOD, HEL_RIFLE, SCATTERGUN},
			Niche:  []Weapon{SHORT_RIFLE, BURST_CANNON},
			DoNot:  []Weapon{SNIPER},
			Notes: []string{
				"A \"typical\" wave meaning some strikers and shooters + a giant",
				"Special mention to the HEL Revolver in \"Best\"",
			},
		},
		GIANTS: {
			Best:   []Weapon{SCATTERGUN, HEL_RIFLE},
			Good:   []Weapon{BURST_CANNON, SNIPER},
			Decent: []Weapon{HEL_GUN, HIGH_CAL, CHOKE_MOD, SHOTGUN, COMBAT_SHOTGUN},
			Niche:  []Weapon{SHORT_RIFLE, VERUTA, ARBALIST, HEAVY_ASSAULT_RIFLE},
			DoNot:  []Weapon{PRECISION_RIFLE, REVOLVER},
			Notes:  []string{},
		},
		KRAKEN: {
			Best:   []Weapon{SNIPER},
			Good:   []Weapon{PRECISION_RIFLE, CHOKE_MOD, BURST_CANNON, HEL_RIFLE},
			Decent: []Weapon{HEL_GUN, REVOLVER, HIGH_CAL},
			Niche:  []Weapon{},
			DoNot:  []Weapon{COMBAT_SHOTGUN, SHOTGUN, SCATTERGUN, SHORT_RIFLE, ARBALIST, VERUTA, HEAVY_ASSAULT_RIFLE},
			Notes:  []string{},
		},
		MOTHERS: {
			Best:   []Weapon{SCATTERGUN},
			Good:   []Weapon{SHOTGUN, BURST_CANNON, CHOKE_MOD, PRECISION_RIFLE},
			Decent: []Weapon{HIGH_CAL, COMBAT_SHOTGUN},
			Niche:  []Weapon{HEL_GUN, HEL_RIFLE},
			DoNot:  []Weapon{SHORT_RIFLE, HEAVY_ASSAULT_RIFLE, ARBALIST, REVOLVER, VERUTA, SNIPER},
			Notes:  []string{},
		},
		NIGHTMARES: {
			Best:   []Weapon{HEL_RIFLE},
			Good:   []Weapon{HEL_GUN, CHOKE_MOD, HIGH_CAL},
			Decent: []Weapon{SHOTGUN, PRECISION_RIFLE, SCATTERGUN, BURST_CANNON, COMBAT_SHOTGUN, VERUTA},
			Niche:  []Weapon{HEAVY_ASSAULT_RIFLE},
			DoNot:  []Weapon{SHORT_RIFLE, ARBALIST, REVOLVER, SNIPER},
			Notes: []string{
				"CHRIS: MENTION THAT THEY ALSO RANKED MAIN WEAPONS, DUMBASS",
			},
		},
		P_MOTHER: {
			Best:   []Weapon{SCATTERGUN},
			Good:   []Weapon{BURST_CANNON, SNIPER, SHOTGUN, PRECISION_RIFLE},
			Decent: []Weapon{CHOKE_MOD, HIGH_CAL, HEL_RIFLE},
			Niche:  []Weapon{HEL_GUN, COMBAT_SHOTGUN},
			DoNot:  []Weapon{REVOLVER, ARBALIST, VERUTA, HEAVY_ASSAULT_RIFLE, SHORT_RIFLE},
			Notes:  []string{},
		},
		REACTOR: {
			Best:   []Weapon{HEL_RIFLE, HEL_GUN},
			Good:   []Weapon{PRECISION_RIFLE, BURST_CANNON},
			Decent: []Weapon{REVOLVER, HIGH_CAL, SNIPER, CHOKE_MOD, HEAVY_ASSAULT_RIFLE, ARBALIST, VERUTA},
			Niche:  []Weapon{COMBAT_SHOTGUN, SHOTGUN, SHORT_RIFLE, SCATTERGUN},
			DoNot:  []Weapon{},
			Notes:  []string{},
		},
		SCOUTS: {
			Best:   []Weapon{SNIPER, HEL_RIFLE, PRECISION_RIFLE},
			Good:   []Weapon{CHOKE_MOD, BURST_CANNON},
			Decent: []Weapon{HIGH_CAL, SHOTGUN, SCATTERGUN},
			Niche:  []Weapon{},
			DoNot:  []Weapon{SHORT_RIFLE, ARBALIST, REVOLVER, COMBAT_SHOTGUN, HEAVY_ASSAULT_RIFLE, HEL_GUN, VERUTA},
			Notes:  []string{},
		},
		SNATCHERS: {
			Best:   []Weapon{SCATTERGUN},
			Good:   []Weapon{CHOKE_MOD, BURST_CANNON},
			Decent: []Weapon{SHOTGUN, HIGH_CAL},
			Niche:  []Weapon{SNIPER, HEL_RIFLE, HEL_GUN, COMBAT_SHOTGUN},
			DoNot:  []Weapon{SHORT_RIFLE, HEAVY_ASSAULT_RIFLE, ARBALIST, REVOLVER, VERUTA, PRECISION_RIFLE},
			Notes:  []string{},
		},
		TANK: {
			Best:   []Weapon{SCATTERGUN},
			Good:   []Weapon{CHOKE_MOD, PRECISION_RIFLE, BURST_CANNON, SHOTGUN},
			Decent: []Weapon{SNIPER, HIGH_CAL, HEL_RIFLE},
			Niche:  []Weapon{HEL_GUN, COMBAT_SHOTGUN},
			DoNot:  []Weapon{REVOLVER, ARBALIST, VERUTA, HEAVY_ASSAULT_RIFLE, SHORT_RIFLE},
			Notes:  []string{},
		},
	}
}
