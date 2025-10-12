package excel

var EquipAffixExcelConfigData []*EquipAffix

func init() {
	load("ExcelBinOutput/EquipAffixExcelConfigData.json", &EquipAffixExcelConfigData)
}

type EquipAffix struct {
	BaseTalent
	AffixId         uint32
	Id              uint32
	Level           uint32
	NameTextMapHash TextMapHash
	DescTextMapHash TextMapHash
}

func (e *EquipAffix) Name() string {
	return e.NameTextMapHash.String()
}

func FindEquipAffix(id, level uint32) *EquipAffix {
	return Find(EquipAffixExcelConfigData, func(v *EquipAffix) bool {
		return v.Id == id && v.Level == level
	})
}
