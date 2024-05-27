package main

import "fmt"

func main() {
	map1 := map[string]bool{"Interview": true, "Bit": true}
	map2 := map[string]bool{"Interview": true, "Questions": true}
	map3 := map1
	map1 = map2 //copy description
	fmt.Println(map1, map2, map3)
}
