// Assignment: Conditional Statements in Go (Golang)
// Topics Covered: if, else if, else, nested if, short if, switch, fallthrough, switch without expression

// -----------------------------------------------------------
// Section A: Easy (Fundamentals)
// -----------------------------------------------------------

// 1. Write a program to check if a number is positive, negative, or zero using if statements.

// =========================================================


package main 
import "fmt"

func main(){
	num := 0

	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0{
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
func main() {
	age := 18 
	if age >= 18{
		fmt.Println("Eligible")
	} else{
		fmt.Println("Not Eligible")
	}
}


// =========================================================


// 3. Write a program that checks whether a given number is even or odd.

// =========================================================


package main
import "fmt"

func main(){
	num := 10
	if num % 2 == 0{
		fmt.Println("Even")
	} else{
		fmt.Println("Odd")
	}
}



// =========================================================


// 4. Write a program that takes a student’s marks and prints “Pass” or “Fail” using an if-else condition.

// =========================================================


package main
import "fmt"

func main(){
	marks := 80
	if marks >= 30{
		fmt.Println("Pass")
	} else{
		fmt.Println("Fail")
	}
}



// =========================================================

// 5. Write a program to check whether a given year is a leap year or not.

// =========================================================


package main
import "fmt"
 func main(){
	year := 2024
	if year % 4 == 0 && year % 100 != 0 || year % 400 == 0{
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Not a leap year")
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
func main(){
	marks := 75
	if marks >= 90 {
		fmt.Println("Grade A")
	} else if marks >= 75{
		fmt.Println("Grade B")
	} else if marks >= 60{
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}
}



// =========================================================

// 2. Write a program that takes two numbers and prints which one is greater using if-else.

// =========================================================


package main 
import "fmt"
func main(){
	a := 9
	b := 56

	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}



// =========================================================

// 3. Write a program that checks if a character is a vowel or consonant using a switch statement.

// =========================================================

package main 
import "fmt"
 func main(){
	alpha := "J"
	switch alpha {
	case "a", "A": 
		fmt.Println("Vowel")
	case "e", "E":
		fmt.Println("Vowel")
	case "i", "I":
		fmt.Println("Vowel")
	case "o", "O":
		fmt.Println("Vowel")
	case "u", "U":
		fmt.Println("Vowel")
	default :
		fmt.Println("Consonant")
	} 
}



// =========================================================

// 4. Write a program that takes a day name (like Monday, Tuesday, etc.) and prints whether it’s a weekday or weekend using switch.

// =========================================================


package main
import "fmt"
func main(){
	day := "Sunday"
	switch day{
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
func main(){
	num := 15
	if num % 3 == 0 && num % 5 == 0{
		fmt.Println("Yes")
	} else{
		fmt.Println("No")
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
func main (){
	username := "Pappu"
	password := "0000"
	if username == "Pappu" && password == "0000"{
		fmt.Println("Login Successful")
	} else if username != "Pappu" {
		fmt.Println("User not found.")
	} else{
		fmt.Println("Invalid Password")
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
func main(){ 
	a := 9
	b := 8
	c := 7
	if a == b && b == c && a == c {
		fmt.Println("Equilateral Triangle")
	} else if a == b || b == c || a == c{
		fmt.Println("Isosceles Triangle")
	} else if a != b && b != c && a != c{
		fmt.Println("Scalene Triangle")
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
func main(){
	x := 1
	switch x {
	case 1 : 
	fmt.Println("Start Game")
	case 2 :
	fmt.Println("Load Game")
	case 3 :
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


package main
import "fmt"
 func main(){
	temp := -8;
	switch {
	case temp > 30:
	fmt.Println("Hot")
	case temp >= 16 :
	fmt.Println("Warm")
	case temp >= 15:
	fmt.Println("Cold")
	case temp > 0:
	fmt.Println("Cold")
	case temp < 0 :
	fmt.Println("Freezing")	
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
func main(){
	
	marks := 89
	maths := true
	science := false

	if marks > 60 && maths && science{
		fmt.Println("Eligible for admission")
	} else if marks < 60 {
		fmt.Println("Marks criteria not cleared")
	} else if maths == false {
		fmt.Println("Not passed in maths")
	} else {
		fmt.Println("Not passed in science")
	}
}



// =========================================================


// -----------------------------------------------------------
// Bonus (Challenge Zone)
// -----------------------------------------------------------

// 1. Write a program that takes three numbers as input and prints the largest using only conditional statements.

// Hint: Use nested if or multiple if-else conditions.


// =========================================================
func main(){
    var a, b , c int
    fmt.Scan(&a , &b ,&c)
    if  a > b{
        if a > c {
            fmt.Print(a)
        }else {
        fmt.Print(c)
        }
    }else if  b > a{
        if b > c {
            fmt.Print(b)
        }else {
        fmt.Print(c)
        }
    }else if  c > b{
        if c > a {
            fmt.Print(c)
        }else {
        fmt.Print(a)
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
package main 
import "fmt"

func main(){
    var op string
    var a, b int
    fmt.Scan(&op,&a, &b )
    if  op == "+" {
        fmt.Print(a + b)
    }else if op == "-" {
        fmt.Print(a - b)
    }else if op == "*" {
        fmt.Print(a * b)
    } else if op == "/"{
        fmt.Print ( a / b)
    }  else {
        fmt.Print("Invalid operator")
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

func main(){
    var a int
    fmt.Scan(&a)
    if  a % 400 == 0 {
        fmt.Print("Century leap year")
    }else if a % 4 == 0 && a%100 != 0{
        fmt.Print("normal leap year")
    }else {
        fmt.Print("I it is not a leap year")
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
var a int
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

func main(){
    var mark , attend int
    var backlog bool
    fmt.Scan(&mark, &attend ,&backlog)
    if  mark >= 85 && attend >90 && backlog {
        fmt.Print("Scholarship Approved")
    }else {
        fmt.Print("Scholarship Denied")
    }
}




// =========================================================
