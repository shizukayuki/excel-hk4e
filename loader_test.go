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
	}
	setup(t, []string{"en", "jp", "kr", "th"}, override)

	for lang, expected := range expected {
		for h, v := range expected {
			if h.Lang(lang) != v {
				t.Errorf("lang=%s,hash=%d results in '%v'. expected '%v'", lang, h, h.Lang(lang), v)
			}
		}
	}
}
