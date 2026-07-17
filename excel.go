package excel

import (
	"fmt"
	"strings"
)

var (
	Languages = []string{"CHS", "CHT", "DE", "EN", "ES", "FR", "ID", "IT", "JP", "KR", "PT", "RU", "TH", "TR", "VI"}
	textMap   = make(map[string]map[TextMapHash]string)

	ManualTextMapConfigData []*ManualTextMap
)

func init() {
	for _, lang := range Languages {
		v := make(map[TextMapHash]string)
		textMap[strings.ToLower(lang)] = v
		load(fmt.Sprintf("TextMap/TextMap%s.json", lang), v)
	}
	load("ExcelBinOutput/ManualTextMapConfigData.json", &ManualTextMapConfigData)
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

type ManualTextMap struct {
	TextMapId                 string
	TextMapContentTextMapHash TextMapHash
	ParamTypes                []string
}

func (m *ManualTextMap) Name() string {
	return m.TextMapContentTextMapHash.String()
}

func (m *ManualTextMap) Lang(lang string) string {
	return m.TextMapContentTextMapHash.Lang(lang)
}

func FindManualTextMap(id string) *ManualTextMap {
	return Find(ManualTextMapConfigData, func(v *ManualTextMap) bool {
		return v.TextMapId == id
	})
}
