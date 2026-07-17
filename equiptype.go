package excel

//go:generate enumer --text --json --type=EquipType $GOFILE
type EquipType uint32

const (
	EQUIP_NONE     EquipType = 0
	EQUIP_BRACER   EquipType = 1
	EQUIP_NECKLACE EquipType = 2
	EQUIP_SHOES    EquipType = 3
	EQUIP_RING     EquipType = 4
	EQUIP_DRESS    EquipType = 5
	EQUIP_WEAPON   EquipType = 6
)
