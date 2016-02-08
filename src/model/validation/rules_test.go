package validation_test
import (
	"testing"
	"model/validation"
)

func Test_RequireNotEmpty(t *testing.T){
	input := "    not empty text "
	passed, msg := validation.RequireNotEmpty(input)
	if passed==false || msg !="" {
		t.Fail()
	}

	emptyinput := "     "
	passed, msg = validation.RequireNotEmpty(emptyinput)
	if passed==true || msg =="" {
		t.Fail()
	}
}

func Test_RequireAlphaNumeric(t *testing.T){
	input :=  "sgsgb453"
	passed, msg := validation.RequireAlphaNumeric(input)
	if passed==false || msg !="" {
		t.Fail()
	}
	input =  "sgsgb"
	passed, msg = validation.RequireAlphaNumeric(input)
	if passed==false || msg !="" {
		t.Fail()
	}
	input =  "4468"
	passed, msg = validation.RequireAlphaNumeric(input)
	if passed==false || msg !="" {
		t.Fail()
	}
	input =  " 4468 "
	passed, msg = validation.RequireAlphaNumeric(input)
	if passed==false || msg !="" {
		t.Fail()
	}
	input =  "ggjfgjgf djdj"
	passed, msg = validation.RequireAlphaNumeric(input)
	if passed==true || msg =="" {
		t.Fail()
	}
	input =  "44,44454"
	passed, msg = validation.RequireAlphaNumeric(input)
	if passed==true || msg =="" {
		t.Fail()
	}
	input =  "4df9#%"
	passed, msg = validation.RequireAlphaNumeric(input)
	if passed==true || msg =="" {
		t.Fail()
	}
}

func Test_RequireEmail(t *testing.T){
	input :=  "lus_ka@gkjka.mgd"
	passed, msg := validation.RequireEmail(input)
	if passed==false || msg !="" {
		t.Fail()
	}
	input =  "sgsgb@dh.ru"
	passed, msg = validation.RequireEmail(input)
	if passed==false || msg !="" {
		t.Fail()
	}
	input =  "hfh4468-lu@gsdg.ru"
	passed, msg = validation.RequireEmail(input)
	if passed==false || msg !="" {
		t.Fail()
	}
	input =  " 44$%%&68 "
	passed, msg = validation.RequireEmail(input)
	if passed==true || msg =="" {
		t.Fail()
	}
	input =  "ggjfgjgf djdj"
	passed, msg = validation.RequireEmail(input)
	if passed==true || msg =="" {
		t.Fail()
	}
	input =  "44,44454.vb"
	passed, msg = validation.RequireEmail(input)
	if passed==true || msg =="" {
		t.Fail()
	}
	input =  "sgsgb@dhdru"
	passed, msg = validation.RequireEmail(input)
	if passed==true || msg =="" {
		t.Fail()
	}
}

//func Test_RequireStrongPassword(t *testing.T){
//	type boolCombination [8][3]bool
//	strengthVersions :=boolCombination{
//		[3]bool{true, true, true},
//		[3]bool{true, true, false},
//		[3]bool{true, false, false},
//		[3]bool{true, false, true},
//		[3]bool{false, true, true},
//		[3]bool{false, true, false},
//		[3]bool{false, false, true},
//		[3]bool{false, false, false},
//	}
//	input := "fgW%d9"
//	//var i int
//	for i:=0; i<8; i++{
//		boolComb := strengthVersions[i]
//		upper, symbol, number := boolComb[0], boolComb[1], boolComb[2]
//		passed, msg :=validation.RequireStrongPassword(input, upper, symbol, number)
//	}
//	passed, msg := validation.RequireStrongPassword(input, true, true, true)
//	if passed==false || msg !="" {
//		t.Fail()
//	}
//	passed, msg = validation.RequireStrongPassword(input, false, false, false)
//	if passed==false || msg !="" {
//		t.Fail()
//	}
//	passed, msg = validation.RequireStrongPassword(input, true, true, false)
//	if passed==false || msg !="" {
//		t.Fail()
//	}
//	passed, msg = validation.RequireStrongPassword(input, true, false, true)
//	if passed==false || msg !="" {
//		t.Fail()
//	}
//	passed, msg = validation.RequireStrongPassword(input, false, true, true)
//	if passed==false || msg !="" {
//		t.Fail()
//	}
//	input = "fgw%d9"
//
//	passed, msg = validation.RequireStrongPassword(input, true, true, true)
//	if passed==true || msg =="" {
//		t.Fail()
//	}
//	passed, msg = validation.RequireStrongPassword(input, false, false, false)
//	if passed==false || msg !="" {
//		t.Fail()
//	}
//	passed, msg = validation.RequireStrongPassword(input, true, true, false)
//	if passed==false || msg !="" {
//		t.Fail()
//	}
//	passed, msg = validation.RequireStrongPassword(input, true, false, true)
//	if passed==false || msg !="" {
//		t.Fail()
//	}
//	passed, msg = validation.RequireStrongPassword(input, false, true, true)
//	if passed==false || msg !="" {
//		t.Fail()
//	}
//}

