package excel

//go:generate enumer --json --type=ElementType -trimprefix=Element $GOFILE
type ElementType uint32

const (
	ElementNone           ElementType = 0
	ElementFire           ElementType = 1
	ElementWater          ElementType = 2
	ElementGrass          ElementType = 3
	ElementElectric       ElementType = 4
	ElementIce            ElementType = 5
	ElementFrozen         ElementType = 6
	ElementWind           ElementType = 7
	ElementRock           ElementType = 8
	ElementAntiFire       ElementType = 9
	ElementVehicleMuteIce ElementType = 10
	ElementMushroom       ElementType = 11
	ElementOverdose       ElementType = 12
	ElementWood           ElementType = 13
)
