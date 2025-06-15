package validator

//TODO:
//create a struct Result, with Password string, OK bool, Reasons []string fields

//TODO:
//Create a Checker interface
//declare a Check function, which takes string as a parameter and returns Result

//TODO:
//Create PasswordChecker struct, with minLength int and blacklist map[string]bool fields

//TODO:
//Create a consturctor for PasswordChecker - Create a function called New, which takes parameters
//minLength int and banned []string
//This function/constructor should return a new PasswordChecker or a pointer to it with minLength and
//blacklisted words initialized. blacklist is created from the slice of strings named 'banned'

//TODO:
//Implement the Checker interface for PasswordChecker
//It should check the string in parameters against various conditions:
//1. is it longer than the minLength ?
//2. is it in the blacklist?
//3. Does it contain an uppercase and lowercase letter?
//4. Does it contain a digit ?
//5. Does it contain a symbol?
//For each condition that is not true, add the reason to the "Result"'s reasons
//If every condition is true, Set the 'OK' in "Result" to true
//return the Result with the string passed to this method as a "password"
