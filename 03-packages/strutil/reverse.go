package strutil

// Reverse will return string
// --> exported function Reverse should have comment or be unexported
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
