package chunks

import (
	"fmt"
	"log"
)

// handlePrefLcle processes the PREF.LCLE chunk.
func handlePrefLcle(data []byte) (StructResult, error) {
	log.Println("Handling PREF.LCLE chunk")

	// struct LocalePrefs {
	//     ULONG lp_Reserved[4];
	//     char  lp_RegionName[32];
	//     char  lp_PreferredLanguages[10][30];
	//     LONG  lp_GMTOffset;
	//     ULONG lp_Flags;			/* The same as loc_Flags in struct Locale */

	//     struct CountryPrefs lp_RegionData;
	// };

	// struct RegionPrefs {
	// 	ULONG cp_Reserved[4];

	// 	ULONG cp_RegionCode;
	// 	ULONG cp_TelephoneCode;
	// 	UBYTE cp_MeasuringSystem;

	// 	char  cp_DateTimeFormat[80];
	// 	char  cp_DateFormat[40];
	// 	char  cp_TimeFormat[40];
	// 	char  cp_ShortDateTimeFormat[80];
	// 	char  cp_ShortDateFormat[40];
	// 	char  cp_ShortTimeFormat[40];

	// 	char  cp_DecimalPoint[10];
	// 	char  cp_GroupSeparator[10];
	// 	char  cp_FracGroupSeparator[10];
	// 	UBYTE cp_Grouping[10];
	// 	UBYTE cp_FracGrouping[10];
	// 	char  cp_MonDecimalPoint[10];
	// 	char  cp_MonGroupSeparator[10];
	// 	char  cp_MonFracGroupSeparator[10];
	// 	UBYTE cp_MonGrouping[10];
	// 	UBYTE cp_MonFracGrouping[10];
	// 	UBYTE cp_MonFracDigits;
	// 	UBYTE cp_MonIntFracDigits;

	// 	char  cp_MonCS[10];
	// 	char  cp_MonSmallCS[10];
	// 	char  cp_MonIntCS[10];

	// 	char  cp_MonPositiveSign[10];
	// 	UBYTE cp_MonPositiveSpaceSep;
	// 	UBYTE cp_MonPositiveSignPos;
	// 	UBYTE cp_MonPositiveCSPos;
	// 	char  cp_MonNegativeSign[10];
	// 	UBYTE cp_MonNegativeSpaceSep;
	// 	UBYTE cp_MonNegativeSignPos;
	// 	UBYTE cp_MonNegativeCSPos;

	// 	UBYTE cp_CalendarType;
	// };

	var offset uint32
	var result StructResult

	// Skip ULONG lp_Reserved[4]
	offset += 16

	// handle char lp_RegionName[32]
	lp_RegionName, err := getStringBuffer(data, &offset, 32)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Region Name", lp_RegionName})

	// handle char lp_PreferredLanguages[10][30]
	for i := 0; i < 10; i++ {
		lp_PreferredLanguages, err := getStringBuffer(data, &offset, 30)
		if err != nil {
			return result, err
		}
		result = append(result, [2]string{"Preferred Language", lp_PreferredLanguages})
	}

	// handle LONG lp_GMTOffset
	lp_GMTOffset, err := getBeLong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"GMT Offset", fmt.Sprintf("%d", lp_GMTOffset)})

	// handle ULONG lp_Flags
	lp_Flags, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Flags", fmt.Sprintf("%032b", lp_Flags)})

	// Read struct CountryPrefs lp_RegionData

	// Skip ULONG cp_Reserved[4]
	offset += 16

	// handle ULONG cp_RegionCode
	cp_RegionCode, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Region Code", fmt.Sprintf("%d", cp_RegionCode)})

	// handle ULONG cp_TelephoneCode
	cp_TelephoneCode, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Telephone Code", fmt.Sprintf("%d", cp_TelephoneCode)})

	// handle UBYTE cp_MeasuringSystem
	cp_MeasuringSystem, err := getByte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Measuring System", fmt.Sprintf("%d", cp_MeasuringSystem)})

	// handle char cp_DateTimeFormat[80]
	cp_DateTimeFormat, err := getStringBuffer(data, &offset, 80)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"DateTime Format", cp_DateTimeFormat})

	// handle char cp_DateFormat[40]
	cp_DateFormat, err := getStringBuffer(data, &offset, 40)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Date Format", cp_DateFormat})

	// handle char cp_TimeFormat[40]
	cp_TimeFormat, err := getStringBuffer(data, &offset, 40)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Time Format", cp_TimeFormat})

	// handle char cp_ShortDateTimeFormat[80]
	cp_ShortDateTimeFormat, err := getStringBuffer(data, &offset, 80)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Short DateTime Format", cp_ShortDateTimeFormat})

	// handle char cp_ShortDateFormat[40]
	cp_ShortDateFormat, err := getStringBuffer(data, &offset, 40)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Short Date Format", cp_ShortDateFormat})

	// handle char cp_ShortTimeFormat[40]
	cp_ShortTimeFormat, err := getStringBuffer(data, &offset, 40)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Short Time Format", cp_ShortTimeFormat})

	// handle char cp_DecimalPoint[10]
	cp_DecimalPoint, err := getStringBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Decimal Point", cp_DecimalPoint})

	// handle char cp_GroupSeparator[10]
	cp_GroupSeparator, err := getStringBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Group Separator", cp_GroupSeparator})

	// handle char cp_FracGroupSeparator[10]
	cp_FracGroupSeparator, err := getStringBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Frac Group Separator", cp_FracGroupSeparator})

	// handle UBYTE cp_Grouping[10]
	cp_Grouping, err := getByteBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Grouping", fmt.Sprintf("%v", cp_Grouping)})

	// handle UBYTE cp_FracGrouping[10]
	cp_FracGrouping, err := getByteBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Frac Grouping", fmt.Sprintf("%v", cp_FracGrouping)})

	// handle char cp_MonDecimalPoint[10]
	cp_MonDecimalPoint, err := getStringBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Decimal Point", cp_MonDecimalPoint})

	// handle char cp_MonGroupSeparator[10]
	cp_MonGroupSeparator, err := getStringBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Group Separator", cp_MonGroupSeparator})

	// handle char cp_MonFracGroupSeparator[10]
	cp_MonFracGroupSeparator, err := getStringBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Frac Group Separator", cp_MonFracGroupSeparator})

	// handle UBYTE cp_MonGrouping[10]
	cp_MonGrouping, err := getByteBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Grouping", fmt.Sprintf("%v", cp_MonGrouping)})

	// handle UBYTE cp_MonFracGrouping[10]
	cp_MonFracGrouping, err := getByteBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Frac Grouping", fmt.Sprintf("%v", cp_MonFracGrouping)})

	// handle UBYTE cp_MonFracDigits
	cp_MonFracDigits, err := getByte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Frac Digits", fmt.Sprintf("%d", cp_MonFracDigits)})

	// handle UBYTE cp_MonIntFracDigits
	cp_MonIntFracDigits, err := getByte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Int Frac Digits", fmt.Sprintf("%d", cp_MonIntFracDigits)})

	// handle char cp_MonCS[10]
	cp_MonCS, err := getStringBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon CS", cp_MonCS})

	// handle char cp_MonSmallCS[10]
	cp_MonSmallCS, err := getStringBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Small CS", cp_MonSmallCS})

	// handle char cp_MonIntCS[10]
	cp_MonIntCS, err := getStringBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Int CS", cp_MonIntCS})

	// handle char cp_MonPositiveSign[10]
	cp_MonPositiveSign, err := getStringBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Positive Sign", cp_MonPositiveSign})

	// handle UBYTE cp_MonPositiveSpaceSep
	cp_MonPositiveSpaceSep, err := getByte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Positive Space Sep", fmt.Sprintf("%d", cp_MonPositiveSpaceSep)})

	// handle UBYTE cp_MonPositiveSignPos
	cp_MonPositiveSignPos, err := getByte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Positive Sign Pos", fmt.Sprintf("%d", cp_MonPositiveSignPos)})

	// handle UBYTE cp_MonPositiveCSPos
	cp_MonPositiveCSPos, err := getByte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Positive CS Pos", fmt.Sprintf("%d", cp_MonPositiveCSPos)})

	// handle char cp_MonNegativeSign[10]
	cp_MonNegativeSign, err := getStringBuffer(data, &offset, 10)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Negative Sign", cp_MonNegativeSign})

	// handle UBYTE cp_MonNegativeSpaceSep
	cp_MonNegativeSpaceSep, err := getByte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Negative Space Sep", fmt.Sprintf("%d", cp_MonNegativeSpaceSep)})

	// handle UBYTE cp_MonNegativeSignPos
	cp_MonNegativeSignPos, err := getByte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Negative Sign Pos", fmt.Sprintf("%d", cp_MonNegativeSignPos)})

	// handle UBYTE cp_MonNegativeCSPos
	cp_MonNegativeCSPos, err := getByte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mon Negative CS Pos", fmt.Sprintf("%d", cp_MonNegativeCSPos)})

	// handle UBYTE cp_CalendarType
	cp_CalendarType, err := getByte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Calendar Type", fmt.Sprintf("%d", cp_CalendarType)})

	return result, nil
}

