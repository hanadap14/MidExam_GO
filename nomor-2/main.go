package main

func main() {
	var array1 = [3]string{"a", "c", "d"}
	var array2 = [3]string{"a", "d", "c"}

	var i int

	println("array 1 : ")
	for i = 0; i < len(array1); i++ {
		print(array1[i], " ")
	}

	println("\narray 2 : ")
	for i = 0; i < len(array2); i++ {
		print(array2[i], " ")
	}

	print("\n")

	for i = 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			println("index ke", i, "berbeda")
		}
	}

}
