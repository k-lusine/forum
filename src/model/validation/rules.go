package validation
import (
	"strings"
	"regexp"
	//"regexp/syntax"
)

func RequireNotEmpty(input string)(passed bool, errorMessage string){
	if strings.Trim(input," ") !="" {
		passed = true
	} else{
		errorMessage = "Some required data is empty"
	}
	return
}

func RequireAlphaNumeric (input string)(passed bool, errorMessage string){
	input = strings.Trim(input," ")
	passed, _ = regexp.MatchString("^[a-zA-Z0-9]+$",input)
	if passed == false{
		errorMessage = "Only alphanumeric input is allowed"
	}
	return
}

func RequireEmail (input string)(passed bool, errorMessage string){
	input = strings.Trim(input," ")
	passed, _ = regexp.MatchString(`^[-_a-zA-Z0-9]+@[a-zA-Z]+\.[a-zA-Z]+$`,input)
	if passed == false{
		errorMessage = "Please provide a valid email address"
	}
	return
}

func RequireUppercase(input string)(passed bool, errorMessage string){
	input = strings.Trim(input," ")
	passed,_ = regexp.MatchString("^.*[A-Z].*$",input)
		if passed == false{
			errorMessage = "There should be a uppercase letter in the input"
		}
	return
}
func RequireSymbol(input string)(passed bool, errorMessage string){
	input = strings.Trim(input," ")
	passed,_ = regexp.MatchString(`^.*[\W].*$`,input)
		if passed == false{
			errorMessage = "There should be a symbol in the input"
		}
	return
}
func RequireNumber(input string)(passed bool, errorMessage string){
	input = strings.Trim(input," ")
	passed,_ = regexp.MatchString(`^.*[\d].*$`,input)
		if passed == false{
			errorMessage = "There should be a numeric digit in the input"
		}
	return
}
func RequireStrongPassword (input string)(passed bool, errorMessage string){
	input = strings.Trim(input," ")
	passed, errorMessage = RequireUppercase(input)
	if passed==false{
		goto ReturnPart
	}
	passed, errorMessage = RequireSymbol(input)
	if passed==false{
		goto ReturnPart
	}
	passed, errorMessage = RequireNumber(input)
	if passed==false{
		goto ReturnPart
	}
	ReturnPart: return
}

func RequireFirstLetter(input string)(passed bool, errorMessage string){
	input = strings.Trim(input," ")
	passed,_ = regexp.MatchString(`^[A-Za-z]`,input)
	if passed == false{
		errorMessage = "There should be a letter on the first position"
	}
	return
}
func RequireMinCount(input string, count int)(passed bool, errorMessage string){
	input = strings.Trim(input," ")
	if len(input) >= count{
		passed = true
	} else{
		errorMessage ="minimum "+string(count)+" symbols are required."
	}
	return
}