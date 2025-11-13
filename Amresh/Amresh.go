// Assignment: Conditional Statements in Go (Golang)
// Topics Covered: if, else if, else, nested if, short if, switch, fallthrough, switch without expression

// -----------------------------------------------------------
// Section A: Easy (Fundamentals)
// -----------------------------------------------------------

// 1. Write a program to check if a number is positive, negative, or zero using if statements.

// =========================================================
// package main
// import "fmt"
// func main(){
// 	num := -2
// 	if num >= 0{
// 		fmt.Println("Positive")
// 	} else{
// 		fmt.Println("Negative")
// 	}
// }


// =========================================================

// 2. Write a program to determine whether a person is eligible to vote based on age.

// =========================================================


// package main
// import "fmt"
// func main(){
// 	age := 12
// 	if age  >= 18{
// 		fmt.Println("You are eligible for vote")
// 	} else{
// 		fmt.Println("You are not eligible")
// 	}
// }




// =========================================================


// 3. Write a program that checks whether a given number is even or odd.

// =========================================================

// package main
// import "fmt"
// func main(){
// 	num := 4
// 	if num % 2 == 0{
// 		fmt.Println("num is even")
// 	} else{
// 		fmt.Println("num is odd")
// 	}
// }



// =========================================================


// 4. Write a program that takes a student’s marks and prints “Pass” or “Fail” using an if-else condition.

// =========================================================

// package main
// import "fmt"
// func main(){
// 	marks := 100
// 	if marks > 30{
// 	 	fmt.Println("pass")
// 	} else{
// 		fmt.Println("fail")
// 	}
// }





// =========================================================

// 5. Write a program to check whether a given year is a leap year or not.

// =========================================================

// package main
// import "fmt"
// func main(){
// 	year := 2020
// 	if year % 4 == 0{
// 		fmt.Print("Leap year")
// 	} else{
// 		fmt.Print("not a leap yaer")
// 	}
// }


// =========================================================


// -----------------------------------------------------------
// Section B: Medium (Logic Building)
// -----------------------------------------------------------

// 1. Write a program to calculate grade based on marks using if-else if conditions.
//    90+ → Grade A
//    75–89 → Grade B
//    60–74 → Grade C
//    Below 60 → Grade F

// =========================================================

// package main
// import "fmt"
// func main(){
// 	marks := 65
// 	if marks > 90{
// 		fmt.Print("grade A")
// 	} else if marks >= 75{
// 		fmt.Print("grade B")
// 	} else if marks >= 60{
// 		fmt.Print("grade C")
// 	} else{
// 		fmt.Print("grade F")
// 	}
// }




// =========================================================

// 2. Write a program that takes two numbers and prints which one is greater using if-else.

// =========================================================
// 
// package main
// import "fmt"
// func main(){
// 	num1 := 12
// 	num2 := 11
// 	if num1 > num2{
// 		fmt.Print(num1)
// 	} else{
// 		fmt.Print(num2)
// 	}
// }




// =========================================================

// 3. Write a program that checks if a character is a vowel or consonant using a switch statement.

// =========================================================


// package main
// import "fmt"
// func main(){
// 	c := "f"
// 	switch c{
// 	case "a" , "e" , "i" , "o" , "u":
// 		fmt.Print("Vowel")
// 	default:
// 		fmt.Print("Consonant")
// 	}
// 
// }





// =========================================================

// 4. Write a program that takes a day name (like Monday, Tuesday, etc.) and prints whether it’s a weekday or weekend using switch.

// =========================================================

// package main
// import "fmt"
// 
// func main() {
// 	day := "Tuesday"
// 
// 	switch day {
// 	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
// 		fmt.Println("Weekdays")
// 	case "Saturday", "Sunday":
// 		fmt.Println("Weekends")
// 	default:
// 		fmt.Println("Invalid day")
// 	}
// }




// =========================================================

// 5. Write a program that uses a short if statement to check if a number is divisible by 3 and 5 at the same time.

// =========================================================

