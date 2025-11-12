// Assignment: Conditional Statements in Go (Golang)
// Topics Covered: if, else if, else, nested if, short if, switch, fallthrough, switch without expression
// -----------------------------------------------------------
// Section A: Easy (Fundamentals)
// -----------------------------------------------------------

// 1. Write a program to check if a number is positive, negative, or zero using if statements.

// =========================================================
package main
import "fmt"
var x int
func main() {
	if x > 0 {
		fmt.Println("Positive")
	} else if x < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
}
// =========================================================

// 2. Write a program to determine whether a person is eligible to vote based on age.

// =========================================================
package main
import "fmt"
var y int
func main() {
	if y >= 18 {
		fmt.Println("Eligible")
	} else {
		fmt.Println("Not Eligible")
	}
}
// =========================================================


// 3. Write a program that checks whether a given number is even or odd.

// =========================================================
package main
import "fmt"
var z int
func main() {
	if z % 2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
}
// =========================================================


// 4. Write a program that takes a student’s marks and prints “Pass” or “Fail” using an if-else condition.

// =========================================================
package main
import "fmt"
var a int
func main() {
	if a >= 40 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}
// =========================================================

// 5. Write a program to check whether a given year is a leap year or not.

// =========================================================
package main
import "fmt"
var b int
func main() {
	if (b % 4 == 0 && b % 100 != 0) || (b % 400 == 0)  {
		fmt.Println("Eligible")
	} else {
		fmt.Println("Not Eligible")
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
package main
import "fmt"
var c int
func main() {
	if y >= 90 {
		fmt.Println("Grade A")
	} else if y >= 75 {
		fmt.Println("Grade B")
	} else if y >= 60 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade F")
	}
}
// =========================================================

// 2. Write a program that takes two numbers and prints which one is greater using if-else.

// =========================================================
package main
import "fmt"
var d int
var e int
func main() {
	if d > e {
		fmt.Println(d)
	} else {
		fmt.Println(e)
	}
}
// =========================================================

// 3. Write a program that checks if a character is a vowel or consonant using a switch statement.

// =========================================================
package main
import "fmt"
var f string 
func main() {
	switch f {
	case "A":
		fmt.Println("Vowel")
	case "E":
		fmt.Println("Vowel")
	case "I":
		fmt.Println("Vowel")
	case "O":
		fmt.Println("Vowel")
	case "U":
		fmt.Println("Vowel")
	case "a":
		fmt.Println("Vowel")
	case "e":
		fmt.Println("Vowel")
	case "i":
		fmt.Println("Vowel")
	case "o":
		fmt.Println("Vowel")
	case "u":
		fmt.Println("Vowel")
	default:
		fmt.Println("Constant")
	}
}
// =========================================================

// 4. Write a program that takes a day name (like Monday, Tuesday, etc.) and prints whether it’s a weekday or weekend using switch.

// =========================================================
package main
import "fmt"
var g string
func main() {
	switch g {
	case "Monday":
		fmt.Println("Weekday")
	case "Tuesday":
		fmt.Println("Weekday")
	case "Wednesday":
		fmt.Println("Weekday")
	case "Thursday":
		fmt.Println("Weekday")
	case "Friday":
		fmt.Println("Weekday")
	case "Saturday":
		fmt.Println("Weekend")
	case "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day mentioned")
	}
}
// =========================================================

// 5. Write a program that uses a short if statement to check if a number is divisible by 3 and 5 at the same time.

