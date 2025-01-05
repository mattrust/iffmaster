// Copyright (c) 2025 Matthias Rustler
// Licensed under the MIT License - see LICENSE for details

package chunks

import (
	"fmt"
	"log"
)

// handlePrefPrhd processes the PREF.PRHD chunk.
func handlePrefPrhd(data []byte) (StructResult, error) {
	log.Println("Handling PREF.PRHD chunk")

	//struct PrefHeader
	//{
	//	UBYTE ph_Version;
	//	UBYTE ph_Type;
	//	ULONG ph_Flags;
	//};

	var offset uint32
	var result StructResult

	// handle ph_Version
	phVersion, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Version", fmt.Sprintf("%d", phVersion)})

	// handle ph_Type
	phType, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Type", fmt.Sprintf("%d", phType)})

	// handle ph_Flags
	phFlags, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Flags", fmt.Sprintf("%032b", phFlags)})

	return result, nil
}

// handlePrefAsl processes the PREF.ASL chunk.
func handlePrefAsl(data []byte) (StructResult, error) {
	log.Println("Handling PREF.ASL chunk")

	//struct AslPrefs
	//{
	//    LONG    ap_Reserved[4];
	//    UBYTE   ap_SortBy;
	//    UBYTE   ap_SortDrawers;
	//    UBYTE   ap_SortOrder;
	//    UBYTE   ap_SizePosition;
	//    WORD    ap_RelativeLeft;
	//    WORD    ap_RelativeTop;
	//    UBYTE   ap_RelativeWidth;
	//    UBYTE   ap_RelativeHeight;
	//} __packed;

	var offset uint32
	var result StructResult

	// Skip ap_Reserved
	offset += 4 * 4

	// handle ap_SortBy
	apSortBy, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Sort By", fmt.Sprintf("%d", apSortBy)})

	// handle ap_SortDrawers
	apSortDrawers, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Sort Drawers", fmt.Sprintf("%d", apSortDrawers)})

	// handle ap_SortOrder
	apSortOrder, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Sort Order", fmt.Sprintf("%d", apSortOrder)})

	// handle ap_SizePosition
	apSizePosition, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Size Position", fmt.Sprintf("%d", apSizePosition)})

	// handle ap_RelativeLeft
	apRelativeLeft, err := getBeWord(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Relative Left", fmt.Sprintf("%d", apRelativeLeft)})

	// handle ap_RelativeTop
	apRelativeTop, err := getBeWord(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Relative Top", fmt.Sprintf("%d", apRelativeTop)})

	// handle ap_RelativeWidth
	apRelativeWidth, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"RelativeWidth", fmt.Sprintf("%d", apRelativeWidth)})

	// handle ap_RelativeHeight
	apRelativeHeight, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"RelativeHeight", fmt.Sprintf("%d", apRelativeHeight)})

	return result, nil
}

// handlePrefFont processes the PREF.FONT chunk.
func handlePrefFont(data []byte) (StructResult, error) {
	log.Println("Handling PREF.FONT chunk")

	// struct FontPrefs
	// {
	//     LONG            fp_Reserved[3]; /* PRIVATE */
	//     UWORD           fp_Reserved2;   /* PRIVATE */
	//     UWORD           fp_Type;        /* see below */
	//     UBYTE           fp_FrontPen;
	//     UBYTE           fp_BackPen;
	//     UBYTE           fp_DrawMode;
	//     UBYTE           fp_pad;
	//     struct TextAttr fp_TextAttr;
	//     BYTE            fp_Name[FONTNAMESIZE];
	// };

	// struct TextAttr
	// {
	//     STRPTR ta_Name;
	//     UWORD  ta_YSize;
	//     UBYTE  ta_Style;
	//     UBYTE  ta_Flags;
	// };

	const FONTNAMESIZE = 128

	var offset uint32
	var result StructResult

	// Skip fp_Reserved
	offset += 4 * 3

	// Skip fp_Reserved2
	offset += 2

	// handle fp_Type
	fpType, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	if fpType == 0 {
		result = append(result, [2]string{"Type", "WBFONT"})
	} else if fpType == 1 {
		result = append(result, [2]string{"Type", "SYSFONT"})
	} else if fpType == 2 {
		result = append(result, [2]string{"Type", "SCREENFONT"})
	}

	// handle fp_FrontPen
	fpFrontPen, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Front Pen", fmt.Sprintf("%d", fpFrontPen)})

	// handle fp_BackPen
	fpBackPen, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Back Pen", fmt.Sprintf("%d", fpBackPen)})

	// handle fp_DrawMode
	fpDrawmode, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Drawmode", fmt.Sprintf("%0b", fpDrawmode)})

	// Skip fp_pad
	offset++

	// Skip fp_TextAttr_ta_Name because it's a STRPTR
	offset += 4

	// struct TextAttr
	fpTextAttrTaYSize, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Size", fmt.Sprintf("%d", fpTextAttrTaYSize)})

	// handle fp_TextAttr_ta_Style
	fpTextAttrTaStyle, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Style", fmt.Sprintf("%d", fpTextAttrTaStyle)})

	// handle fp_TextAttr_ta_Flags
	fpTextAttrTaFlags, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"TextAttr_ta_Flags",
		fmt.Sprintf("%0b", fpTextAttrTaFlags)})

	// handle fp_Name
	fpName, err := getStringBuffer(data, &offset, FONTNAMESIZE)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Name", fpName})

	return result, nil
}

