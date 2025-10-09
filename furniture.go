package excel

var HomeWorldFurnitureExcelConfigData []*HomeWorldFurniture

func init() {
	load("ExcelBinOutput/HomeWorldFurnitureExcelConfigData.json", &HomeWorldFurnitureExcelConfigData)
}

type HomeWorldFurniture struct {
	Id              uint32
	NameTextMapHash TextMapHash
}

func (f *HomeWorldFurniture) Name() string {
	return f.NameTextMapHash.String()
}

func FindHomeWorldFurniture(id uint32) *HomeWorldFurniture {
	return Find(HomeWorldFurnitureExcelConfigData, func(v *HomeWorldFurniture) bool {
		return v.Id == id
	})
}
