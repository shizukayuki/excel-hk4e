package excel

type IdCount struct {
	Id    uint32
	Count uint32
}

type PropValue struct {
	PropType FightProp
	Value    float64
}

type BaseTalent struct {
	OpenConfig string
	AddProps   []*PropValue
	ParamList  []float64
}

type Item struct {
	Id              uint32
	NameTextMapHash TextMapHash
	DescTextMapHash TextMapHash
	Icon            string
	ItemType        ItemType
	Weight          uint32
	Rank            uint32
	GadgetId        uint32
	Dropable        bool
	UseLevel        uint32
	GlobalItemLimit uint32
}

func (i *Item) Name() string {
	return i.NameTextMapHash.String()
}

type GrowCurveInfo struct {
	Type  GrowCurveType
	Arith string // ArithType
	Value float64
}

type Curve struct {
	Level      uint32
	CurveInfos []*GrowCurveInfo
}

func (c *Curve) Type(typ GrowCurveType) (string, float64) {
	ci := Find(c.CurveInfos, func(v *GrowCurveInfo) bool {
		return v.Type == typ
	})
	if ci == nil {
		return "", 0
	}
	return ci.Arith, ci.Value
}

type FightPropGrow struct {
	Type      FightProp
	GrowCurve GrowCurveType
}

func FindCurveData(list []*Curve, level uint32) *Curve {
	return Find(list, func(v *Curve) bool {
		return v.Level == level
	})
}
