package main

import(
	"fmt"
)
func main() {

	slc := []int{1,2,3,4,5}
	fmt.Println(slc)
	
	slc = append(slc,6)
	fmt.Println(slc)

	//copy---------------------------

	copy(slc[1:3],[]int{7,8})
	fmt.Println(slc)

	var brands = []string{"Apple", "Samsung", "LG", "Nokia"}

	fmt.Printf("the brands are: %q\n", brands)

	var clothBrands = make([]string,5,100)
	clothBrands2 := []string{"Adidas", "Nike", "New Balance", "Zara", "Puma"}//if i add "jordan" here it doesnt append to base slice
	copy(clothBrands,clothBrands2)
	clothBrands = append(clothBrands,"Jordan")
	fmt.Printf("the cloth brands are: %v\n. len: %d\n. cap:%d\n", clothBrands, len(clothBrands), cap(clothBrands))

	//remove element----------------------

	fmt.Println(slc)
	newSlc := append(slc[:3], slc[4:]...)
	
	fmt.Println(newSlc)

}