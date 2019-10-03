package main

import "fmt"

func main() {
	/*
		array statement

		var a [5]int
		a[1] = 100
		fmt.Println(a)


		var x [5]float64
		x[0] = 98
		x[1] = 93
		x[2] = 77
		x[3] = 82
		x[4] = 83

		var total float64 = 0
		for i := 0; i < len(x); i++ {
			total += x[i]
		}
		fmt.Println(total / float64(len(x)))

	*/

	//	Slices
	/*
		// xample slices
		arr := [5]float64{1,2,3,4,5}
		x := arr[3:5]
		fmt.Println(x)
	*/

	/*
		//Slice Append
		slice1 := []int{1,2,3}
		slice2 := append(slice1, 4,5)
		fmt.Println(slice1,slice2)
	*/

	/*
		//	Maps
		x := make(map[int]int)
		x[1] = 10
		fmt.Println(x[1]);
	*/

	/*
		// example new
		//Get Index array maps
		elements := make(map[string]string)
		elements["H"] = "Hydrogen"
		elements["HE"] = "Helium"
		elements["Li"] = "Lithium"
		elements["Be"] = "Beryllium"
		elements["B"] = "Boron"
		elements["C"] = "Carbon"
		elements["N"] = "Nitrogen"
		elements["O"] = "Oxygen"
		elements["F"] = "Fluorine"
		elements["Ne"] = "Neon"

		if name, ok := elements["No"]; ok {
			fmt.Println(name, ok)
		}

	*/

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

}
