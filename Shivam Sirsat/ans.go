package main

import "fmt"

// Assignment: Conditional Statements in Go (Golang)
// Topics Covered: if, else if, else, nested if, short if, switch, fallthrough, switch without expression

// -----------------------------------------------------------
// Section A: Easy (Fundamentals)
// -----------------------------------------------------------

// 1. Write a program to check if a number is positive, negative, or zero using if statements.

// =========================================================

func main(){
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	if n > 0{
		fmt.Print("positive")

	}else if n<0{
		fmt.Print("negative")
	}else {
		fmt.Print("zero")
	}


// =========================================================

// 2. Write a program to determine whether a person is eligible to vote based on age.

// =========================================================
	age := 15
	if age >= 18 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}







// =========================================================


// 3. Write a program that checks whether a given number is even or odd.

// =========================================================
num := 10
if  num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}



// =========================================================


// 4. Write a program that takes a student’s marks and prints “Pass” or “Fail” using an if-else condition.

// =========================================================
score := 85
	if score >= 40 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Grade: F")
	}



// =========================================================

// 5. Write a program to check whether a given year is a leap year or not.

// =========================================================
year := 10
if  year%4 == 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
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
s := 85
	if s >= 90 {
		fmt.Println("Grade: A")
	} else if s >= 75 {
		fmt.Println("Grade: B")
	} else if s >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}



// =========================================================

// 2. Write a program that takes two numbers and prints which one is greater using if-else.

// =========================================================
m:=4
l:=5
if l > m{
	fmt.Print("l")
}else{
	fmt.Print("m")
}






// =========================================================

// 3. Write a program that checks if a character is a vowel or consonant using a switch statement.

// =========================================================
var ch = 'a'
switch ch {
    case 'a', 'e', 'i', 'o', 'u':
        fmt.Println("Vowel")
    default:
        fmt.Println("Consonant")
    }



// =========================================================

// 4. Write a program that takes a day name (like Monday, Tuesday, etc.) and prints whether it’s a weekday or weekend using switch.

// =========================================================
var day = "Monday"
 switch day {
    case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
        fmt.Println("Weekday")
    case "Saturday", "Sunday":
        fmt.Println("Weekend")
    default:
        fmt.Println("Invalid day name")
    }


// =========================================================

// 5. Write a program that uses a short if statement to check if a number is divisible by 3 and 5 at the same time.

// =========================================================
k:= 87
if k%3 == 0 && k%5 == 0 {
        fmt.Println("Divisible by both 3 and 5")
    } else {
        fmt.Println("Not divisible by both 3 and 5")
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
var username, password string
	cU := "admin"
    cP := "12345"

    fmt.Print("Enter username: ")
    fmt.Scan(&username)

    fmt.Print("Enter password: ")
    fmt.Scan(&password)

    if username == cU && password == cP {
        fmt.Println("Login successful")
    } else if username == cU && password != cP {
        fmt.Println("Invalid password")
    } else {
        fmt.Println("User not found")
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
var a, b, c int
    fmt.Print("Enter three sides of the triangle: ")
    fmt.Scan(&a, &b, &c)

    if a == b && b == c {
        fmt.Println("Equilateral triangle")
    } else {
        if a == b || b == c || a == c {
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
}