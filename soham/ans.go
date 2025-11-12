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
    var num float64
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num > 0 {
        fmt.Println("Number is positive")
    } else if num < 0 {
        fmt.Println("Number is negative")
    } else {
        fmt.Println("Number is zero")
    }
}



// =========================================================

// 2. Write a program to determine whether a person is eligible to vote based on age.

// =========================================================
package main
import "fmt"

func main() {
    var age int
    fmt.Print("Enter your age: ")
    fmt.Scan(&age)

    if age >= 18 {
        fmt.Println("You can vote!")
    } else {
        fmt.Println("Sorry, you are not eligible to vote.")
    }
}



// =========================================================


// 3. Write a program that checks whether a given number is even or odd.

// =========================================================

// 3. Check even or odd number
package main
import "fmt"

func main() {
    var n int
    fmt.Print("Enter a number: ")
    fmt.Scan(&n)

    if n%2 == 0 {
        fmt.Println("Even number")
    } else {
        fmt.Println("Odd number")
    }
}



// =========================================================


// 4. Write a program that takes a student’s marks and prints “Pass” or “Fail” using an if-else condition.

// =========================================================

package main
import "fmt"

func main() {
    var marks int
    fmt.Print("Enter your marks: ")
    fmt.Scan(&marks)

    if marks >= 60 {
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

func main() {
    var year int
    fmt.Print("Enter year: ")
    fmt.Scan(&year)

    if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
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

package main
import "fmt"

func main() {
    var year int
    fmt.Print("Enter year: ")
    fmt.Scan(&year)

    if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
        fmt.Println("Leap Year")
    } else {
        fmt.Println("Not a Leap Year")
    }
}


// =========================================================

// 2. Write a program that takes two numbers and prints which one is greater using if-else.

// =========================================================

package main
import "fmt"

func main() {
    var a, b int
    fmt.Print("Enter two numbers: ")
    fmt.Scan(&a, &b)

    if a > b {
        fmt.Println(a, "is greater")
    } else if b > a {
        fmt.Println(b, "is greater")
    } else {
        fmt.Println("Both are equal")
    }
}


// =========================================================

// 3. Write a program that checks if a character is a vowel or consonant using a switch statement.

// =========================================================

package main
import "fmt"

func main() {
    var ch string
    fmt.Print("Enter a letter: ")
    fmt.Scan(&ch)

    switch ch {
    case "a", "e", "i", "o", "u":
        fmt.Println("Vowel")
    default:
        fmt.Println("Consonant")
    }
}


// =========================================================

// 4. Write a program that takes a day name (like Monday, Tuesday, etc.) and prints whether it’s a weekday or weekend using switch.

// =========================================================

package main
import "fmt"

func main() {
    var day string
    fmt.Print("Enter day name: ")
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

package main
import "fmt"

func main() {
    var num int
    fmt.Print("Enter number: ")
    fmt.Scan(&num)

    if num%3 == 0 && num%5 == 0 {
        fmt.Println("Divisible by 3 and 5")
    } else {
        fmt.Println("Not divisible by both")
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
    const username = "admin"
    const password = "1234"

    var user, pass string

    fmt.Print("Enter username: ")
    fmt.Scan(&user)
    fmt.Print("Enter password: ")
    fmt.Scan(&pass)

    if user == username && pass == password {
        fmt.Println("Login Successful!")
    } else if user == username {
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
    var s1, s2, s3 float64

    fmt.Print("Enter 3 sides: ")
    fmt.Scan(&s1, &s2, &s3)

    if s1 == s2 && s2 == s3 {
        fmt.Println("Equilateral triangle")
    } else if s1 == s2 || s2 == s3 || s1 == s3 {
        fmt.Println("Isosceles triangle")
    } else {
        fmt.Println("Scalene triangle")
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
    var choice int
    fmt.Println("1. Start Game")
    fmt.Println("2. Load Game")
    fmt.Println("3. Exit")
    fmt.Print("Enter choice: ")
    fmt.Scan(&choice)

    switch choice {
    case 1:
        fmt.Println("Game Started!")
    case 2:
        fmt.Println("Game Loaded!")
    case 3:
        fmt.Println("Goodbye!")
    default:
        fmt.Println("Invalid choice")
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
    var temp int
    fmt.Print("Enter temperature: ")
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
    var marks float64
    var math, science bool

    fmt.Print("Enter total marks: ")
    fmt.Scan(&marks)
    fmt.Print("Passed in Math (true/false)? ")
    fmt.Scan(&math)
    fmt.Print("Passed in Science (true/false)? ")
    fmt.Scan(&science)

    if marks >= 60 {
        if math && science {
            fmt.Println("Eligible for admission")
        } else if !math && !science {
            fmt.Println("Not passed in Math and Science")
        } else if !math {
            fmt.Println("Not passed in Math")
        } else {
            fmt.Println("Not passed in Science")
        }
    } else {
        fmt.Println("Marks less than 60% - Not eligible")
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
    var a, b, c int
    fmt.Print("Enter 3 numbers: ")
    fmt.Scan(&a, &b, &c)

    if a >= b && a >= c {
        fmt.Println(a, "is largest")
    } else if b >= a && b >= c {
        fmt.Println(b, "is largest")
    } else {
        fmt.Println(c, "is largest")
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
    var n1, n2 float64
    var op string

    fmt.Print("Enter first number: ")
    fmt.Scan(&n1)
    fmt.Print("Enter second number: ")
    fmt.Scan(&n2)
    fmt.Print("Enter operator (+ - * /): ")
    fmt.Scan(&op)

    switch op {
    case "+":
        fmt.Println("Result:", n1+n2)
    case "-":
        fmt.Println("Result:", n1-n2)
    case "*":
        fmt.Println("Result:", n1*n2)
    case "/":
        if n2 != 0 {
            fmt.Println("Result:", n1/n2)
        } else {
            fmt.Println("Cannot divide by zero")
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
import "fmt"

func main() {
    var year int
    fmt.Print("Enter year: ")
    fmt.Scan(&year)

    if year%400 == 0 {
        fmt.Println("Century leap year")
    } else if year%4 == 0 && year%100 != 0 {
        fmt.Println("Normal leap year")
    } else {
        fmt.Println("Not a leap year")
    }
}


// =========================================================
