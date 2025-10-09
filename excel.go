package excel

import (
	"fmt"
	"strings"
)

var (
	Languages = []string{"CHS", "CHT", "DE", "EN", "ES", "FR", "ID", "IT", "JP", "KR", "PT", "RU", "TH", "TR", "VI"}
	textMap   = make(map[string]map[TextMapHash]string)
)

func init() {
	for _, lang := range Languages {
		v := make(map[TextMapHash]string)
		textMap[strings.ToLower(lang)] = v
		load(fmt.Sprintf("TextMap/TextMap%s.json", lang), v)
	}
}

type TextMapHash uint32

func (h TextMapHash) Lang(lang string) string {
	textMap := textMap[strings.ToLower(lang)]
	if textMap != nil {
		return textMap[h]
	}
	return ""
}

func (h TextMapHash) String() string {
	return h.Lang("en")
}

type CurveInfo struct {
	Type  string
	Arith string
	Value float32
}

type CurveData struct {
	Level      uint32
	CurveInfos []*CurveInfo
}

func (c *CurveData) Type(typ string) (string, float32) {
	ci := Find(c.CurveInfos, func(v *CurveInfo) bool {
		return v.Type == typ
	})
	if ci == nil {
		return "", 0
	}
	return ci.Arith, ci.Value
}

type PropGrowCurves struct {
	Type      FightProp
	GrowCurve string
}

type FightPropData struct {
	PropType FightProp
	Value    float32
}

func FindCurveData(list []*CurveData, level uint32) *CurveData {
	return Find(list, func(v *CurveData) bool {
		return v.Level == level
	})
}
