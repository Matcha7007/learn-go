package main

import (
	"fmt"
	"maps"
)

func main() {
	// mapVariable map[keyType]valueType
	// mapVariable = make(map[keyType]valueType)

	// using a map literal
	// mapVariable = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)
	
	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println("key1:",myMap["key1"])
	fmt.Println("key:",myMap["key"])
	
	myMap["code"] = 21
	fmt.Println(myMap)
	
	delete(myMap, "key1")
	fmt.Println(myMap)
	
	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap)
	
	// clear map
	// clear(myMap)
	// fmt.Println(myMap)
	
	value, unknownValue := myMap["key1"]
	if unknownValue {
		fmt.Println("A value exists with key1")
	} else {
		fmt.Println("No value exists with key1")
	}
	fmt.Println("Value key1:", value)
	fmt.Println("Is a value associated with key1:", unknownValue)
	
	myMap2 := map[string]int{"a":1, "b":2}
	fmt.Println(myMap2)
	
	myMap3 := map[string]int{"a":1, "b":2}

	// compare map
	if maps.Equal(myMap3, myMap2){
		fmt.Println("myMap3 and myMap2 are equal")
	}

	// iterate map
	for k, v := range myMap3 {
		fmt.Println(k, v)
	}

	// initialize map to nil value
	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value.")
	} else {
		fmt.Println("The map is not initialized to nil value.")
	}

	val := myMap4["key"]
	fmt.Println(val)
	
	// add key and value
	// when we initialized map like myMap4, we can't add or update key and value direct to myMap4 like below
	// myMap4["key"] = "value" -> error nil deference in map update

	// we need use make() to add key and value to myMap4
	myMap4 = make(map[string]string)
	// after use make we can add key and value
	myMap4["key"] = "value" 
	fmt.Println(myMap4)

	// get length the map
	fmt.Println("myMap4 length is", len(myMap4))

	// multi dimensional maps or nested maps
	myMap5 := make(map[string]map[string]string)

	myMap5["map1"] = myMap4
	fmt.Println(myMap5)

}