package word

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"unicode"
)

// ToUpper 转换为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 转换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderScoreToUpperCamelCase 下划线转换为大写驼峰
func UnderScoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1) //先将_转换为空格
	s = cases.Title(language.English).String(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderScoreToLowerCamelCase 下划线转换为小写驼峰
func UnderScoreToLowerCamelCase(s string) string {
	s = UnderScoreToUpperCamelCase(s) //只需要处理首字母，其他与大写驼峰一致
	return string(unicode.ToLower(rune(s[0]))) + s[1:]

}

// CamelCaseToUnderScore 将驼峰转化为下划线
func CamelCaseToUnderScore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
			output = append(output, unicode.ToLower(r))
		}
		output = append(output, r)
	}
	return string(output)
}
