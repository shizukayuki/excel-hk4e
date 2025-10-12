package excel

//go:generate enumer --text --json --type=ItemType $GOFILE
type ItemType uint32

const (
	ITEM_NONE ItemType = iota
	ITEM_VIRTUAL
	ITEM_MATERIAL
	ITEM_RELIQUARY
	ITEM_WEAPON
	ITEM_DISPLAY
	ITEM_FURNITURE
)
