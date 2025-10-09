package excel

var GadgetExcelConfigData []*Gadget

func init() {
	load("ExcelBinOutput/GadgetExcelConfigData.json", &GadgetExcelConfigData)
}

type Gadget struct {
	Type                    string
	JSONName                string
	InteractNameTextMapHash TextMapHash
	Id                      uint32
	CampId                  uint32
}

func FindGadget(id uint32) *Gadget {
	return Find(GadgetExcelConfigData, func(v *Gadget) bool {
		return v.Id == id
	})
}
