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
	for _, v := range GadgetExcelConfigData {
		if v.Id == id {
			return v
		}
	}
	return nil
}