// handlePrefOscn processes the PREF.OSCN chunk.
func handlePrefOscn(data []byte) (StructResult, error) {
	log.Println("Handling PREF.OSCN chunk")

	// struct OverscanPrefs {
	//     ULONG os_Reserved;
	//     ULONG os_Magic;
	//     UWORD os_HStart;
	//     UWORD os_HStop;
	//     UWORD os_VStart;
	//     UWORD os_VStop;
	//     ULONG os_DisplayID;
	//     Point os_ViewPos;
	//     Point os_Text;
	//     Rectangle os_Standard;
	// };

	var offset uint32
	var result StructResult

	// Skip ULONG os_Reserved
	offset += 4

	// handle ULONG os_Magic
	os_Magic, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Magic", fmt.Sprintf("%d", os_Magic)})

	// handle UWORD os_HStart
	os_HStart, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"HStart", fmt.Sprintf("%d", os_HStart)})

	// handle UWORD os_HStop
	os_HStop, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"HStop", fmt.Sprintf("%d", os_HStop)})

	// handle UWORD os_VStart
	os_VStart, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"VStart", fmt.Sprintf("%d", os_VStart)})

	// handle UWORD os_VStop
	os_VStop, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"VStop", fmt.Sprintf("%d", os_VStop)})

	// handle ULONG os_DisplayID
	os_DisplayID, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"DisplayID", fmt.Sprintf("%032b", os_DisplayID)})

	// handle Point os_ViewPos
	os_ViewPosX, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	os_ViewPosY, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"ViewPos", fmt.Sprintf("(%d, %d)", os_ViewPosX, os_ViewPosY)})

	// handle Point os_Text
	os_TextX, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	os_TextY, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Text", fmt.Sprintf("(%d, %d)", os_TextX, os_TextY)})

	// handle struct Rectangle os_Standard (4 WORDs)
	os_StandardMinX, err := getBeWord(data, &offset)
	if err != nil {
		return result, err
	}
	os_StandardMinY, err := getBeWord(data, &offset)
	if err != nil {
		return result, err
	}
	os_StandardMaxX, err := getBeWord(data, &offset)
	if err != nil {
		return result, err
	}
	os_StandardMaxY, err := getBeWord(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Standard", fmt.Sprintf("(%d, %d, %d, %d)",
		os_StandardMinX, os_StandardMinY, os_StandardMaxX, os_StandardMaxY)})

	return result, nil
}

