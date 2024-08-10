package main

import "fmt"

func main(){

	numbers := [] int{}

	for count :=1; count <= 10; count++{
	
	numbers = append(numbers, count)

	}


	evennumbers := [] int {}
	oddnumbers := [] int {}

	for _, counter := range numbers{

	if counter % 2 == 0{
	evennumbers = append(evennumbers, counter)

	}else{
		
		oddnumbers = append(oddnumbers, counter)
		
	}
	
	}

	fmt.Println("your even numbers are", evennumbers)
	fmt.Println("your odd numbers are", oddnumbers)
	fmt.Println(" your numbers are", numbers)
	fmt.Println("your total numbers in your slice", len(numbers))
}
