package excel

var ConstValueExcelConfigData []*ConstValue

func init() {
	load("ExcelBinOutput/ConstValueExcelConfigData.json", &ConstValueExcelConfigData)
}

type ConstValue struct {
	Name  string
	Value []string
}

func FindConstValue(name string) *ConstValue {
	return Find(ConstValueExcelConfigData, func(v *ConstValue) bool {
		return v.Name == name
	})
}
