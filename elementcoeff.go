package excel

var ElementCoeffExcelConfigData []*ElementCoeff

func init() {
	load("ExcelBinOutput/ElementCoeffExcelConfigData.json", &ElementCoeffExcelConfigData)
}

type ElementCoeff struct {
	Level                uint32
	CrashCo              float32
	ElementLevelCo       float32
	PlayerElementLevelCo float32
	PlayerShieldLevelCo  float32
}

func FindElementCoeff(level uint32) *ElementCoeff {
	for _, v := range ElementCoeffExcelConfigData {
		if v.Level == level {
			return v
		}
	}
	return nil
}
