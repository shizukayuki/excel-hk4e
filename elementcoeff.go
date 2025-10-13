package excel

var ElementCoeffExcelConfigData []*ElementCoeff

func init() {
	load("ExcelBinOutput/ElementCoeffExcelConfigData.json", &ElementCoeffExcelConfigData)
}

type ElementCoeff struct {
	Level                uint32
	CrashCo              float64
	ElementLevelCo       float64
	PlayerElementLevelCo float64
	PlayerShieldLevelCo  float64
}

func FindElementCoeff(level uint32) *ElementCoeff {
	return Find(ElementCoeffExcelConfigData, func(v *ElementCoeff) bool {
		return v.Level == level
	})
}
