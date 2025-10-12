package excel

var HomeWorldFurnitureExcelConfigData []*HomeWorldFurniture

func init() {
	load("ExcelBinOutput/HomeWorldFurnitureExcelConfigData.json", &HomeWorldFurnitureExcelConfigData)
}

type HomeWorldFurniture struct {
	Item
	SurfaceType string // FurnitureDeploySurfaceType
	RankLevel   uint32
}

func FindHomeWorldFurniture(id uint32) *HomeWorldFurniture {
	return Find(HomeWorldFurnitureExcelConfigData, func(v *HomeWorldFurniture) bool {
		return v.Id == id
	})
}
