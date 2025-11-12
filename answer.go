// Section A: Easy (Fundamentals)
// -----------------------------------------------------------

// 1. Write a program to check if a number is positive, negative, or zero using if statements.

package main
import "fmt"
func main() {
    var num float64
    num = 8

    if num > 0 {
        fmt.Println("The number is positive.")
    } else if num < 0 {
        fmt.Println("The number is negative.")
    } else {
        fmt.Println("The number is zero.")
    }
}





// =========================================================

// 2. Write a program to determine whether a person is eligible to vote based on age.

// =========================================================

package main
import "fmt"
func main() {
    var age int
    age = 20

    if age >= 18 {
        fmt.Println("Eligible to vote.")
    } else {
        fmt.Println("Not eligible to vote.")
    }
}


// =========================================================


// 3. Write a program that checks whether a given number is even or odd.

// =========================================================

package main
import "fmt"
func main() {
    var num int
    num = 7

    if num%2 == 0 {
        fmt.Println("The number is even.")
    } else {
        fmt.Println("The number is odd.")
    }
}

// =========================================================


// 4. Write a program that takes a student’s marks and prints “Pass” or “Fail” using an if-else condition.

// =========================================================

package main
import "fmt"
func main() {
    var marks int
    marks = 75

    if marks >= 60 {
        fmt.Println("Pass")
    } else {
        fmt.Println("Fail")
    }
}



// =========================================================

// 5. Write a program to check whether a given year is a leap year or not.

package main
import "fmt"
func main() {
    var year int
    year = 2020

    if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
        fmt.Println("Leap Year")
    } else {
        fmt.Println("Not a Leap Year")
    }
}

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

package main
import "fmt"
func main() {
    var marks int
    marks = 85

    if marks >= 90 {
        fmt.Println("Grade A")
    } else if marks >= 75 {
        fmt.Println("Grade B")
    } else if marks >= 60 {
        fmt.Println("Grade C")
    } else {
        fmt.Println("Grade F")
    }
}



// =========================================================

// 2. Write a program that takes two numbers and prints which one is greater using if-else.

package main
import "fmt"
func main() {
    var num1, num2 int
    num1 = 10
    num2 = 20

    if num1 > num2 {
        fmt.Println("Number 1 is greater.")
    } else if num1 < num2 {
        fmt.Println("Number 2 is greater.")
    } else {
        fmt.Println("Both numbers are equal.")
    }
}

// =========================================================




// =========================================================

// 3. Write a program that checks if a character is a vowel or consonant using a switch statement.

// =========================================================

package main
import "fmt"
func main() {
    var char string
    char = "a"

    switch char {
    case "a", "e", "i", "o", "u":
        fmt.Println("Vowel")
    default:
        fmt.Println("Consonant")
    }
}


// =========================================================

// 4. Write a program that takes a day name (like Monday, Tuesday, etc.) and prints whether it’s a weekday or weekend using switch.

package main
import "fmt"
func main() {
    var day string
    day = "Saturday"

    switch day {
    case "Saturday", "Sunday":
        fmt.Println("Weekend")
    default:
        fmt.Println("Weekday")
    }
}

// =========================================================




// =========================================================

// 5. Write a program that uses a short if statement to check if a number is divisible by 3 and 5 at the same time.

// =========================================================

package main
import "fmt"
func main() {
    var num int
    num = 15

    if num%3 == 0 && num%5 == 0 {
        fmt.Println("The number is divisible by both 3 and 5.")
    } else {
        fmt.Println("The number is not divisible by both 3 and 5.")
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

import (
	"fmt"
)
func main() {
	// Predefined username and password
	const correctUsername = "admin"
	const correctPassword = "12345"

	var username, password string
    // Taking username input
    fmt.Print("Enter username: ")
    fmt.Scan(&username)
    // Taking password input
    fmt.Print("Enter password: ")
    fmt.Scan(&password)

    if username == correctUsername && password == correctPassword {
        fmt.Println("Login successful")
    } else if username == correctUsername && password != correctPassword {
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

package main

import "fmt"

func main() {
    var side1, side2, side3 float64

    fmt.Print("Enter the length of side 1: ")
    fmt.Scan(&side1)    
    fmt.Print("Enter the length of side 2: ")
    fmt.Scan(&side2)
    fmt.Print("Enter the length of side 3: ")
    fmt.Scan(&side3)

    if side1 == side2 && side2 == side3 {
        fmt.Println("The triangle is Equilateral.")
    } else if side1 == side2 || side2 == side3 || side1 == side3 {
        fmt.Println("The triangle is Isosceles.")
    } else {
        fmt.Println("The triangle is Scalene.")
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
import (
	"fmt"
	"internal/fmtsort"
)

func main() {
    var choice int16    
    fmt.Println("1. Start Game")
    fmt.Println("2. Load Game")
    fmt.Println("3. Exit")
    fmt.Print("Enter your choice: ")
    fmt.Scan(&choice)

    switch choice {
    case 1:
        fmtsort.StartGame()
    case 2:
        fmtsort.LoadGame()
    case 3:
        fmtsort.ExitGame()
    default:
        fmt.Println("Invalid option")
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
    var temperature int
    fmt.Print("Enter the current temperature: ")
    fmt.Scan(&temperature)

    switch {
    case temperature < 0:
        fmt.Println("Freezing")
    case temperature >= 0 && temperature <= 15:
        fmt.Println("Cold")
    case temperature >= 16 && temperature <= 30:
        fmt.Println("Warm")
    case temperature > 30:
        fmt.Println("Hot")
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

import (
	"fmt"
)

func main() {
    var totalMarks float64
    var passedMath, passedScience bool

    fmt.Print("Enter total marks (in percentage): ")
    fmt.Scan(&totalMarks)
    fmt.Print("Did the student pass Math? (true/false): ")
    fmt.Scan(&passedMath)
    fmt.Print("Did the student pass Science? (true/false): ")
    fmt.Scan(&passedScience)

    if totalMarks >= 60 {
        if passedMath && passedScience {
            fmt.Println("Eligible for admission")
        } else {
            if !passedMath && !passedScience {
                fmt.Println("Rejected: Did not pass Math and Science")
            } else if !passedMath {
                fmt.Println("Rejected: Did not pass Math")
            } else {
                fmt.Println("Rejected: Did not pass Science")
            }
        }
    } else {
        fmt.Println("Rejected: Insufficient marks")
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

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Print("Enter third number: ")
	fmt.Scan(&c)

	if a >= b && a >= c {
		fmt.Println("The largest number is:", a)
	} else if b >= a && b >= c {
		fmt.Println("The largest number is:", b)
	} else {
		fmt.Println("The largest number is:", c)
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

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// Take inputs
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	// Perform operation using switch
	switch operator {
	case "+":
		fmt.Println("Result:", num1+num2)
	case "-":
		fmt.Println("Result:", num1-num2)
	case "*":
		fmt.Println("Result:", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Println("Result:", num1/num2)
		} else {
			fmt.Println("Error: Division by zero is not allowed")
		}
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

package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Print("Enter a year: ")
	fmt.Scan(&year)

	if year%400 == 0 {
		fmt.Println(year, "is a century leap year.")
	} else if year%4 == 0 {
		fmt.Println(year, "is a normal leap year.")
	} else {
		fmt.Println(year, "is not a leap year.")
	}
}

// =========================================================
