// 문자열 연산을 제공하는 패키지
package duckstring

import (
	"strings"
)
// 문자열더하기 a와 b 문자열 합치기
func AddStr(a string, b string) string {
	return a + b
}
// 문자열 빼기. a에서 b부분을 찾아 제거
func SubStr(a string, b string) string {
	return strings.Replace(a, b, "", -1)
}
// 문자열 곱하기. a문자열을 x번 반복 리턴
func MulStr(a string, x int) string {
	var result string
	for i := 0; i < x; i++ {
		result += a
	}
	return result
}
// 문자열 나누기. a문자열중 b문자열이 몇번 반복되는지 횟수리턴
func DivStr(a string, b string) int {
	return strings.Count(a, b)
}
// 문자열 나머지연산. a문자열중 b문자열을 제거한 나머지 리턴
func ReamindStr(a string, b string) string {
	count := strings.Count(a, b)
	return strings.Replace(a, b, "", count)
}
//문자열이 같은지 확인. int, float64타입도 포함 ( 그외에는 x)
func IsSame(a interface{}, b interface{}) bool {
	switch aType := a.(type) {
	case string:
		bType, ok := b.(string)
		return ok && aType == bType
	case int:
		bType, ok := b.(int)
		return ok && aType == bType
	case float64:
		bType, ok := b.(float64)
		const epsilon = 1e-9
		return ok && (aType-bType) < epsilon && (bType-aType) < epsilon
	default:
		return false
	}
}
// 문자열 비교. a와 b문자열중 더 큰 문자열 리턴. ex) aaa < aab
func IsBig(a string, b string) string {
	result := strings.Compare(a, b)
	if result > 0 {
		return a
	}
	return b
}
