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
	return Find(MaterialExcelConfigData, func(v *Material) bool {
		return v.Id == id
	})
}
