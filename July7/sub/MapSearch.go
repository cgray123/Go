package sub

import "fmt"

func MapSearch() {
	m := make(map[string]string)
	m["grape"] = "fruit"
	m["apple"] = "fruit"
	m["pear"] = "fruit"
	m["blueberry"] = "fruit"
	m["rasberry"] = "fruit"
	m["onion"] = "vegetable"
	m["garlic"] = "vegetable"
	m["greenbean"] = "vegetable"
	m["potato"] = "vegetable"
	m["kale"] = "vegetable"
retry:
	fmt.Println("input 'fruit'or 'vegetable' or a specific fruit or vegetable")
	var input string
	output := make(map[string]string)
	fmt.Scanln(&input)
	for k, v := range m {
		if k == input {
			output[k] = v
			break
		} else if v == input {
			output[k] = v
		}
	}
	if len(output) == 0 {
		fmt.Println("input not found, would you like to try again? (input 'y' to try again, input anything else to exit)")
		var input string
		fmt.Scanln(&input)
		if input == "y" {
			goto retry
		}
	} else {
		fmt.Println(output)
	}
}
