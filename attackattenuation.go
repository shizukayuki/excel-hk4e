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
	return Find(AttackAttenuationExcelConfigData, func(v *AttackAttenuation) bool {
		return v.Group == group
	})
}