// handlePrefIctl processes the PREF.ICTL chunk.
func handlePrefIctl(data []byte) (StructResult, error) {
	log.Println("Handling PREF.ICTL chunk")

	// struct IControlPrefs {
	// 	LONG  ic_Reserved[4];
	// 	UWORD ic_TimeOut;
	// 	WORD  ic_MetaDrag;
	// 	ULONG ic_Flags;
	// 	UBYTE ic_WBtoFront;
	// 	UBYTE ic_FrontToBack;
	// 	UBYTE ic_ReqTrue;
	// 	UBYTE ic_ReqFalse;
	// 	UWORD ic_Reserved2;
	// 	UWORD ic_VDragModes[2];
	// };

	var offset uint32
	var result StructResult

	// Skip ic_Reserved
	offset += 4 * 4

	// handle ic_TimeOut
	icTimeOut, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Timeout", fmt.Sprintf("%d", icTimeOut)})

	// handle ic_MetaDrag
	icMetaDrag, err := getBeWord(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Meta Drag", fmt.Sprintf("%0b", icMetaDrag)})

	// handle ic_Flags
	icFlags, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	if icFlags&(1<<0) != 0 {
		result = append(result, [2]string{"Flag", "ICF_NOACTIVEWINDOW"})
	}
	if icFlags&(1<<1) != 0 {
		result = append(result, [2]string{"Flag", "ICF_COERCE_LACE"})
	}
	if icFlags&(1<<2) != 0 {
		result = append(result, [2]string{"Flag", "ICF_STRGAD_FILTER"})
	}
	if icFlags&(1<<3) != 0 {
		result = append(result, [2]string{"Flag", "ICF_MENUSNAP"})
	}
	if icFlags&(1<<4) != 0 {
		result = append(result, [2]string{"Flag", "ICF_MODEPROMOTE"})
	}
	if icFlags&(1<<31) != 0 {
		result = append(result, [2]string{"Flag", "ICF_STICKYMENUS (MorphOS)"})
	}
	if icFlags&(1<<30) != 0 {
		result = append(result, [2]string{"Flag", "ICF_OPAQUEMOVE (MorphOS)"})
	}
	if icFlags&(1<<29) != 0 {
		result = append(result, [2]string{"Flag", "ICF_PRIVILEDGEDREFRESH (MorphOS)"})
	}
	if icFlags&(1<<28) != 0 {
		result = append(result, [2]string{"Flag", "ICF_OFFSCREENLAYERS (MorphOS)"})
	}
	if icFlags&(1<<27) != 0 {
		result = append(result, [2]string{"Flag", "ICF_DEFPUBSCREEN (MorphOS)"})
	}
	if icFlags&(1<<26) != 0 {
		result = append(result, [2]string{"Flag", "ICF_SCREENACTIVATION (MorphOS)"})
	}
	if icFlags&(1<<17) != 0 {
		result = append(result, [2]string{"Flag", "ICF_PULLDOWNTITLEMENUS (AROS)"})
	}
	if icFlags&(1<<16) != 0 {
		result = append(result, [2]string{"Flag", "ICF_POPUPMENUS (AROS)"})
	}
	if icFlags&(1<<15) != 0 {
		result = append(result, [2]string{"Flag", "ICF_3DMENUS (AROS)"})
	}
	if icFlags&(1<<14) != 0 {
		result = append(result, [2]string{"Flag", "ICF_AVOIDWINBORDERERASE (AROS)"})
	}

	//
	icWBtoFront, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"WBtoFront", fmt.Sprintf("%d", icWBtoFront)})

	// handle ic_FrontToBack
	icFrontToBack, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"FrontToBack", fmt.Sprintf("%d", icFrontToBack)})

	// handle ic_ReqTrue
	icReqTrue, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"ReqTrue", fmt.Sprintf("%d", icReqTrue)})

	// handle ic_ReqFalse
	icReqFalse, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"ReqFalse", fmt.Sprintf("%d", icReqFalse)})

	// Skip ic_Reserved2
	offset += 2

	for i := 0; i < 2; i++ {
		// TODO: Parse ic_VDragModes
		icVDragModes, err := getBeUword(data, &offset)
		if err != nil {
			return result, err
		}
		result = append(result, [2]string{fmt.Sprintf("VDragModes %d", i),
			fmt.Sprintf("%d", icVDragModes)})
	}

	return result, nil
}

