package excel

var GadgetExcelConfigData []*Gadget

func init() {
	load("ExcelBinOutput/GadgetExcelConfigData.json", &GadgetExcelConfigData)
}

type Gadget struct {
	Entity
	Type                    string // EntityType
	JsonName                string
	InteractNameTextMapHash TextMapHash
}

func FindGadget(id uint32) *Gadget {
	return Find(GadgetExcelConfigData, func(v *Gadget) bool {
		return v.Id == id
	})
}