// package main
// import "fmt"
// func main(){
// 	num := 12
// 	if (num % 3 == 0) && (num % 5 == 0){
// 		fmt.Print("yes , divisble by both 3 and 5")
// 	} else{
// 		fmt.Print("No , its not divisble by 3 and 5")
// 	}
// }




// =========================================================


// // -----------------------------------------------------------
// Section C: Hard (Real-World Logic & Nesting)
// -----------------------------------------------------------

// Note: Use the following syntax to take input from the user in Go:
//
// var variableName datatype
// fmt.Print("Enter something: ")
// fmt.Scan(&variableName)
//
// Example:
// var age int
// fmt.Print("Enter your age: ")
// fmt.Scan(&age)


// =========================================================

// 1. Write a program to simulate a login system:
//    - Ask the user to enter a username and password.
//    - If both match predefined values, print “Login successful”.
//    - If username is correct but password is wrong, print “Invalid password”.
//    - Otherwise, print “User not found”.

// Hint: You can compare input strings using the == operator.


// =========================================================
// package main
// import "fmt"
// func main(){
// 	const username = "amresh"
// 	const password = 1234
// 	var curruser string;
// 	var pass int;
// 	fmt.Scan(&user)
// 	fmt.Scan(&pass)
// 	if username == curruser && password == pass{
// 		fmt.Print("Login Succesfull")
// 	} else if (username == user && password != pass){
// 		fmt.Print("Invalid password")
// 	} else{
// 		fmt.Print("User not found")
// 	}
// }



// =========================================================

// 2. Write a program that checks whether a given triangle is:
//    - Equilateral (all sides equal)
//    - Isosceles (two sides equal)
//    - Scalene (all sides different)
//
//    - Ask the user to enter the lengths of all three sides.
//    - Use nested if statements to determine the triangle type.

// =========================================================


// package main
// import "fmt"
// func main(){
// 	var a, b, c int
// 	fmt.Scan(&a,&b,&c)
// 	if a +b > c && b + c > a && c+a > b{
// 		if a == b && b == c && c == a{
// 			fmt.Print("Equilateral Triangle")
// 		} else{
// 			if a == b || b == c || c == a{
// 				fmt.Print("Isosceles Triangle")
// 			} else{
// 				fmt.Print("Scalene Triangle")
// 			}
// 		}		
// 	} else
// }



// =========================================================

// 3. Write a program that simulates a menu system using a switch statement:
//    - Display a simple menu:
//         1 → “Start Game”
//         2 → “Load Game”
//         3 → “Exit”
//    - Ask the user to enter their choice (1, 2, or 3).
//    - Use a switch statement to print the corresponding message.
//    - If input doesn’t match any option, print “Invalid option”.

// =========================================================

// package main
// import "fmt"
// func main(){
// 	fmt.Println("1 → Start Game")
// 	fmt.Println("2 → Load Game")
// 	fmt.Println("3 → Exit")
// 	var a int
// 	fmt.Scan(&a)
// 	switch a{
// 	case 1:
// 		fmt.Print("Start Game")
//     case 2:
// 		fmt.Print("Load Game")	
// 	case 3:
// 		fmt.Print("Exit")	
// 
// 	}
// }




// =========================================================

// 4. Write a program that uses a switch without an expression to classify temperature:
//    - Ask the user to input the current temperature (integer).
//    - Based on the value, print one of the following:
//         Below 0 → “Freezing”
//         0–15 → “Cold”
//         16–30 → “Warm”
//         Above 30 → “Hot”

// Hint: Use logical operators (>, <, <=, >=) inside switch cases.

// =========================================================
// package main
// import "fmt"
// func main(){
// 	var a int
// 	fmt.Scan(&a)
// 	switch {
// 	case a < 0:
// 		fmt.Println("Freezing")
// 	case a <=15:
// 		fmt.Println("Cold")	
//     case a <= 30:
// 		fmt.Println("Warm")
//     case a > 30:
// 		fmt.Println("Hot")		
// 	default:
// 		fmt.Println("Invalid")
// 	}
// }




// =========================================================

// 5. Write a program to check admission eligibility for a student based on nested conditions:
//    - Ask the user to input total marks (in percentage).
//    - Ask if the student passed Math and Science (true/false).
//    - Conditions:
//         Minimum marks: 60%
//         Must have passed both Math and Science
//    - If both conditions are true, print “Eligible for admission”.
//    - Otherwise, print the specific reason for rejection.

