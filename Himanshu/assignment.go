// Assignment: Conditional Statements in Go (Golang)
// Topics Covered: if, else if, else, nested if, short if, switch, fallthrough, switch without expression

// -----------------------------------------------------------
// Section A: Easy (Fundamentals)
// -----------------------------------------------------------

// 1. Write a program to check if a number is positive, negative, or zero using if statements.

// =========================================================
package main

import "fmt"

func main() {

	var number int;
	fmt.Scan(&number)

	if number >0 {
		fmt.Println("Positive number")
	} else if number ==0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Negative number")
	}
}



// =========================================================

// 2. Write a program to determine whether a person is eligible to vote based on age.

// =========================================================




// DOUBT: How i can run both function one by one?

func second() {

	var age int;
	fmt.Scan(&age)

	if age >=18 {
		fmt.Println("Eligible for vote")

	} else {
		fmt.Println("Not eligible")
	}
}



// =========================================================


// 3. Write a program that checks whether a given number is even or odd.

// =========================================================

func third() {

	var n int;
	fmt.Scan(&n)

	if n%2==0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}


// =========================================================


// 4. Write a program that takes a student’s marks and prints “Pass” or “Fail” using an if-else condition.

// =========================================================


func forth() {
	var marks int;

	fmt.Scan(&marks)

	if marks >= 40 {
		fmt.Println("Pass")

	} else {
		fmt.Println("Fail")
	}
}

// =========================================================

// 5. Write a program to check whether a given year is a leap year or not.

// =========================================================

func fifth() {
	var year int;

	fmt.Scan(&year)

	if year % 100==0 {
		fmt.Println("Not a leap")

	} else if year % 4 == 0{
		fmt.Println("leap")

	} else{
		fmt.Println("Not a leap")
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

func bone() {
	var marks int;

	fmt.Scan(&marks)

	if marks>= 90 {
		fmt.Println("Grade A")

	} else if marks>= 75 {
		fmt.Println("Grade B")
		
	} else if marks>= 60 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade F")

	}
}


// =========================================================

// 2. Write a program that takes two numbers and prints which one is greater using if-else.

// =========================================================

func btwo() {
	var a int;
	var b int;

	fmt.Scan(&a)
	fmt.Scan(&b)

	if a>b {
		fmt.Println(a, "is greater")

	} else if a == b {
		fmt.Println("Equal")
	} else {
		fmt.Println(b, " is greater")
	}

}


// =========================================================

// 3. Write a program that checks if a character is a vowel or consonant using a switch statement.

// =========================================================

func bthree() {
	var a string
	fmt.Scan(&a)

	switch a {
	case "a", "e", "i", "o", "u":
		fmt.Println(a,"is vowel.")
	default:
		fmt.Println(a,"is consonant.")
	}
}


// =========================================================

// 4. Write a program that takes a day name (like Monday, Tuesday, etc.) and prints whether it’s a weekday or weekend using switch.

// =========================================================

func bfour() {
	var day string
	fmt.Scan(&day)

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}
}


// =========================================================

// 5. Write a program that uses a short if statement to check if a number is divisible by 3 and 5 at the same time.

// =========================================================


func bfive() {
	var num int;
	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println(num, "is divisible by both 3 & 5.")
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
//    - If username is trial but password is wrong, print “Invalid password”.
//    - Otherwise, print “User not found”.

// Hint: You can compare input strings using the == operator.


// =========================================================

func cone() {
	var username, password string;
	const trialUsername = "himanshu"
	const trialPassword = "abc123"

	fmt.Scan(&username)
	fmt.Scan(&password)

	if username == trialUsername && password == trialPassword {
		fmt.Println("Login successful")
	} else if username == trialUsername && password != trialPassword {
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

func ctwo() {
	var a int;
	var b int;
	var c int;

	fmt.Print("Enter side a: ")
	fmt.Scan(&a)
	fmt.Print("Enter side b: ")
	fmt.Scan(&b)
	fmt.Print("Enter side c: ")
	fmt.Scan(&c)

	if a == b {
		if b == c {
			fmt.Println("Equilateral.")
		} else {
			fmt.Println("Isosceles.")
		}
	} else if b == c || a == c {
		fmt.Println("Isosceles.")
	} else {
		fmt.Println("Scalene.")
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

func cthree() {
	var choice int;

	fmt.Println("Main Menu")
	fmt.Println("1 → Start Game")
	fmt.Println("2 → Load Game")
	fmt.Println("3 → Exit")

	fmt.Print("Enter your choice (1, 2, or 3): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Starting game...")
	case 2:
		fmt.Println("Loading game...")
	case 3:
		fmt.Println("Exiting... Goodbye!")
	default:
		fmt.Println("Invalid option.")
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



func cfour() {
    var temp int

    fmt.Print("Enter the current temperature: ")
    fmt.Scan(&temp)

    switch {
    case temp < 0:
        fmt.Println("Freezing")
    case temp >= 0 && temp <= 15:
        fmt.Println("Cold")
    case temp >= 16 && temp <= 30:
        fmt.Println("Warm")
    case temp > 30:
        fmt.Println("Hot")
    default:
        fmt.Println("Not a valid input")
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



func cfive() {
	var tmarks float64
	var ispassedMath, ispassedScience bool

	fmt.Print("Enter total marks (in percentage): ")
	fmt.Scan(&tmarks)

	fmt.Print("Passed Math? (true/false): ")
	fmt.Scan(&ispassedMath)

	fmt.Print("Passed Science? (true/false): ")
	fmt.Scan(&ispassedScience)

	if tmarks >= 60 {
		if ispassedMath {
			if ispassedScience {
				fmt.Println("eligible for admission")
			} else {
				fmt.Println("not eligible failed in science")
			}
		} else {
			fmt.Println("not eligible failed in math")
		}
	} else {
		fmt.Println("not eligible: marks less than 60%")
	}
}


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