func handlePrefPalt(data []byte) (StructResult, error) {
	log.Println("Handling PREF.PALT chunk")

	// struct PalettePrefs
	// {
	// 	LONG	     pap_Reserved[4];
	// 	UWORD	     pap_4ColorPens[32];
	// 	UWORD	     pap_8ColorPens[32];
	// 	struct ColorSpec pap_Colors[32];
	// } __packed;

	var offset uint32
	var result StructResult

	// Skip ULONG pap_Reserved[4]
	offset += 16

	// handle UWORD pap_4ColorPens[32]
	for i := 0; i < 32; i++ {
		pap_4ColorPens, err := getBeUword(data, &offset)
		if err != nil {
			return result, err
		}
		result = append(result, [2]string{fmt.Sprintf("4 Color Pen %d", i),
			fmt.Sprintf("%d", pap_4ColorPens)})
	}

	// handle UWORD pap_8ColorPens[32]
	for i := 0; i < 32; i++ {
		pap_8ColorPens, err := getBeUword(data, &offset)
		if err != nil {
			return result, err
		}
		result = append(result, [2]string{fmt.Sprintf("8 Color Pen %d", i),
			fmt.Sprintf("%d", pap_8ColorPens)})
	}

	// 	struct ColorSpec
	// {
	//     WORD  ColorIndex;
	//     UWORD Red;
	//     UWORD Green;
	//     UWORD Blue;
	// };

	// handle struct ColorSpec pap_Colors[32]
	for i := 0; i < 32; i++ {
		// handle WORD ColorIndex
		ColorIndex, err := getBeWord(data, &offset)
		if err != nil {
			return result, err
		}

		// handle UWORD Red
		Red, err := getBeUword(data, &offset)
		if err != nil {
			return result, err
		}

		// handle UWORD Green
		Green, err := getBeUword(data, &offset)
		if err != nil {
			return result, err
		}

		// handle UWORD Blue
		Blue, err := getBeUword(data, &offset)
		if err != nil {
			return result, err
		}

		result = append(result, [2]string{fmt.Sprintf("Color %d", i),
			fmt.Sprintf("%d: %d, %d, %d", ColorIndex, Red, Green, Blue)})
	}

	return result, nil
}

