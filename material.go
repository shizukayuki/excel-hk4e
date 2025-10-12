package excel

var MaterialExcelConfigData []*Material

func init() {
	load("ExcelBinOutput/MaterialExcelConfigData.json", &MaterialExcelConfigData)
}

type Material struct {
	Item
	MaterialType string // MaterialType
	StackLimit   uint32
	MaxUseCount  uint32
	RankLevel    uint32
	FoodQuality  string // FoodQualityType
}

func FindMaterial(id uint32) *Material {
	return Find(MaterialExcelConfigData, func(v *Material) bool {
		return v.Id == id
	})
}
