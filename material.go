package excel

var MaterialExcelConfigData []*Material

func init() {
	load("ExcelBinOutput/MaterialExcelConfigData.json", &MaterialExcelConfigData)
}

type Material struct {
	Id              uint32
	NameTextMapHash TextMapHash
}

func (m *Material) Name() string {
	return m.NameTextMapHash.String()
}

func FindMaterial(id uint32) *Material {
	for _, v := range MaterialExcelConfigData {
		if v.Id == id {
			return v
		}
	}
	return nil
}
