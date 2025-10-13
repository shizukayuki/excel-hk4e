package excel

var (
	WeaponCodexExcelConfigData   []*WeaponCodex
	WeaponCurveExcelConfigData   []*Curve
	WeaponExcelConfigData        []*Weapon
	WeaponPromoteExcelConfigData []*WeaponPromote
)

func init() {
	load("ExcelBinOutput/WeaponCodexExcelConfigData.json", &WeaponCodexExcelConfigData)
	load("ExcelBinOutput/WeaponCurveExcelConfigData.json", &WeaponCurveExcelConfigData)
	load("ExcelBinOutput/WeaponExcelConfigData.json", &WeaponExcelConfigData)
	load("ExcelBinOutput/WeaponPromoteExcelConfigData.json", &WeaponPromoteExcelConfigData)
}

type WeaponCodex struct {
	WeaponId         uint32
	IsDisuse         bool
	ShowOnlyUnlocked bool
}

func (w *WeaponCodex) Weapon() *Weapon {
	return FindWeapon(w.WeaponId)
}

type WeaponProperty struct {
	PropType  FightProp
	InitValue float64
	Type      GrowCurveType
}

type Weapon struct {
	Item
	WeaponType                 WeaponType
	RankLevel                  uint32
	SkillAffix                 []uint32
	AwakenMaterial             uint32
	WeaponProp                 []*WeaponProperty
	AwakenIcon                 string
	WeaponPromoteId            uint32
	StoryId                    uint32
	AwakenCosts                []uint32
	EnhanceRule                uint32
	DestroyRule                string // MaterialDestroyType
	DestroyReturnMaterial      []uint32
	DestroyReturnMaterialCount []uint32
	InitialLockState           uint32
}

func (w *Weapon) Codex() *WeaponCodex {
	return Find(WeaponCodexExcelConfigData, func(v *WeaponCodex) bool {
		return v.WeaponId == w.Id
	})
}

func (w *Weapon) Affix(level uint32) *EquipAffix {
	if len(w.SkillAffix) == 0 {
		return nil
	}
	return FindEquipAffix(w.SkillAffix[0], level)
}

func (w *Weapon) Curve(level uint32) *Curve {
	return FindCurveData(WeaponCurveExcelConfigData, level)
}

func (w *Weapon) Promote(level uint32) *WeaponPromote {
	return Find(WeaponPromoteExcelConfigData, func(v *WeaponPromote) bool {
		return v.WeaponPromoteId == w.WeaponPromoteId && v.PromoteLevel == level
	})
}

type WeaponPromote struct {
	WeaponPromoteId     uint32
	PromoteLevel        uint32
	CostItems           []*IdCount
	CoinCost            uint32
	AddProps            []*PropValue
	UnlockMaxLevel      uint32
	RequiredPlayerLevel uint32
}

func FindWeapon(id uint32) *Weapon {
	return Find(WeaponExcelConfigData, func(v *Weapon) bool {
		return v.Id == id
	})
}