// =========================================================
package main
import "fmt"
var h int
func main() {
	if h % 3 == 0 && h % 5 == 0 {
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
package main
import "fmt"
func main() {
	fmt.Println("Enter Username and Password")
	var a string
	var b string
	user := "kartik"
	pass := "kartik@123"
	fmt.Scan(&a)
	fmt.Scan(&b)
	if user == a && pass == b {
		fmt.Println("Login Successful")
	} else if user == a && pass != b {
		fmt.Println("Invalid Password")
	} else {
		fmt.Println("User not Found")
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
package main
import "fmt"
func main() {
	fmt.Println("Enter lengths:")
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	if a == b && b == c {
		fmt.Println("Equilateral")
	} else if a != b && b != c && a != c {
		fmt.Println("Scalene")
	} else {
		fmt.Println("Isosceles")
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
package main
import "fmt"
func main() {
    var c int
    fmt.Scan(&c)
    switch c {
        case 1:
        fmt.Println("Start Game")
        case 2:
        fmt.Println("Load Game")
        case 3:
        fmt.Println("Exit")
        default:
        fmt.Println("Invalid Option")
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
package main
import "fmt"
func main() {
	var t int
	fmt.Scan(&t)
	switch {
	case t < 0:
		fmt.Println("Freezing")
	case 15 >= t:
		fmt.Println("Cold")
	case 30 >= t:
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
package main
import "fmt"
func main() {
	var Marks int
	var Math bool
	var Science bool
	fmt.Println("Enter Percentage:")
	fmt.Scan(&Marks)
	fmt.Println("Math: Pass or Fail")
	fmt.Scan(&Math)
	fmt.Println("Science: Pass or Fail")
	fmt.Scan(&Science)
	if Marks >= 60 {
		if Math == true && Science == true {
			fmt.Println("Eligible")
		} else if Math == false && Science == true {
			fmt.Println("Ineligible due to Math")
		} else if Math == true && Science == false {
			fmt.Println("Ineligible due to Science")
		} else {
			fmt.Println("Ineligible due to Math and Science")
		} 
	} else {
		fmt.Println("Ineligible due to Percentage")
}
}
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
	var a,b,c int
	fmt.Scan(&a,&b,&c)
	if a > b && a > c {
		fmt.Println(a)
	} else if b > a && b > c {
		fmt.Println(b)
	} else {
		fmt.Println(c)
	}
}
// =========================================================

// 2. Create a small calculator using a switch statement:
//    - Ask the user to input two numbers and an operator (+, -, *, /).
//    - Perform the operation based on the chosen operator.
//    - Display the result.
//    - If the operator is invalid, print “Invalid operator”.

// =========================================================
package main
import "fmt"
func main() {
	var x string
	var a,b int
	fmt.Scan(&x)
	fmt.Scan(&a,&b)
	if x == "+" {
		fmt.Println(a+b)
	} else if x == "-" {
		fmt.Println(a-b)
	} else if x == "*" {
		fmt.Println(a*b)
	} else {
		fmt.Println(a/b)
	}
}
// =========================================================

// 3. Write a program that determines whether a given year is a century leap year or not.
//    - Ask the user to enter a year.
//    - A year is a century leap year if it is divisible by 400.
//    - A year is a normal leap year if it is divisible by 4 but not by 100.
//    - Otherwise, it is not a leap year.

// =========================================================
package main
import "fmt"
func main() {
var x int
	fmt.Println("Enter a Year:")
	fmt.Scan(&x)
	if x % 400 == 0 {
		fmt.Println("Century Leap Year")
	} else if x % 4 == 0 && x % 100 != 0 {
			fmt.Println("Leap Year")
	} else {
		fmt.Println("Not a Leap Year")
	}
}
// =========================================================

// 4. Write a program that uses fallthrough in a switch to demonstrate cumulative conditions:
//    - Example: Enter a number and print messages for all cases up to that number using fallthrough.
//    - For example, if number = 2, output should print case 1 and case 2 messages.

// =========================================================
package main
import "fmt"
func main() {
	var x int
	fmt.Scan(&x)
	switch x {
	case 1:
		fmt.Println("case 1")
		fallthrough
	case 2:
		fmt.Println("case 2")
		fallthrough
	case 3:
		fmt.Println("case 3")
		fallthrough
	case 4:
		fmt.Println("case 4")
		fallthrough
	case 5:
		fmt.Println("case 5")
	default:
		fmt.Println("out of range")
	}
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
package main
import "fmt"
func main() {
	fmt.Println("Enter Marks")
	fmt.Println("Enter Attendence")
	fmt.Println("Enter Backlog Status")
	var x int
	var a int
	var t bool
	fmt.Scan(&x)
	fmt.Scan(&a)
	fmt.Scan(&t)
	if x >= 85 && a >= 90 && t == false {
		fmt.Println("Scholarship Approved")
	} else {
		fmt.Println("Scholarship Denied")
	}
}
// =========================================================
