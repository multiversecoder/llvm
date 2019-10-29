// Code generated by "string2enum -linecomment -type DwarfLang ../../ir/enum"; DO NOT EDIT.

package enum

import (
	"fmt"

	"github.com/llir/llvm/ir/enum"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.DwarfLangC89-1]
	_ = x[enum.DwarfLangC-2]
	_ = x[enum.DwarfLangAda83-3]
	_ = x[enum.DwarfLangCPlusPlus-4]
	_ = x[enum.DwarfLangCobol74-5]
	_ = x[enum.DwarfLangCobol85-6]
	_ = x[enum.DwarfLangFortran77-7]
	_ = x[enum.DwarfLangFortran90-8]
	_ = x[enum.DwarfLangPascal83-9]
	_ = x[enum.DwarfLangModula2-10]
	_ = x[enum.DwarfLangJava-11]
	_ = x[enum.DwarfLangC99-12]
	_ = x[enum.DwarfLangAda95-13]
	_ = x[enum.DwarfLangFortran95-14]
	_ = x[enum.DwarfLangPLI-15]
	_ = x[enum.DwarfLangObjC-16]
	_ = x[enum.DwarfLangObjCPlusPlus-17]
	_ = x[enum.DwarfLangUPC-18]
	_ = x[enum.DwarfLangD-19]
	_ = x[enum.DwarfLangPython-20]
	_ = x[enum.DwarfLangOpenCL-21]
	_ = x[enum.DwarfLangGo-22]
	_ = x[enum.DwarfLangModula3-23]
	_ = x[enum.DwarfLangHaskell-24]
	_ = x[enum.DwarfLangCPlusPlus03-25]
	_ = x[enum.DwarfLangCPlusPlus11-26]
	_ = x[enum.DwarfLangOCaml-27]
	_ = x[enum.DwarfLangRust-28]
	_ = x[enum.DwarfLangC11-29]
	_ = x[enum.DwarfLangSwift-30]
	_ = x[enum.DwarfLangJulia-31]
	_ = x[enum.DwarfLangDylan-32]
	_ = x[enum.DwarfLangCPlusPlus14-33]
	_ = x[enum.DwarfLangFortran03-34]
	_ = x[enum.DwarfLangFortran08-35]
	_ = x[enum.DwarfLangRenderScript-36]
	_ = x[enum.DwarfLangBLISS-37]
	_ = x[enum.DwarfLangMipsAssembler-32769]
	_ = x[enum.DwarfLangGoogleRenderScript-36439]
	_ = x[enum.DwarfLangBorlandDelphi-45056]
}

const (
	_DwarfLang_name_0 = "DW_LANG_C89DW_LANG_CDW_LANG_Ada83DW_LANG_C_plus_plusDW_LANG_Cobol74DW_LANG_Cobol85DW_LANG_Fortran77DW_LANG_Fortran90DW_LANG_Pascal83DW_LANG_Modula2DW_LANG_JavaDW_LANG_C99DW_LANG_Ada95DW_LANG_Fortran95DW_LANG_PLIDW_LANG_ObjCDW_LANG_ObjC_plus_plusDW_LANG_UPCDW_LANG_DDW_LANG_PythonDW_LANG_OpenCLDW_LANG_GoDW_LANG_Modula3DW_LANG_HaskellDW_LANG_C_plus_plus_03DW_LANG_C_plus_plus_11DW_LANG_OCamlDW_LANG_RustDW_LANG_C11DW_LANG_SwiftDW_LANG_JuliaDW_LANG_DylanDW_LANG_C_plus_plus_14DW_LANG_Fortran03DW_LANG_Fortran08DW_LANG_RenderScriptDW_LANG_BLISS"
	_DwarfLang_name_1 = "DW_LANG_Mips_Assembler"
	_DwarfLang_name_2 = "DW_LANG_GOOGLE_RenderScript"
	_DwarfLang_name_3 = "DW_LANG_BORLAND_Delphi"
)

var (
	_DwarfLang_index_0 = [...]uint16{0, 11, 20, 33, 52, 67, 82, 99, 116, 132, 147, 159, 170, 183, 200, 211, 223, 245, 256, 265, 279, 293, 303, 318, 333, 355, 377, 390, 402, 413, 426, 439, 452, 474, 491, 508, 528, 541}
)

func DwarfLangFromString(s string) enum.DwarfLang {
	if len(s) == 0 {
		return 0
	}
	for i := range _DwarfLang_index_0[:len(_DwarfLang_index_0)-1] {
		if s == _DwarfLang_name_0[_DwarfLang_index_0[i]:_DwarfLang_index_0[i+1]] {
			return enum.DwarfLang(i + 1)
		}
	}
	if s == _DwarfLang_name_1 {
		return enum.DwarfLang(32769)
	}
	if s == _DwarfLang_name_2 {
		return enum.DwarfLang(36439)
	}
	if s == _DwarfLang_name_3 {
		return enum.DwarfLang(45056)
	}
	panic(fmt.Errorf("unable to locate DwarfLang enum corresponding to %q", s))
}
