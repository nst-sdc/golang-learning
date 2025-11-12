// Assignment: Conditional Statements in Go (Golang)
// Topics Covered: if, else if, else, nested if, short if, switch, fallthrough, switch without expression

// -----------------------------------------------------------
// Section A: Easy (Fundamentals)
// -----------------------------------------------------------

// 1. Write a program to check if a number is positive, negative, or zero using if statements.
// package main

// import "fmt"

// func main(){
// 	var a int
// 	fmt.Scan(&a)

// 	if a>0{
// 		fmt.Println("positive")
// 	}else if a<0{
// 		fmt.Println("negative")
// 	}else{
// 		fmt.Println("zero")
// 	}
// }

// =========================================================




// =========================================================

// 2. Write a program to determine whether a person is eligible to vote based on age.
// package main

// import "fmt"

// func main(){
// 	var a int
// 	fmt.Scan(&a)
// 	if a>18{
// 		fmt.Println("positive")
// 	}else if a<18{
// 		fmt.Println("negative")
// 	}
// }

// =========================================================




// =========================================================


// 3. Write a program that checks whether a given number is even or odd.
// package main

// import "fmt"

// func main(){
// 	var a int
// 	fmt.Scan(&a)
// 	if a%2==0{
// 		fmt.Println("even")
// 	}else{
// 		fmt.Println("odd")
// 	}
// }

// =========================================================




// =========================================================


// 4. Write a program that takes a student’s marks and prints “Pass” or “Fail” using an if-else condition.
// package main

// import "fmt"

// func main(){
// 	var a int
// 	fmt.Scan(&a)
// 	if a%2==0{
// 		fmt.Println("even")
// 	}else{
// 		fmt.Println("odd")
// 	}
// }


// =========================================================




// =========================================================

// 5. Write a program to check whether a given year is a leap year or not.
// package main

// import "fmt"

// func main(){
// 	var a int
// 	fmt.Scan(&a)
// 	if a%2==0{
// 		fmt.Println("even")
// 	}else{
// 		fmt.Println("odd")
// 	}
// }

// =========================================================




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
// 	var a int
// 	fmt.Scan(&a)
// 	if a>=90{
//             fmt.Println("Grade A")
//     }else if a>=80{
//             fmt.Println("Grade B")
//     }else if a>=70{
//             fmt.Println("Grade C")
//     }else if a>=60{
//             fmt.Println("Grade D")
//     }else{
//             fmt.Println("Grade F")
//     }


// }





// =========================================================

// 2. Write a program that takes two numbers and prints which one is greater using if-else.

// package main

// import "fmt"

// func main(){
// 	var a,b int
// 	fmt.Scan(&a,&b)

// 	if a>b{
// 		fmt.Println(a)
// 	}else{
// 		fmt.Println(b)
// 	}
// }
// =========================================================




// =========================================================

// 3. Write a program that checks if a character is a vowel or consonant using a switch statement.

// =========================================================




// =========================================================

// 4. Write a program that takes a day name (like Monday, Tuesday, etc.) and prints whether it’s a weekday or weekend using switch.

// package main
// import "fmt"

// func main(){
// 	var a string
// 	fmt.Scan(&a)
// 	switch a{
// 	case "monday":
// 		fmt.Println("weekday")
// 	case "tuesday":
// 		fmt.Println("weekday")
// 	case "Wednesday":
// 		fmt.Println("weekday")
// 	case "Thursday":
// 		fmt.Println("weekday")
// 	case "Friday":
// 		fmt.Println("weekday")
// 	case "Saturday":
// 		fmt.Println("weekend")
// 	case "Sunday":
// 		fmt.Println("weekend")
// 	}
// }

// =========================================================




// =========================================================

// 5. Write a program that uses a short if statement to check if a number is divisible by 3 and 5 at the same time.

// package main
// import "fmt"

// func main(){
// 	var a int
// 	fmt.Scan(&a)

// 	if a%3==0 && a%5==0{
// 		fmt.Println("yes")
// 	}else{
// 		fmt.Println("no")
// 	}

// }


// =========================================================




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

// package main
// import "fmt"

// func main(){
// 	 email:="abc@gmail.com"
// 	 pass:=123
//   var a string
// 	var b int
// 	fmt.Scan(&a,&b)

// 	if a==email && b==pass{
// 		fmt.Println("login successful")
// 	}else if a==email && b!=pass{
// 		fmt.Println("Invalid password")
// 	}else{
// 		fmt.Println("User not found")
// 	}
// }

// =========================================================




// =========================================================

// 2. Write a program that checks whether a given triangle is:
//    - Equilateral (all sides equal)
//    - Isosceles (two sides equal)
//    - Scalene (all sides different)
//
//    - Ask the user to enter the lengths of all three sides.
//    - Use nested if statements to determine the triangle type.

// package main
// import "fmt"

// func main(){
// 	var a,b,c int
// 	fmt.Scan(&a,&b,&c)
//   if a==b==c{
// 		fmt.Println("Equilateral")
// 	}else{
// 		if a==b or b==c or c==a{
// 			fmt.Println("isocellss")
// 		}else{
// 			fmt.Println("Scalene")
// 		}
// 	}

// }


// =========================================================




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




// =========================================================


// -----------------------------------------------------------
// Bonus (Challenge Zone)
// -----------------------------------------------------------

// 1. Write a program that takes three numbers as input and prints the largest using only conditional statements.

// Hint: Use nested if or multiple if-else conditions.


// =========================================================




// =========================================================

// 2. Create a small calculator using a switch statement:
//    - Ask the user to input two numbers and an operator (+, -, *, /).
//    - Perform the operation based on the chosen operator.
//    - Display the result.
//    - If the operator is invalid, print “Invalid operator”.

// =========================================================




// =========================================================

// 3. Write a program that determines whether a given year is a century leap year or not.
//    - Ask the user to enter a year.
//    - A year is a century leap year if it is divisible by 400.
//    - A year is a normal leap year if it is divisible by 4 but not by 100.
//    - Otherwise, it is not a leap year.

// =========================================================




// =========================================================

// 4. Write a program that uses fallthrough in a switch to demonstrate cumulative conditions:
//    - Example: Enter a number and print messages for all cases up to that number using fallthrough.
//    - For example, if number = 2, output should print case 1 and case 2 messages.

// =========================================================




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




// =========================================================