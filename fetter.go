package excel

var FetterInfoExcelConfigData []*FetterInfo

func init() {
	load("ExcelBinOutput/FetterInfoExcelConfigData.json", &FetterInfoExcelConfigData)
}

type FetterInfo struct {
	FetterId                            uint32
	AvatarId                            uint32
	InfoBirthMonth                      uint32
	InfoBirthDay                        uint32
	AvatarNativeTextMapHash             TextMapHash
	AvatarVisionBeforTextMapHash        TextMapHash
	AvatarConstellationBeforTextMapHash TextMapHash
	AvatarTitleTextMapHash              TextMapHash
	AvatarDetailTextMapHash             TextMapHash
	AvatarAssocType                     string // AssocType
	CvChineseTextMapHash                TextMapHash
	CvJapaneseTextMapHash               TextMapHash
	CvEnglishTextMapHash                TextMapHash
	CvKoreanTextMapHash                 TextMapHash
	AvatarVisionAfterTextMapHash        TextMapHash
	AvatarConstellationAfterTextMapHash TextMapHash
}

func (f *FetterInfo) Avatar() *Avatar {
	return FindAvatar(f.AvatarId)
}
