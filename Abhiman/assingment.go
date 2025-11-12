// Assignment: Conditional Statements in Go (Golang)
// Topics Covered: if, else if, else, nested if, short if, switch, fallthrough, switch without expression

// -----------------------------------------------------------
// Section A: Easy (Fundamentals)
// -----------------------------------------------------------

// 1. Write a program to check if a number is positive, negative, or zero using if statements.

// =========================================================

func main() {

	var a int
	fmt.Print("Enter a number: ")
	fmt.Scan(&a)

	if a > 0 {
		fmt.Println("Positive")
	} else if a == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Negative")
	}

}

// =========================================================

// 2. Write a program to determine whether a person is eligible to vote based on age.

// =========================================================

func main() {

	var a int
	fmt.Print("Enter Your Age: ")
	fmt.Scan(&a)

	if a > 18 {
		fmt.Println("Eligible")
	} else {
		fmt.Println("Not Eligible")
	}

}

// =========================================================

// 3. Write a program that checks whether a given number is even or odd.

// =========================================================

func main() {

	var a int
	fmt.Print("Enter a Number: ")
	fmt.Scan(&a)

	if a%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}

// =========================================================

// 4. Write a program that takes a student’s marks and prints “Pass” or “Fail” using an if-else condition.

// =========================================================

func main() {

	var a int
	fmt.Print("Enter your marks: ")
	fmt.Scan(&a)

	if a > 33 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

}

// =========================================================

// 5. Write a program to check whether a given year is a leap year or not.

// =========================================================

func main() {

	var a int
	fmt.Print("Enter a Year: ")
	fmt.Scan(&a)

	if a%4 == 0 && (a%100 != 0 || a%400 == 0) {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Not a Leap Year")
	}

}

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

func main() {

	var a int
	fmt.Print("Enter your marks: ")
	fmt.Scan(&a)

	if a >= 90 {
		fmt.Println("Grade A")
	} else if a >= 75 {
		fmt.Println("Grade B")
	} else if a >= 60 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade F")
	}

}

// =========================================================

// 2. Write a program that takes two numbers and prints which one is greater using if-else.

// =========================================================
func main() {

	a := 6
	b := 9

	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}

}

// =========================================================

// 3. Write a program that checks if a character is a vowel or consonant using a switch statement.

// =========================================================

func main() {

	day := "A"
	switch day {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		fmt.Println("Vowel")
	default:
		fmt.Println("consonant")
	}

}

// =========================================================

// 4. Write a program that takes a day name (like Monday, Tuesday, etc.) and prints whether it’s a weekday or weekend using switch.

// =========================================================
func main() {

	day := "Sunday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	}

}

// =========================================================

// 5. Write a program that uses a short if statement to check if a number is divisible by 3 and 5 at the same time.

// =========================================================

func main() {

	var a int
	fmt.Print("Enter a number: ")
	fmt.Scan(&a)

	if a%15 == 0 {
		fmt.Println("Divisible")
	} else {
		fmt.Println("Not Divisible")
	}

}

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

func main() {

	var username string
	fmt.Print("Enter your username: ")
	fmt.Scan(&username)

	var password string
	fmt.Print("Enter your password: ")
	fmt.Scan(&password)

	if username == "admin" && password == "password123" {
		fmt.Println("Login successful")
	} else if username == "admin" {
		fmt.Println("Invalid password")
	} else {
		fmt.Println("User not found")
	}

}

// =========================================================

// 2. Write a program that checks whether a given triangle is:
//    - Equilateral (all sides equal)
//    - Isosceles (two sides equal)
//    - Scalene (all sides different)
//
//    - Ask the user to enter the lengths of all three sides.
//    - Use nested if statements to determine the triangle type.

