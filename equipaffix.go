package excel

var EquipAffixExcelConfigData []*EquipAffix

func init() {
	load("ExcelBinOutput/EquipAffixExcelConfigData.json", &EquipAffixExcelConfigData)
}

type EquipAffix struct {
	AffixId         uint32
	Id              uint32
	Level           uint32
	OpenConfig      string
	NameTextMapHash TextMapHash
	AddProps        []FightPropData
	ParamList       []float32
}

func (e *EquipAffix) Name() string {
	return e.NameTextMapHash.String()
}

func FindEquipAffix(id, level uint32) *EquipAffix {
	return Find(EquipAffixExcelConfigData, func(v *EquipAffix) bool {
		return v.Id == id && v.Level == level
	})
}
