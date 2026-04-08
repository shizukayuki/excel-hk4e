package excel_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/shizukayuki/excel-hk4e"
)

func setup(t *testing.T, langs []string, override map[string]string) {
	err := excel.LoadResources(excel.LoaderConfig{
		Languages: langs,
		ReadFile: func(root string, name string) ([]byte, error) {
			data := []byte("[]")
			if strings.HasPrefix(strings.ToLower(name), "textmap") {
				if b, ok := override[name]; ok {
					fmt.Printf("loaded %s...\n", name)
					data = []byte(b)
				} else {
					fmt.Printf("trying %s...\n", name)
				}
			}
			return data, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTextMapMerge(t *testing.T) {
	override := map[string]string{
		"TextMap/TextMapEN.json": `{
			"1": "1",
			"2": "2",
			"3": "3"
		}`,

		"TextMap/TextMapJP.json":   `{ "1": "1" }`,
		"TextMap/TextMapJP_1.json": `{ "2": "2" }`,
		"TextMap/TextMapJP_2.json": `{ "3": "3" }`,

		"TextMap/TextMapKR_0.json": `{ "1": "1" }`,
		"TextMap/TextMapKR_1.json": `{ "2": "2" }`,

		"TextMap/TextMapTH_1.json": `{ "1": "1" }`,
		"TextMap/TextMapTH_2.json": `{ "2": "2" }`,

		"TextMap/TextMapES_0.json":      `{ "1": "1" }`,
		"TextMap/TextMapES_1.json":      `{ "2": "2" }`,
		"TextMap/TextMap_MediumES.json": `{ "3": "3" }`,

		"TextMap/TextMapFR.json":          `{ "1": "1" }`,
		"TextMap/TextMap_MediumFR_0.json": `{ "2": "2" }`,
		"TextMap/TextMap_MediumFR_1.json": `{ "3": "3" }`,
	}
	expected := map[string]map[excel.TextMapHash]string{
		"en": {
			1: "1",
			2: "2",
			3: "3",
		},
		"jp": {
			1: "1",
			2: "",
			3: "",
		},
		"kr": {
			1: "1",
			2: "2",
			3: "",
		},
		"th": {
			1: "1",
			2: "2",
			3: "",
		},
		"es": {
			1: "1",
			2: "2",
			3: "3",
		},
		"fr": {
			1: "1",
			2: "2",
			3: "3",
		},
	}
	var langs []string
	for k := range expected {
		langs = append(langs, k)
	}
	setup(t, langs, override)

	for lang, expected := range expected {
		for h, v := range expected {
			if h.Lang(lang) != v {
				t.Errorf("lang=%s,hash=%d results in '%v'. expected '%v'", lang, h, h.Lang(lang), v)
			}
		}
	}
}
