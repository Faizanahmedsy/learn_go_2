package main

import "fmt"

/*
Let's simplify this even more.

### Imagine This as a Game:
- You have a number, and your job is to flip it around and see if it looks the same.

### Let's Play with the Number "12321":
1. **Start with 12321:**
   - Imagine writing the number down: `12321`.
   - We want to flip it around, so we look at the last digit first.

2. **First Step:**
   - Take the last digit (1) and start a new number with it. Now, our new number is `1`.

3. **Next Step:**
   - Move to the second last digit (2). Add it to our new number, so now it’s `12`.

4. **Keep Going:**
   - Add the next digit (3). Now, our new number is `123`.
   - Add the next digit (2). Now, our new number is `1232`.
   - Add the final digit (1). Now, our new number is `12321`.

5. **Check the Game:**
   - The new number we made by flipping the digits is `12321`, which is the same as the number we started with!
   - Since they match, we say "Yes, it’s a palindrome!" (a fancy word meaning it looks the same forwards and backwards).

### Key Point:
- We’re just flipping the number around like flipping a pancake, and then checking if it looks the same. If it does, it’s a palindrome.



*/

func main() {
	var number int

	fmt.Println("Please enter a number :)")
	fmt.Scan(&number)

	if check_palindrome(number) == true {
		fmt.Println("Yes")
	} else {
		fmt.Println(("No"))
	}

}

func check_palindrome(n int) bool {

	// We start with reverse = 0 and temp = 12321.
	var reverse int = 0
	var temp int = n

	/*
		Take the last digit: To get the last digit of temp, we do temp % 10 (which means "what’s the remainder when you divide by 10?").
		For 12321 % 10, the remainder is 1.
		Now, we update reverse by multiplying it by 10 (which shifts its digits left) and adding the last digit: reverse = 0 * 10 + 1 = 1.
		Then, remove the last digit from temp by dividing it by 10: temp = 12321 / 10 = 1232.

	*/

	for temp != 0 {
		reverse = (reverse * 10) + (temp % 10)
		temp = temp / 10
	}

	return (reverse == n)
}
