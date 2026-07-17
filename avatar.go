package excel

var (
	AvatarCodexExcelConfigData      []*AvatarCodex
	AvatarCurveExcelConfigData      []*Curve
	AvatarExcelConfigData           []*Avatar
	AvatarPromoteExcelConfigData    []*AvatarPromote
	AvatarSkillDepotExcelConfigData []*AvatarSkillDepot
	AvatarSkillExcelConfigData      []*AvatarSkill
	AvatarTalentExcelConfigData     []*AvatarTalent
	ProudSkillExcelConfigData       []*ProudSkill
)

func init() {
	load("ExcelBinOutput/AvatarCodexExcelConfigData.json", &AvatarCodexExcelConfigData)
	load("ExcelBinOutput/AvatarCurveExcelConfigData.json", &AvatarCurveExcelConfigData)
	load("ExcelBinOutput/AvatarExcelConfigData.json", &AvatarExcelConfigData)
	load("ExcelBinOutput/AvatarPromoteExcelConfigData.json", &AvatarPromoteExcelConfigData)
	load("ExcelBinOutput/AvatarSkillDepotExcelConfigData.json", &AvatarSkillDepotExcelConfigData)
	load("ExcelBinOutput/AvatarSkillExcelConfigData.json", &AvatarSkillExcelConfigData)
	load("ExcelBinOutput/AvatarTalentExcelConfigData.json", &AvatarTalentExcelConfigData)
	load("ExcelBinOutput/ProudSkillExcelConfigData.json", &ProudSkillExcelConfigData)
}

type AvatarCodex struct {
	AvatarId uint32
}

func (a *AvatarCodex) Avatar() *Avatar {
	return FindAvatar(a.AvatarId)
}

type Avatar struct {
	Creature
	UseType             string // AvatarUseType
	BodyType            string // BodyType
	IconName            string
	SideIconName        string
	QualityType         string // QualityType
	ChargeEfficiency    float64
	IsRangeAttack       bool
	WeaponType          WeaponType
	ImageName           string
	SkillDepotId        uint32
	StaminaRecoverSpeed float64
	CandSkillDepotIds   []uint32
	DescTextMapHash     TextMapHash
	AvatarIdentityType  string // AvatarIdentityType
	AvatarPromoteId     uint32
	FeatureTagGroupId   uint32
	InfoDescTextMapHash TextMapHash
}

func (a *Avatar) Codex() *AvatarCodex {
	return Find(AvatarCodexExcelConfigData, func(v *AvatarCodex) bool {
		return v.AvatarId == a.Id
	})
}

func (a *Avatar) SkillDepot() *AvatarSkillDepot {
	return FindSkillDepot(a.SkillDepotId)
}

func (a *Avatar) Curve(level uint32) *Curve {
	return FindCurveData(AvatarCurveExcelConfigData, level)
}

func (a *Avatar) Promote(level uint32) *AvatarPromote {
	return Find(AvatarPromoteExcelConfigData, func(v *AvatarPromote) bool {
		return v.AvatarPromoteId == a.AvatarPromoteId && v.PromoteLevel == level
	})
}

func (a *Avatar) FetterInfo() *FetterInfo {
	return Find(FetterInfoExcelConfigData, func(v *FetterInfo) bool {
		return v.AvatarId == a.Id
	})
}

type AvatarPromote struct {
	AvatarPromoteId     uint32
	PromoteLevel        uint32
	ScoinCost           uint32
	CostItems           []*IdCount
	UnlockMaxLevel      uint32
	AddProps            []*PropValue
	RequiredPlayerLevel uint32
}

type ProudSkillOpen struct {
	ProudSkillGroupId      uint32
	NeedAvatarPromoteLevel uint32
}

type AvatarSkillDepot struct {
	Id                      uint32
	EnergySkill             uint32
	Skills                  []uint32
	SubSkills               []uint32
	AttackModeSkill         uint32
	Talents                 []uint32
	TalentStarName          string
	InherentProudSkillOpens []*ProudSkillOpen
	SkillDepotAbilityGroup  string
	ArkheType               string `json:"__exp_arkheType"` // custom field
}

type AvatarSkill struct {
	Id                 uint32
	NameTextMapHash    TextMapHash
	AbilityName        string
	DescTextMapHash    TextMapHash
	SkillIcon          string
	IsRanged           bool
	CDTime             float64
	IgnoreCDMinusRatio bool
	CostStamina        float64
	CostElemType       ElementType
	CostElemVal        float64
	MaxChargeNum       int
	TriggerID          int
	ProudSkillGroupId  uint32
	CDSlot             uint32
	SpecialEnergyMax   uint32
	SpecialEnergyMin   uint32
	SpecialEnergyType  string
}

func (a *AvatarSkill) Name() string {
	return a.NameTextMapHash.String()
}

func (a *AvatarSkill) ProudSkill(level uint32) *ProudSkill {
	return FindProudSkill(a.ProudSkillGroupId, level)
}

type AvatarTalent struct {
	BaseTalent
	TalentId          uint32
	NameTextMapHash   TextMapHash
	DescTextMapHash   TextMapHash
	Icon              string
	PrevTalent        uint32
	MainCostItemId    uint32
	MainCostItemCount uint32
}

func (a *AvatarTalent) Name() string {
	return a.NameTextMapHash.String()
}

func (a *AvatarTalent) Prev() *AvatarTalent {
	return FindTalent(a.PrevTalent)
}

type ProudSkill struct {
	BaseTalent
	ProudSkillId      uint32
	ProudSkillGroupId uint32
	Level             uint32
	ProudSkillType    uint32
	NameTextMapHash   TextMapHash
	DescTextMapHash   TextMapHash
	Icon              string
	CoinCost          uint32
	CostItems         []*IdCount
	Breaklevel        uint32
	ParamDescList     []TextMapHash
}

func (a *ProudSkill) Name() string {
	return a.NameTextMapHash.String()
}

func (a *ProudSkill) FindSkill() *AvatarSkill {
	return Find(AvatarSkillExcelConfigData, func(v *AvatarSkill) bool {
		return v.ProudSkillGroupId == a.ProudSkillGroupId
	})
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

func FindProudSkill(id, level uint32) *ProudSkill {
	return Find(ProudSkillExcelConfigData, func(v *ProudSkill) bool {
		return v.ProudSkillGroupId == id && v.Level == level
	})
}

func FindTalent(id uint32) *AvatarTalent {
	return Find(AvatarTalentExcelConfigData, func(v *AvatarTalent) bool {
		return v.TalentId == id
	})
}
