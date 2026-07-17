package excel

//go:generate enumer --text --json --type=ItemType $GOFILE
type ItemType uint32

const (
	ITEM_NONE      ItemType = 0
	ITEM_VIRTUAL   ItemType = 1
	ITEM_MATERIAL  ItemType = 2
	ITEM_RELIQUARY ItemType = 3
	ITEM_WEAPON    ItemType = 4
	ITEM_DISPLAY   ItemType = 5
	ITEM_FURNITURE ItemType = 6
)
