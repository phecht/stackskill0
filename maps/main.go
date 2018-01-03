package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
func map1() {

	fmt.Println("--------- map1 ----------")
	defer fmt.Println("----- end of map1 ------")
	aMap := map[string]string{
		"ser":   "to be (perm)",
		"tomar": "to take/drink",
		"creer": "to believe",
	}
	fmt.Println(aMap)
	delete(aMap, "tomar")
	fmt.Println(aMap)
	if val, exists := aMap["ser"]; exists {
		delete(aMap, "ser")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

}

func map2() {
	fmt.Println("--------- map2 ----------")
	defer fmt.Println("----- end of map2 ------")
	verbs := map[string]map[string]string{
		"ser": map[string]string{
			"s1": "soy",
			"s2": "eres",
			"s3": "es",
			"p1": "somos",
			"p2": "sois",
			"p3": "son",
		},
		"ir": map[string]string{
			"s1": "voy",
			"s2": "vas",
			"s3": "va",
			"p1": "vamos",
			"p2": "vais",
			"p3": "van",
		},
	}
	for key, value := range verbs {
		fmt.Printf("verb:%v\n", key)
		for vKey, vValue := range value {
			fmt.Printf("Tense:%v --> %v\n", vKey, vValue)
		}
	}
	//fmt.Println(verbs)

}
func main() {
	fmt.Println("maps:")
	map1()
	map2()

}
