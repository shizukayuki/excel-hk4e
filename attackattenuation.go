package excel

var AttackAttenuationExcelConfigData []*AttackAttenuation

func init() {
	load("ExcelBinOutput/AttackAttenuationExcelConfigData.json", &AttackAttenuationExcelConfigData)
}

type AttackAttenuation struct {
	Group              string
	ResetCycle         float32
	DurabilitySequence []float32
	EnbreakSequence    []float32
	DamageSequence     []float32
}

func FindAttackAttenuation(group string) *AttackAttenuation {
	for _, v := range AttackAttenuationExcelConfigData {
		if v.Group == group {
			return v
		}
	}
	return nil
}
