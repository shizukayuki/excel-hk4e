package excel

var (
	MonsterCurveExcelConfigData    []*Curve
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

type MonsterDrop struct {
	DropId    uint32
	HpPercent float64
}

type Monster struct {
	Id                uint32
	CampId            uint32
	HpBase            float64
	AttackBase        float64
	DefenseBase       float64
	FireSubHurt       float64
	GrassSubHurt      float64
	WaterSubHurt      float64
	ElecSubHurt       float64
	WindSubHurt       float64
	IceSubHurt        float64
	RockSubHurt       float64
	PropGrowCurves    []*FightPropGrow
	PhysicalSubHurt   float64
	MonsterName       string
	Type              string // MonsterType
	Affix             []uint32
	Equips            []uint32
	HpDrops           []*MonsterDrop
	KillDropId        uint32
	FeatureTagGroupId uint32
	DescribeId        uint32
}

func (m *Monster) Describe() *MonsterDescribe {
	return FindMonsterDescribe(m.DescribeId)
}

func (m *Monster) Curve(level uint32) *Curve {
	return FindCurveData(MonsterCurveExcelConfigData, level)
}

func FindMonster(id uint32) *Monster {
	return Find(MonsterExcelConfigData, func(v *Monster) bool {
		return v.Id == id
	})
}

func FindMonsterDescribe(id uint32) *MonsterDescribe {
	return Find(MonsterDescribeExcelConfigData, func(v *MonsterDescribe) bool {
		return v.Id == id
	})
}
