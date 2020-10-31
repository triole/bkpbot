package rx

import "regexp"

// Lib a struct, you know...
type Lib struct {
	AnsiSequence      string
	ControlCharacters string
	AfterLastSlash    string
	UpToLastSlash     string
}

// InitLib returns a set of regular expressions
func InitLib() (lib Lib) {
	lib = Lib{
		AnsiSequence:      `\\x1b[^m]*m`,
		ControlCharacters: `(\\\\t|\\\\r|\\n)`,
		AfterLastSlash:    `[^/]+$`,
		UpToLastSlash:     `^(.*[\\\\/])`,
	}
	return
}

// FilterArrayByRegex filters an array and returns only the values that match the regex
func FilterArrayByRegex(arr []string, rx string) []string {
	returnArr := []string{}
	r, _ := regexp.Compile(rx)
	for _, str := range arr {
		if r.Match([]byte(str)) == true {
			returnArr = append(returnArr, str)
		}
	}
	return returnArr
}

// Find returns the substring that matches the given regex scheme
func Find(rx string, content string) (r string) {
	temp, _ := regexp.Compile(rx)
	r = temp.FindString(content)
	return
}

// Match returns true or false depending if the regex matches
func Match(rx string, str string) (b bool) {
	re, _ := regexp.Compile(rx)
	b = re.MatchString(str)
	return
}

// Sub inspired by Pythons re.sub function, string substitition
func Sub(rx string, str string, rep string) (r string) {
	re := regexp.MustCompile(rx)
	r = re.ReplaceAllString(str, rep)
	return
}
