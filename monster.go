package excel

var (
	MonsterCurveExcelConfigData    []*CurveData
	MonsterDescribeExcelConfigData []*MonsterDescribe
	MonsterExcelConfigData         []*Monster
)

func init() {
	load("ExcelBinOutput/MonsterCurveExcelConfigData.json", &MonsterCurveExcelConfigData)
	load("ExcelBinOutput/MonsterDescribeExcelConfigData.json", &MonsterDescribeExcelConfigData)
	load("ExcelBinOutput/MonsterExcelConfigData.json", &MonsterExcelConfigData)
}

type MonsterDescribe struct {
	Id               uint32
	NameTextMapHash  TextMapHash
	TitleId          uint32
	SpecialNameLabId uint32
	Icon             string
}

func (m *MonsterDescribe) Name() string {
	return m.NameTextMapHash.String()
}

type Monster struct {
	MonsterName string
	Type        string
	HpDrops     []struct {
		DropId    uint32
		HpPercent float32
	}
	KillDropId        uint32
	FeatureTagGroupId uint32
	DescribeId        uint32
	HpBase            float32
	AttackBase        float32
	DefenseBase       float32
	FireSubHurt       float32
	GrassSubHurt      float32
	WaterSubHurt      float32
	ElecSubHurt       float32
	WindSubHurt       float32
	IceSubHurt        float32
	RockSubHurt       float32
	PhysicalSubHurt   float32
	PropGrowCurves    []PropGrowCurves
	Id                uint32
	CampId            uint32
}

type HpDrop struct {
	DropId    uint32
	HpPercent float32
}

func (m *Monster) Describe() *MonsterDescribe {
	return FindMonsterDescribe(m.DescribeId)
}

func (m *Monster) Curve(level uint32) *CurveData {
	return FindCurveData(MonsterCurveExcelConfigData, level)
}

func FindMonster(id uint32) *Monster {
	for _, v := range MonsterExcelConfigData {
		if v.Id == id {
			return v
		}
	}
	return nil
}

func FindMonsterDescribe(id uint32) *MonsterDescribe {
	for _, v := range MonsterDescribeExcelConfigData {
		if v.Id == id {
			return v
		}
	}
	return nil
}