// =========================================================
func main() {

	var l1 string
	fmt.Print("Enter the length of the triangle: ")
	fmt.Scan(&l1)

	var b1 string
	fmt.Print("Enter the breadth of the triangle: ")
	fmt.Scan(&b1)

	var h1 string
	fmt.Print("Enter the height of the triangle: ")
	fmt.Scan(&h1)

	if l1 == b1 && b1 == h1 && l1 == h1 {
		fmt.Println("Equilateral")
	} else if l1 == b1 || b1 == h1 || l1 == h1 {
		fmt.Println("Isosceles")
	} else {
		fmt.Println("Scalene")
	}

}

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
func main() {

	var choice int
	fmt.Print("Enter the choice (1,2,3): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Start the Game")
	case 2:
		fmt.Println("Load the Game")
	case 3:
		fmt.Println("Exit")
	default:
		fmt.Println("Invalid option")
	}
}

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

func main() {

	var temp int
	fmt.Print("Enter the temp:")
	fmt.Scan(&temp)

	switch {
	case temp < 0:
		fmt.Println("Freezing")
	case temp < 15:
		fmt.Println("Cold")
	case temp < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}
}


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
package main

import "fmt"

func main() {

	var a int
	var b int
	var c int

	fmt.Print("Enter three numbers:")
	fmt.Scan(&a, &b, &c)

	if a > b {
		if a > c {
			fmt.Printf("%d", a)
		} else {
			fmt.Printf("%d", c)
		}
	} else if b > c {
		if b > c {
			fmt.Printf("%d", b)
		} else {
			fmt.Printf("%d", c)
		}
	} else {
		if c > a {
			fmt.Printf("%d", c)
		} else {
			fmt.Printf("%d", a)
		}
	}

}

// =========================================================

// 2. Create a small calculator using a switch statement:
//    - Ask the user to input two numbers and an operator (+, -, *, /).
//    - Perform the operation based on the chosen operator.
//    - Display the result.
//    - If the operator is invalid, print “Invalid operator”.

// =========================================================
func main() {

	var a int
	var b int
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Printf("%d", a+b)
	case "-":
		fmt.Printf("%d", a-b)
	case "*":
		fmt.Printf("%d", a*b)
	case "/":
		fmt.Printf("%d", a*b)
	default:
		fmt.Println("Invalid operator")
	}
}
// =========================================================

// 3. Write a program that determines whether a given year is a century leap year or not.
//    - Ask the user to enter a year.
//    - A year is a century leap year if it is divisible by 400.
//    - A year is a normal leap year if it is divisible by 4 but not by 100.
//    - Otherwise, it is not a leap year.

// =========================================================
func main() {

	var year int
	fmt.Print("Enter a year: ")
	fmt.Scanln(&year)

	if year%400 == 0 {
		fmt.Printf("%d is a century leap year.\n", year)
	} else if year%4 == 0 && year%100 != 0 {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}

// =========================================================

// 4. Write a program that uses fallthrough in a switch to demonstrate cumulative conditions:
//    - Example: Enter a number and print messages for all cases up to that number using fallthrough.
//    - For example, if number = 2, output should print case 1 and case 2 messages.

// =========================================================
var a int
	fmt.Println("Enter a number(1-5)")
	fmt.Scan(&a)
	switch a { 
	case 1:
		fmt.Println("Case 1")
		fallthrough
	case 2:
		fmt.Println("Case 2")
		fallthrough
	case 3:
		fmt.Println("Case 3")
		fallthrough
	case 4:
		fmt.Println("Case 4")
		fallthrough
	case 5:
		fmt.Println("Case 5")
	default:
		fmt.Println("Invalid Case")
	}
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
func main() {

	var marks int
		fmt.Print("Enter a marks: ")
		fmt.Scanln(&marks)
	
	var attendance int 
	fmt.Print("Enter attendance percentage: ")
	fmt.Scanln(&attendance)
  

	var backlog bool
	fmt.Print("Enter you have any backlogs (true/false): ")
	fmt.Scanln(&backlog)

	if marks>=85 && attendance>=75 && !backlog {
		fmt.Print("Scholarship Approved")
	}else {
		fmt.Print("Scholarship Denied")
	}

}
// =========================================================
