package excel

//go:generate enumer --text --json --type=WeaponType $GOFILE
type WeaponType uint32

const (
	WEAPON_NONE           WeaponType = 0
	WEAPON_SWORD_ONE_HAND WeaponType = 1
	WEAPON_CROSSBOW       WeaponType = 2
	WEAPON_STAFF          WeaponType = 3
	WEAPON_DOUBLE_DAGGER  WeaponType = 4
	WEAPON_KATANA         WeaponType = 5
	WEAPON_SHURIKEN       WeaponType = 6
	WEAPON_STICK          WeaponType = 7
	WEAPON_SPEAR          WeaponType = 8
	WEAPON_SHIELD_SMALL   WeaponType = 9
	WEAPON_CATALYST       WeaponType = 10
	WEAPON_CLAYMORE       WeaponType = 11
	WEAPON_BOW            WeaponType = 12
	WEAPON_POLE           WeaponType = 13
)
