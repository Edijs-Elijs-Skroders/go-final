package main

import (
    "flag"
    "fmt"

    "github.com/fatih/color"
    "final-task/internal/validator"
)

//TODO:
//using the standard flag package,parse the following command-line flags: 
// --passwords, type string, list of passwords, separated by a comma
// --min, type int, min password length
// --banned,  type string, list of banned passwords, separated by a comma
//Validate, that the passwords is not empty - if it is, print an error and exit the program
//Create the necessary data structures for PasswordChecker, checks and results
//Crate a new PasswordChecker using the New function
//Check all the passwords using PasswordChecker Check method 
//Print out results using "github.com/fatih/color": 
// if the result is OK, print out "OK: " and add the password. Make this text's color green
// else, print out"FAIL: "and add the password plus all the reasons. Make this text's color red
func main(){

}