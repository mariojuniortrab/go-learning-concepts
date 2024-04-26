package main

import "fmt"

func main() {
	i := 2.4
	fmt.Println(getType(i))
}

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "unknown"
	}
}
