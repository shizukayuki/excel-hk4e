package excel

var FeatureTagGroupExcelConfigData []*FeatureTagGroup

func init() {
	load("ExcelBinOutput/FeatureTagGroupExcelConfigData.json", &FeatureTagGroupExcelConfigData)
}

type FeatureTagGroup struct {
	GroupID uint32
	TagIDs  []uint32
}

func FindFeatureTag(id uint32) *FeatureTagGroup {
	return Find(FeatureTagGroupExcelConfigData, func(v *FeatureTagGroup) bool {
		return v.GroupID == id
	})
}
