package excel

var (
	AvatarCodexExcelConfigData      []*AvatarCodex
	AvatarCurveExcelConfigData      []*CurveData
	AvatarExcelConfigData           []*Avatar
	AvatarPromoteExcelConfigData    []*AvatarPromote
	AvatarSkillDepotExcelConfigData []*AvatarSkillDepot
	AvatarSkillExcelConfigData      []*AvatarSkill
	ProudSkillExcelConfigData       []*ProudSkill
)

func init() {
	load("ExcelBinOutput/AvatarCodexExcelConfigData.json", &AvatarCodexExcelConfigData)
	load("ExcelBinOutput/AvatarCurveExcelConfigData.json", &AvatarCurveExcelConfigData)
	load("ExcelBinOutput/AvatarExcelConfigData.json", &AvatarExcelConfigData)
	load("ExcelBinOutput/AvatarPromoteExcelConfigData.json", &AvatarPromoteExcelConfigData)
	load("ExcelBinOutput/AvatarSkillDepotExcelConfigData.json", &AvatarSkillDepotExcelConfigData)
	load("ExcelBinOutput/AvatarSkillExcelConfigData.json", &AvatarSkillExcelConfigData)
	load("ExcelBinOutput/ProudSkillExcelConfigData.json", &ProudSkillExcelConfigData)
}

type AvatarCodex struct {
	AvatarId uint32
}

func (a *AvatarCodex) Avatar() *Avatar {
	return FindAvatar(a.AvatarId)
}

type Avatar struct {
	UseType             string
	BodyType            string
	IconName            string
	SideIconName        string
	QualityType         string
	ChargeEfficiency    float32
	WeaponType          string
	ImageName           string
	SkillDepotId        uint32
	StaminaRecoverSpeed float32
	CandSkillDepotIds   []uint32
	AvatarPromoteId     uint32
	HpBase              float32
	AttackBase          float32
	DefenseBase         float32
	Critical            float32
	CriticalHurt        float32
	PropGrowCurves      []PropGrowCurves
	Id                  uint32
	NameTextMapHash     TextMapHash
}

func (a *Avatar) Name() string {
	return a.NameTextMapHash.String()
}

func (a *Avatar) Codex() *AvatarCodex {
	return Find(AvatarCodexExcelConfigData, func(v *AvatarCodex) bool {
		return v.AvatarId == a.Id
	})
}

func (a *Avatar) SkillDepot() *AvatarSkillDepot {
	return FindSkillDepot(a.SkillDepotId)
}

func (a *Avatar) Curve(level uint32) *CurveData {
	return FindCurveData(AvatarCurveExcelConfigData, level)
}

func (a *Avatar) Promote(level uint32) *AvatarPromote {
	return Find(AvatarPromoteExcelConfigData, func(v *AvatarPromote) bool {
		return v.AvatarPromoteId == a.AvatarPromoteId && v.PromoteLevel == level
	})
}

type AvatarPromote struct {
	AvatarPromoteId uint32
	PromoteLevel    uint32
	UnlockMaxLevel  uint32
	AddProps        []FightPropData
}

type AvatarSkillDepot struct {
	Id                      uint32
	EnergySkill             uint32
	Skills                  []uint32
	SubSkills               []uint32
	AttackModeSkill         uint32
	Talents                 []uint32
	TalentStarName          string
	InherentProudSkillOpens []struct {
		ProudSkillGroupId      uint32
		NeedAvatarPromoteLevel uint32
	}
}

type AvatarSkill struct {
	Id                 uint32
	NameTextMapHash    TextMapHash
	AbilityName        string
	DescTextMapHash    TextMapHash
	SkillIcon          string
	CDTime             float32
	IgnoreCDMinusRatio bool
	CostElemType       ElementType
	CostElemVal        float32
	MaxChargeNum       int
	ProudSkillGroupId  uint32
	CDSlot             uint32
}

func (a *AvatarSkill) Name() string {
	return a.NameTextMapHash.String()
}

func (a *AvatarSkill) ProudSkill(level uint32) *ProudSkill {
	return Find(ProudSkillExcelConfigData, func(v *ProudSkill) bool {
		return v.ProudSkillGroupId == a.ProudSkillGroupId && v.Level == level
	})
}

type ProudSkill struct {
	ProudSkillGroupId uint32
	Level             uint32
	ProudSkillType    uint32
	Icon              string
	Breaklevel        uint32
	ParamDescList     []TextMapHash
	OpenConfig        string
	ParamList         []float32
}

func FindAvatar(id uint32) *Avatar {
	return Find(AvatarExcelConfigData, func(v *Avatar) bool {
		return v.Id == id
	})
}

func FindSkillDepot(id uint32) *AvatarSkillDepot {
	return Find(AvatarSkillDepotExcelConfigData, func(v *AvatarSkillDepot) bool {
		return v.Id == id
	})
}

func FindSkill(id uint32) *AvatarSkill {
	return Find(AvatarSkillExcelConfigData, func(v *AvatarSkill) bool {
		return v.Id == id
	})
}