func Test_RequireUppercase(t *testing.T){
	input := "gshsdhjn5453./hdf3$$&@"
		passed, msg := validation.RequireUppercase(input)
		if passed==true || msg =="" {
			t.Fail()
		}
	input = "gdgj5435^#G"
	passed, msg = validation.RequireUppercase(input)
	if passed==false || msg !="" {
		t.Fail()
	}
}

func Test_RequireSymbol(t *testing.T){
	input := "aafgj45#"
		passed, msg := validation.RequireSymbol(input)
		if passed==false || msg !="" {
			t.Fail()
		}
	input = "hdhgdh$"
		passed, msg = validation.RequireSymbol(input)
		if passed==false || msg !="" {
			t.Fail()
		}
	input = "hdhgdh/"
		passed, msg = validation.RequireSymbol(input)
		if passed==false || msg !="" {
			t.Fail()
		}
	input = "hdhgdh!"
		passed, msg = validation.RequireSymbol(input)
		if passed==false || msg !="" {
			t.Fail()
		}
	input = "hd*gdh"
		passed, msg = validation.RequireSymbol(input)
		if passed==false || msg !="" {
			t.Fail()
		}
	input = "jfj6"
	passed, msg = validation.RequireSymbol(input)
	if passed==true || msg =="" {
		t.Fail()
	}
	input = "j fj6"
	passed, msg = validation.RequireSymbol(input)
	if passed==false || msg !="" {
		t.Fail()
	}
}
func Test_RequireNumber(t *testing.T){
	input := "hd6hsd4hs"
	passed, msg := validation.RequireNumber(input)
	if passed==false || msg !="" {
		t.Fail()
	}
	input = "jgdsgfj"
	passed, msg = validation.RequireNumber(input)
	if passed==true || msg =="" {
		t.Fail()
	}
}
func Test_RequireStrongPassword(t *testing.T){
	input := "hdj87&"
	passed, msg := validation.RequireStrongPassword(input)
	if passed==true || msg =="" {
		t.Fail()
	}
	input = "hdj87A"
	passed, msg = validation.RequireStrongPassword(input)
	if passed==true || msg =="" {
		t.Fail()
	}
	input = "hdj%A"
	passed, msg = validation.RequireStrongPassword(input)
	if passed==true || msg =="" {
		t.Fail()
	}
	input = "hdjdgsd8"
	passed, msg = validation.RequireStrongPassword(input)
	if passed==true || msg =="" {
		t.Fail()
	}
	input = "gG8ye^"
	passed, msg = validation.RequireNumber(input)
	if passed==false || msg !="" {
		t.Fail()
	}
}

func Test_RequireFirstLetter(t *testing.T){
	input := "#hdg987"
	passed, msg := validation.RequireFirstLetter(input)
	if passed==true || msg =="" {
		t.Fail()
	}
	input = "8hdg987"
	passed, msg = validation.RequireFirstLetter(input)
	if passed==true || msg =="" {
		t.Fail()
	}
	input = "A8hdg987"
	passed, msg = validation.RequireFirstLetter(input)
	if passed==false || msg !="" {
		t.Fail()
	}
	input = "d8hdg987"
	passed, msg = validation.RequireFirstLetter(input)
	if passed==false || msg !="" {
		t.Fail()
	}
}

func Test_RequireMinCount(t *testing.T){
	input := "gsdsdh"
	passed, msg := validation.RequireMinCount(input,8)
	if passed==true || msg =="" {
		t.Fail()
	}
	passed, msg = validation.RequireMinCount(input,6)
	if passed==false || msg !="" {
		t.Fail()
	}
	passed, msg = validation.RequireMinCount(input,4)
	if passed==false || msg !="" {
		t.Fail()
	}
}