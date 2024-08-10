package main

import "fmt"

func main(){

	var number int 
	var reverse int 
	
	var remainder int 
	fmt.Println("Enter a set of numbers")
	fmt.Scanln(&number)
	 orgnumber  := number

	for number !=0{
		remainder = number % 10
		reverse = reverse * 10 + remainder 
		number = number / 10
	}
		
	if orgnumber == reverse{

		fmt.Println("Your number  is a  palindrome \n")
		

	}else{
		fmt.Println("your number is not palindrome \n")




}
	
		
}
		
	