// handlePrefPntr processes the PREF.PNTR chunk.
func handlePrefPntr(data []byte) (StructResult, error) {
	log.Println("Handling PREF.PNTR chunk")

	// struct PointerPrefs
	// {
	//     ULONG pp_Reserved[4];
	//     UWORD pp_Which;
	//     UWORD pp_Size;
	//     UWORD pp_Width;
	//     UWORD pp_Height;
	//     UWORD pp_Depth;
	//     UWORD pp_YSize;
	//     UWORD pp_X, pp_Y;
	// };

	var offset uint32
	var result StructResult

	// Skip ULONG pp_Reserved[4]
	offset += 16

	// handle UWORD pp_Which
	pp_Which, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Which", fmt.Sprintf("%d", pp_Which)})

	// handle UWORD pp_Size
	pp_Size, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Size", fmt.Sprintf("%d", pp_Size)})

	// handle UWORD pp_Width
	pp_Width, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Width", fmt.Sprintf("%d", pp_Width)})

	// handle UWORD pp_Height
	pp_Height, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Height", fmt.Sprintf("%d", pp_Height)})

	// handle UWORD pp_Depth
	pp_Depth, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Depth", fmt.Sprintf("%d", pp_Depth)})

	// handle UWORD pp_YSize
	pp_YSize, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"YSize", fmt.Sprintf("%d", pp_YSize)})

	// handle UWORD pp_X
	pp_X, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}

	// handle UWORD pp_Y
	pp_Y, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Position", fmt.Sprintf("(%d, %d)", pp_X, pp_Y)})

	return result, nil
}

// handlePrefNptr processes the PREF.NPTR chunk.
func handlePrefNptr(data []byte) (StructResult, error) {
	log.Println("Handling PREF.NPTR chunk")

	// struct NewPointerPrefs
	// {
	//     UWORD npp_Which;
	//     UWORD npp_AlphaValue;
	//     ULONG npp_WhichInFile;
	//     UWORD npp_X, npp_Y;
	//     char  npp_File[0];
	// };

	var offset uint32
	var result StructResult

	// handle UWORD npp_Which
	npp_Which, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Which", fmt.Sprintf("%d", npp_Which)})

	// handle UWORD npp_AlphaValue
	npp_AlphaValue, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Alpha Value", fmt.Sprintf("%d", npp_AlphaValue)})

	// handle ULONG npp_WhichInFile
	npp_WhichInFile, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Which In File", fmt.Sprintf("%d", npp_WhichInFile)})

	// handle UWORD npp_X
	npp_X, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}

	// handle UWORD npp_Y
	npp_Y, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Hotspot Coordinates", fmt.Sprintf("(%d, %d)", npp_X, npp_Y)})

	// handle char npp_File[0]
	// Read until the end of the chunk
	npp_File, err := getStringBuffer(data, &offset, uint32(len(data)-int(offset)))
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"File", npp_File})

	return result, nil
}
