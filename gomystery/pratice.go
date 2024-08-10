package main

import "fmt"

func main(){

	numbers := [] int {}
	for count := 1; count <= 10; count++{

		numbers = append(numbers, count)
			if count % 2 ==0{
				fmt.Println( numbers)

		

			}else{
				fmt.Println(numbers)
}
		
}



		fmt.Print("your numbers are" , numbers, "  ")
		fmt.Println("the final count is", len(numbers))




}











