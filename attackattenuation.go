package excel

var AttackAttenuationExcelConfigData []*AttackAttenuation

func init() {
	load("ExcelBinOutput/AttackAttenuationExcelConfigData.json", &AttackAttenuationExcelConfigData)
}

type AttackAttenuation struct {
	Group              string
	ResetCycle         float64
	DurabilitySequence []float64
	EnbreakSequence    []float64
	DamageSequence     []float64
}

func FindAttackAttenuation(group string) *AttackAttenuation {
	return Find(AttackAttenuationExcelConfigData, func(v *AttackAttenuation) bool {
		return v.Group == group
	})
}