// =========================================================

// package main
// import "fmt"
// func main(){
// 	var marks int
// 	var pass bool
// 	fmt.Scan(&marks)
// 	fmt.Scan(&pass)
// 	if marks >= 60{
// 		if pass == true{
// 			fmt.Print("Eligible for admission")
// 		}else{
// 			fmt.Print("Not Eligible for admission")
// 		}
// 	}else{
// 		fmt.Println("Not good marks")
// 	}
// }




// =========================================================


// -----------------------------------------------------------
// Bonus (Challenge Zone)
// -----------------------------------------------------------

// 1. Write a program that takes three numbers as input and prints the largest using only conditional statements.

// Hint: Use nested if or multiple if-else conditions.


// =========================================================

// package main
// import "fmt"
// func main(){
// 	var a int
// 	var b int
// 	var c int
// 	fmt.Scan(&a,&b,&c)
// 	if a > b && a > c {
// 		fmt.Print(a)
// 	}else if b > c{
// 		fmt.Print(b)
// 	}else{
// 		fmt.Print(c)
// 	}
// }


// =========================================================

// 2. Create a small calculator using a switch statement:
//    - Ask the user to input two numbers and an operator (+, -, *, /).
//    - Perform the operation based on the chosen operator.
//    - Display the result.
//    - If the operator is invalid, print “Invalid operator”.

// =========================================================

// package main
// import "fmt"
// func main(){
// 	var fir int
// 	var sec int
// 	var ope string
// 	fmt.Scan(&fir,&sec,&ope)
// 	if ope == "+"{
// 		fmt.Println(fir+sec)
// 	}else if ope == "-"{
// 		fmt.Println(fir-sec)
// 
// 	}else if ope == "*"{
// 		fmt.Println(fir*sec)
// 	}else if ope == "/"{
// 		fmt.Println(fir/sec)
// 	}else{
// 		fmt.Println("invalid")
// 	}
// }


// =========================================================

// 3. Write a program that determines whether a given year is a century leap year or not.
//    - Ask the user to enter a year.
//    - A year is a century leap year if it is divisible by 400.
//    - A year is a normal leap year if it is divisible by 4 but not by 100.
//    - Otherwise, it is not a leap year.

// =========================================================

// package main
// import "fmt"
// func main(){
// 	var year int
// 	fmt.Scan(&year)
// 	if year%400 == 0 || (year%4 == 0 && year%100 != 0){
// 		fmt.Println("Leap year")
// 
// 	}else{
// 		fmt.Println("Not a Leap year")
// 	}
// }



// =========================================================

// 4. Write a program that uses fallthrough in a switch to demonstrate cumulative conditions:
//    - Example: Enter a number and print messages for all cases up to that number using fallthrough.
//    - For example, if number = 2, output should print case 1 and case 2 messages.

// =========================================================

// package main
// import "fmt"
// func main(){
// 	number := 2
// 	switch {
// 	case number >= 1:
// 		fmt.Println("one")
// 		fallthrough
// 	case number >= 2:
// 		fmt.Println("Two")
// 	default:
// 		fmt.Println("Something else")
// 	}
// }




// =========================================================

// 5. Write a program that checks if a student qualifies for a scholarship:
//    - Ask for the student’s marks (percentage).
//    - Ask for attendance percentage.
//    - Ask if the student has any backlogs (true/false).
//    - Conditions:
//         Must have 85% or higher marks
//         Attendance above 90%
//         No backlog
//    - If all conditions are true, print “Scholarship Approved”.
//    - Otherwise, print “Scholarship Denied”.

// =========================================================
// package main
// import "fmt"
// func main(){
// 	var a int
// 	var b int
// 	var c bool
// 	fmt.Scan(&a,&b,&c)
// 	if a >=85 && b >90 && c == false{
// 		fmt.Print("Scholarship Approved")
// 	}else{
// 		fmt.Print("Scholarship Denied")
// 	}
// }


// =========================================================