// handlePrefInpt processes the PREF.INPT chunk.
func handlePrefInpt(data []byte) (StructResult, error) {
	log.Println("Handling PREF.INPT chunk")

	// struct FileInputPrefs
	// {
	//     char    ip_Keymap[16];
	//     UBYTE   ip_PointerTicks[2];
	//     UBYTE   ip_DoubleClick_secs[4];
	//     UBYTE   ip_DoubleClick_micro[4];
	//     UBYTE   ip_KeyRptDelay_secs[4];
	//     UBYTE   ip_KeyRptDelay_micro[4];
	//     UBYTE   ip_KeyRptSpeed_secs[4];
	//     UBYTE   ip_KeyRptSpeed_micro[4];
	//     UBYTE   ip_MouseAccel[2];
	//     UBYTE   ip_ClassicKeyboard[4];
	//     char    ip_KeymapName[64];
	//     UBYTE   ip_SwitchMouseButtons[4];
	// };

	// struct InputPrefs {
	// 	char           ip_Keymap[16];
	// 	UWORD          ip_PointerTicks;
	// 	struct timeval ip_DoubleClick;
	// 	struct timeval ip_KeyRptDelay;
	// 	struct timeval ip_KeyRptSpeed;
	// 	WORD           ip_MouseAccel;

	// 	ULONG          ip_ClassicKeyboard;
	// 	char           ip_KeymapName[64];
	// 	ULONG          ip_SwitchMouseButtons;
	// };

	// struct timeval
	// {
	//     ULONG tv_secs;
	//     ULONG tv_micro;
	// };

	const KEYMAPSIZE = 16
	const KEYMAPNAMESIZE = 64

	var offset uint32
	var result StructResult

	// handle ip_Keymap
	ipKeymap, err := getStringBuffer(data, &offset, KEYMAPSIZE)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Keymap", ipKeymap})

	// handle ip_PointerTicks
	ipPointerTicks, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Pointer Ticks",
		fmt.Sprintf("%d", ipPointerTicks)})

	// handle ip_DoubleClick_secs
	ipDoubleClickSecs, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"DoubleClick Seconds",
		fmt.Sprintf("%d", ipDoubleClickSecs)})

	// handle ip_DoubleClick_micro
	ipDoubleClickMicro, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"DoubleClick Micro",
		fmt.Sprintf("%d", ipDoubleClickMicro)})

	// handle ip_KeyRptDelaySecs
	ipKeyRptDelaySecs, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Key Repeat Seconds",
		fmt.Sprintf("%d", ipKeyRptDelaySecs)})

	// handle ip_KeyRptDelayMicro
	ipKeyRptDelayMicro, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Key Repeat Delay Micro",
		fmt.Sprintf("%d", ipKeyRptDelayMicro)})

	// handle ip_KeyRptSpeedSecs
	ipKeyRptSpeedSecs, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Key Repeat Speed Seconds",
		fmt.Sprintf("%d", ipKeyRptSpeedSecs)})

	// handle ip_KeyRptSpeedMicro
	ipKeyRptSpeedMicro, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Key Repeat Speed Micro",
		fmt.Sprintf("%d", ipKeyRptSpeedMicro)})

	// handle ip_MouseAccel
	ipMouseAccel, err := getBeWord(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Mouse Acceleration",
		fmt.Sprintf("%d", ipMouseAccel)})

	// handle ip_ClassicKeyboard
	ipClassicKeyboard, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Classic Keyboard",
		fmt.Sprintf("%d", ipClassicKeyboard)})

	// handle ipKeymapName
	ipKeymapName, err := getStringBuffer(data, &offset, KEYMAPNAMESIZE)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"KeymapName", ipKeymapName})

	// handle ipSwitchMouseButtons
	ipSwitchMouseButtons, err := getBeUlong(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Switch Mouse Buttons",
		fmt.Sprintf("%d", ipSwitchMouseButtons)})

	return result, nil
}

// handlePrefKMSW processes the PREF.KMSW chunk.
func handlePrefKmsw(data []byte) (StructResult, error) {
	log.Println("Handling PREF.KMSW chunk")

	// 	struct KMSPrefs
	// {
	//     UBYTE kms_Enabled;
	//     UBYTE kms_Reserved;
	//     UWORD kms_SwitchQual;
	//     UWORD kms_SwitchCode;
	//     char  kms_AltKeymap[64];
	// };

	const ALTKEYMAPSIZE = 64

	var offset uint32
	var result StructResult

	// handle kms_Enabled
	kmsEnabled, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Enabled", fmt.Sprintf("%d", kmsEnabled)})

	// handle kms_Reserved
	kmsReserved, err := getUbyte(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Reserved", fmt.Sprintf("%d", kmsReserved)})

	// handle kms_SwitchQual
	kmsSwitchQual, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Switch Qualifier", fmt.Sprintf("%032b", kmsSwitchQual)})

	// handle kms_SwitchCode
	kmsSwitchCode, err := getBeUword(data, &offset)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Switch Code", fmt.Sprintf("%032b", kmsSwitchCode)})

	// handle kms_AltKeymap
	kmsAltKeymap, err := getStringBuffer(data, &offset, ALTKEYMAPSIZE)
	if err != nil {
		return result, err
	}
	result = append(result, [2]string{"Alternative Keymap", kmsAltKeymap})

	return result, nil
}
