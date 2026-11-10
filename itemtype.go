package excel

//go:generate enumer --text --json --type=ItemType $GOFILE
type ItemType uint32

const (
	ITEM_NONE                 ItemType = 0
	ITEM_VIRTUAL              ItemType = 1
	ITEM_MATERIAL             ItemType = 2
	ITEM_RELIQUARY            ItemType = 3
	ITEM_WEAPON               ItemType = 4
	ITEM_DISPLAY              ItemType = 5
	ITEM_FURNITURE            ItemType = 6
	ITEM_BEYOND_MATERIAL      ItemType = 7
	ITEM_BEYOND_HALL_FACILITY ItemType = 8
	ITEM_TPS_WEAPON           ItemType = 9
)
