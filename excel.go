package excel

import (
	"fmt"
	"strings"
)

var (
	Languages = []string{"CHS", "CHT", "DE", "EN", "ES", "FR", "ID", "IT", "JP", "KR", "PT", "RU", "TH", "TR", "VI"}
	textMap   = map[string]*map[TextMapHash]string{}
)

func init() {
	for _, lang := range Languages {
		var v map[TextMapHash]string
		textMap[strings.ToLower(lang)] = &v
		load(fmt.Sprintf("TextMap/TextMap%s.json", lang), &v)
	}
}

type TextMapHash uint32

func (h TextMapHash) Lang(lang string) string {
	textMap := textMap[strings.ToLower(lang)]
	if textMap != nil && *textMap != nil {
		return (*textMap)[h]
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
	for _, v := range c.CurveInfos {
		if v.Type == typ {
			return v.Arith, v.Value
		}
	}
	return "", 0
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
	for _, v := range list {
		if v.Level != level {
			continue
		}
		return v
	}
	return nil
}
