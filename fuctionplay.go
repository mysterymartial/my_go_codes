package main

import "fmt"


	func add_numbers( number1 , number2 int) int{

		total := number1 + number2


		return total


}

	func power( base_number, exponent_number int) int{
		power := 1
		for count := 1; count <= exponent_number; count++{
			power *= base_number
}
			
	return power

}

	func length_char(  word string) int{
		return len(word)
}

	func reversal(number []int) [] int  {
		for counter, count:= 0, len(number)-1; counter < count; counter, count = counter + 1, count - 1{
			number[counter], number[count] = number[count], number[counter]
}

			return number

		}
			
	

	func main(){
		 number1  := 2
		 number2  := 3
		 name := "bigboy"

	result := add_numbers(number1,number2)
	square := power(number1, number2)
	count := length_char(name)
	numbers := [] int { 7, 8,1,5,9}
	 reverse := reversal(numbers)
	fmt.Println(result)
	fmt.Println(square)
	fmt.Println(count)
	fmt.Println(reverse)
	
}

		