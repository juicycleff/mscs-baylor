package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Set represents a collection of unique string values.
type Set map[string]struct{}

// readSet reads a set of strings from the provided `reader` and returns it as a `Set`.
// It expects the first line of input to specify the number of strings to read.
// Each subsequent line is read and added to the set.
// The strings are trimmed of leading and trailing whitespaces before being added to the set.
func readSets(reader *bufio.Reader) (Set, Set) {
	inputA := make(Set)
	inputB := make(Set)

	// Read first input
	var n int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		inputA[line] = struct{}{}
	}

	// Read second input
	mLine, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(mLine))
	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		inputB[line] = struct{}{}
	}

	return inputA, inputB
}

// main is the entry point function of the program.
// It reads two sets from standard input, computes the union, intersection, and Cartesian product of the sets,
// and prints the results.
func main() {
	// we read in stand input put
	reader := bufio.NewReader(os.Stdin)
	setA, setB := readSets(reader)

	unionOutput := union(setA, setB)
	intersectOutput := intersection(setA, setB)
	cartesianProdOutput := cartesianProduct(setA, setB)

	emptyTitle := ""

	printSet(nil, unionOutput)
	printSet(&emptyTitle, intersectOutput)
	printCartesianProduct("", cartesianProdOutput)
}

// union returns the union of two sets.
// It takes two sets, setA and setB, as input parameters and returns another set that contains all the elements from both setA and setB.
func union(setA, setB Set) Set {
	unionSet := make(Set)
	for key := range setA {
		unionSet[key] = struct{}{}
	}
	for key := range setB {
		unionSet[key] = struct{}{}
	}
	return unionSet
}

// intersect returns the intersection of two sets.
// It takes two sets, setA and setB, as input and returns a new set
// containing only the elements that are present in both setA and setB.
// The function initializes an empty intersectionSet.
// It then iterates over all keys in setA.
// If the current key is found in setB, it adds the key to the intersectionSet.
// Finally, it returns the intersectionSet.
func intersection(setA, setB Set) Set {
	intersectionSet := make(Set)
	for key := range setA {
		if _, found := setB[key]; found {
			intersectionSet[key] = struct{}{}
		}
	}
	return intersectionSet
}

// cartesianProduct calculates the Cartesian product of two sets and returns it as a slice of strings.
// Each element in the resulting slice represents an ordered pair of elements, one from setA and one from setB.
// The elements are enclosed in parentheses and separated by a comma.
// For example, if setA = {"a", "b"} and setB = {"x", "y"}, the Cartesian product would be [("a", "x"), ("a", "y"), ("b", "x"), ("b", "y")].
// The input sets should be of type Set, which is a map[string]struct{}.
// Each set represents a collection of unique elements.
// This function does not modify the input sets.
// The Cartesian product is calculated by iterating through each element of setA and setB and forming the ordered pair.
// The resulting ordered pairs are stored in a slice, which is then returned.
func cartesianProduct(setA, setB Set) []string {
	var cartesian []string
	for a := range setA {
		for b := range setB {
			cartesian = append(cartesian, fmt.Sprintf("%s %s", a, b))
		}
	}
	return cartesian
}

// printSet takes a title string and a set as input and prints the set elements in sorted order, preceded by the title.
func printSet(title *string, set Set) {
	var sorted []string
	for key := range set {
		sorted = append(sorted, key)
	}
	sort.Strings(sorted)

	if title != nil {
		fmt.Println(*title)
	}

	for _, key := range sorted {
		fmt.Println(key)
	}
}

// printCartesianProduct prints the title followed by each product in the Cartesian product of the two sets.
// The products are sorted in lexicographical order.
func printCartesianProduct(title string, cartesian []string) {
	sort.Strings(cartesian)
	fmt.Println(title)
	for _, product := range cartesian {
		fmt.Println(product)
	}
}
