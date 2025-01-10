package str

/*
Reverse a string without using extra space.
For example:
	Input: "hello"
	Output: "olleh"

In Go, strings are immutable, meaning you cannot directly modify a string.
Instead, you need to work with a slice of rune (or byte for ASCII-only strings).
Once reversed, you can convert the slice back to a string.

*/

func ReverseStr(s string) string {
	runes := []rune(s)
	pt1, pt2 := 0, len(s)-1

	for pt1 < pt2 {
		runes[pt1], runes[pt2] = runes[pt2], runes[pt1]
		pt1++
		pt2--
	}
	return string(runes)
}

/*
Example: Given a string s, find the length of the longest substring that contains only unique characters.
	Input:
    	s = "abcabcbb"
	Output:
		Length = 3
		Longest substring = "abc"

Steps:
	Use Two Pointers:

	start tracks the beginning of the current window.
	end expands the window by iterating through the string.
	HashMap for Tracking:

	Use a map to store the characters in the current window and their last seen positions.
	Adjust the Window:

	If a duplicate character is found, move the start pointer to the position after the duplicate’s last occurrence.
	Update Maximum Length:

	At each step, calculate the length of the current window and update the maximum length.

*/

func LongestSubstring(s string) int {
	hasSet := make(map[rune]bool)
	start := 0
	lsub := 0
	for _, ele := range s {
		if hasSet[ele] {
			// delete till ele not deleted
			for hasSet[ele] {
				delete(hasSet, rune(s[start]))
				start++
			}
		}
		hasSet[ele] = true
		if lsub < len(hasSet) {
			lsub = len(hasSet)
		}
	}

	return lsub
}

/*
Problem Description:
	Given two strings, return true if they are anagrams of each other, and false otherwise.

Definition of Anagrams
	Two strings are anagrams if:
		They have the same length.
		Rearranging the characters of one string can result in the other string.
Example:
	"listen" and "silent" are anagrams.
	"triangle" and "integral" are anagrams.
	"hello" and "world" are not anagrams.
*/

// 4. Check if two strings are anagrams of each other.
func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	hasSet := make(map[rune]int)

	for _, ele := range s1 {
		hasSet[ele]++
	}
	for _, ele := range s2 {
		if hasSet[ele] == 0 {
			return false
		}
		hasSet[ele]--
		if hasSet[ele] == 0 {
			delete(hasSet, ele)
		}
	}
	return len(hasSet) == 0
}

// 6. Find the first non-repeating character in a string.
func FirstNonRepeatingChar(s string) string {
	hasSet := make(map[rune]int)
	for _, ele := range s {
		hasSet[ele]++
	}

	for _, ele := range s {
		if hasSet[ele] == 1 {
			return string(ele)
		}
	}

	return ""
}

// 7. Check if a string is a palindrome.
func Palindrome(s string) bool {
	pt1, pt2 := 0, len(s)-1
	for pt1 < pt2 {
		if s[pt1] != s[pt2] {
			return false
		}
		pt1++
		pt2--
	}
	return true
}

