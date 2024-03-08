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
	for _, v := range HomeWorldFurnitureExcelConfigData {
		if v.Id == id {
			return v
		}
	}
	return nil
}
