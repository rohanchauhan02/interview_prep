package main

import (
	"fmt"

	array "github.com/rohanchauhan02/interview_prep/Array"
	recursion "github.com/rohanchauhan02/interview_prep/Recursion"
	str "github.com/rohanchauhan02/interview_prep/String"
)

func main() {
	// _01()
	_02()
	// _03()
	// _04()
	// _05()
	// _06()
	// _07()
}

// array
func _01() {
	// 1, 2, 3, 4, 5, 6, 7, 8, 9
	nums := []int{1, 3, 5, 8, 7, 4, 2, 9}
	fmt.Println("01. ", array.MissingNumber(nums))
	fmt.Println("02. ", array.ReverseArray(nums))
}

// string
func _02() {
	// s := "hello"
	// s1 := "ababac"
	// fmt.Println("02. ", str.ReverseStr(s))
	// fmt.Println("03. ", str.LongestSubstring(s1))
	fmt.Println("04. ", str.IsAnagram("listen", "silent"))
	fmt.Println("06. ", str.FirstNonRepeatingChar("helloh"))
	fmt.Println("07. ", str.Palindrome("racecar"))
}

// recursion
func _03() {
	keyboard := []string{",;", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	fmt.Println("03. Subsequence", recursion.Subsequence("abc"))
	fmt.Println("03. KeyboardCombination", recursion.KeyboardCombination(keyboard, "456"))
}

func _04() {

}

func _05() {

}

func _06() {

}

func _07() {

}

/*

1. Arrays & Strings
	1. Find the missing number in a given array of size n containing numbers from 1 to n+1.
	2. Reverse a string without using extra space.
	3. Find the longest substring without repeating characters in a string.
	4. Check if two strings are anagrams of each other.
	5. Implement string matching algorithms (Naive, KMP).
	6. Find the first non-repeating character in a string.
	7. Check if a string is a palindrome.
2. Linked List
	1. Reverse a singly linked list.
	2. Detect a cycle in a linked list (Floyd’s cycle-finding algorithm).
	3. Find the intersection point of two linked lists.
	4. Merge two sorted linked lists.
	5. Find the middle element of a linked list.
	6. Flatten a linked list (nested linked list).
	7. Remove N-th node from the end of a linked list.
3. Stacks & Queues
	1. Implement a stack using two queues.
	2. Implement a queue using two stacks.
	3. Evaluate a postfix expression.
	4. Check for balanced parentheses in an expression.
	5. Implement a circular queue.
	6. Implement a stack supporting push, pop, and min operations.
4. Trees
	1. Traverse a binary tree (In-order, Pre-order, Post-order).
	2. Check if a binary tree is balanced (Height-balanced tree).
	3. Find the Lowest Common Ancestor (LCA) of two nodes in a binary tree.
	4. Convert a binary tree to a doubly linked list.
	5. Implement a binary search tree (BST).
	6. Find the diameter of a binary tree.
	7. Level-order traversal of a binary tree.
	8. Serialize and deserialize a binary tree.
	9. Check if a binary tree is a valid BST.
5. Heaps
	1. Implement a min-heap.
	2. Find the kth largest element in an array using a heap.
	3. Merge k sorted arrays using a min-heap.
	4. Implement a priority queue.
6. Graphs
	1. Implement Depth-First Search (DFS) and Breadth-First Search (BFS).
	2. Find the shortest path in an unweighted graph (Breadth-First Search).
	3. Detect a cycle in a directed graph (using DFS).
	4. Find the topological sort of a directed acyclic graph (DAG).
	5. Detect a cycle in an undirected graph.
	6. Find the strongly connected components (SCC) using Tarjan’s or Kosaraju’s algorithm.
	7. Dijkstra’s algorithm for shortest paths.
	8. Find the shortest path between two nodes in a weighted graph.
7. Dynamic Programming
	1. Find the longest increasing subsequence (LIS).
	2. Find the minimum number of coins required for a given value.
	3. 0/1 Knapsack problem.
	4. Longest common subsequence (LCS).
	5. Matrix chain multiplication.
	6. Find the number of unique paths in a grid from top-left to bottom-right.
	7. Palindrome Partitioning: Find minimum cuts to make a string palindrome.
8. Backtracking
	1. Solve the N-Queens problem.
	2. Generate all subsets of a given set.
*/
