package main

func NewStrings() []string {
	result := []string{}
	for _, v := range []string{"a", "b", "c"} {
		result = append(result, v)
	}
	return result
}
func main() {
	result := []string{}
	for _, v := range []string{"a", "b", "c"} {
		result = append(result, v)
	}
	println("main2")
}
