package excel

import (
	"cmp"
	"slices"
)

var (
	ReliquaryAffixExcelConfigData    []*ReliquaryAffix
	ReliquaryCodexExcelConfigData    []*ReliquaryCodex
	ReliquaryExcelConfigData         []*Reliquary
	ReliquaryLevelExcelConfigData    []*ReliquaryLevel
	ReliquaryMainPropExcelConfigData []*ReliquaryMainProp
	ReliquarySetExcelConfigData      []*ReliquarySet
)

func init() {
	load("ExcelBinOutput/ReliquaryAffixExcelConfigData.json", &ReliquaryAffixExcelConfigData)
	load("ExcelBinOutput/ReliquaryCodexExcelConfigData.json", &ReliquaryCodexExcelConfigData)
	load("ExcelBinOutput/ReliquaryExcelConfigData.json", &ReliquaryExcelConfigData)
	load("ExcelBinOutput/ReliquaryLevelExcelConfigData.json", &ReliquaryLevelExcelConfigData)
	load("ExcelBinOutput/ReliquaryMainPropExcelConfigData.json", &ReliquaryMainPropExcelConfigData)
	load("ExcelBinOutput/ReliquarySetExcelConfigData.json", &ReliquarySetExcelConfigData)
}

type ReliquaryAffix struct {
	Id        uint32
	DepotId   uint32
	GroupId   uint32
	PropType  FightProp
	PropValue float32
}

func (r *ReliquaryAffix) Min() *ReliquaryAffix {
	return r.Rolls()[0]
}

func (r *ReliquaryAffix) Max() *ReliquaryAffix {
	v := r.Rolls()
	return v[len(v)-1]
}

func (r *ReliquaryAffix) Rolls() []*ReliquaryAffix {
	s := Filter(ReliquaryAffixExcelConfigData, func(v *ReliquaryAffix) bool {
		return v.DepotId == r.DepotId && v.PropType == r.PropType
	})
	slices.SortFunc(s, func(a, b *ReliquaryAffix) int { return cmp.Compare(a.PropValue, b.PropValue) })
	return s
}

type ReliquaryCodex struct {
	Id        uint32
	SuitId    uint32
	Level     uint32
	CupId     uint32
	LeatherId uint32
	CapId     uint32
	FlowerId  uint32
	SandId    uint32
}

func (r *ReliquaryCodex) Reliquary(typ EquipType) *Reliquary {
	return FindReliquary(r.equipType(typ))
}

func (r *ReliquaryCodex) Set() *ReliquarySet {
	return FindReliquarySet(r.SuitId)
}

func (r *ReliquaryCodex) equipType(typ EquipType) uint32 {
	switch typ {
	case EQUIP_BRACER:
		return r.FlowerId
	case EQUIP_NECKLACE:
		return r.LeatherId
	case EQUIP_SHOES:
		return r.SandId
	case EQUIP_RING:
		return r.CupId
	case EQUIP_DRESS:
		return r.CapId
	default:
		return 0
	}
}

type Reliquary struct {
	Item
	EquipType                  EquipType
	RankLevel                  uint32
	MainPropDepotId            uint32
	AppendPropDepotId          uint32
	AppendPropNum              uint32
	SetId                      uint32
	AddPropLevels              []uint32
	BaseConvExp                uint32
	MaxLevel                   uint32
	StoryId                    uint32
	DestroyRule                string // MaterialDestroyType
	DestroyReturnMaterial      []uint32
	DestroyReturnMaterialCount []uint32
	InitialLockState           uint32
}

func (r *Reliquary) Codex() *ReliquaryCodex {
	return Find(ReliquaryCodexExcelConfigData, func(v *ReliquaryCodex) bool {
		return v.equipType(r.EquipType) == r.Id
	})
}

func (r *Reliquary) Set() *ReliquarySet {
	return FindReliquarySet(r.SetId)
}

func (r *Reliquary) Level(level uint32) *ReliquaryLevel {
	return Find(ReliquaryLevelExcelConfigData, func(v *ReliquaryLevel) bool {
		return v.Rank == r.RankLevel && v.Level == level
	})
}

func (r *Reliquary) MainProp(id uint32) *ReliquaryMainProp {
	return Find(ReliquaryMainPropExcelConfigData, func(v *ReliquaryMainProp) bool {
		return v.PropDepotId == r.MainPropDepotId && v.Id == id
	})
}

func (r *Reliquary) AppendProp(id uint32) *ReliquaryAffix {
	return Find(ReliquaryAffixExcelConfigData, func(v *ReliquaryAffix) bool {
		return v.DepotId == r.AppendPropDepotId && v.Id == id
	})
}

type ReliquaryLevel struct {
	Rank     uint32
	Level    uint32
	Exp      uint32
	AddProps []*PropValue
}

func (r *ReliquaryLevel) Stat(prop FightProp) float32 {
	p := Find(r.AddProps, func(v *PropValue) bool {
		return v.PropType == prop
	})
	if p == nil {
		return 0
	}
	return p.Value
}

type ReliquaryMainProp struct {
	Id          uint32
	PropDepotId uint32
	PropType    FightProp
	AffixName   string // ReliquaryMainAffixName
}

type ReliquarySet struct {
	SetId         uint32
	SetIcon       string
	SetNeedNum    []uint32
	EquipAffixId  uint32
	DisableFilter uint32
	ContainsList  []uint32
	BagSortValue  uint32
	DungeonGroup  []uint32
	TextList      []uint32
}

func (r *ReliquarySet) Codex(level uint32) *ReliquaryCodex {
	return Find(ReliquaryCodexExcelConfigData, func(v *ReliquaryCodex) bool {
		return v.SuitId == r.SetId && v.Level == level
	})
}

func (r *ReliquarySet) Affix(level uint32) *EquipAffix {
	return FindEquipAffix(r.EquipAffixId, level)
}

func FindReliquary(id uint32) *Reliquary {
	return Find(ReliquaryExcelConfigData, func(v *Reliquary) bool {
		return v.Id == id
	})
}

func FindReliquarySet(id uint32) *ReliquarySet {
	return Find(ReliquarySetExcelConfigData, func(v *ReliquarySet) bool {
		return v.SetId == id
	})
}
