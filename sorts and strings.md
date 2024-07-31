# Go Standard Library Packages

The Go programming language comes with a standard library of packages that provide a variety of functions and features. 
These packages can be imported and used in Go programs to extend the language's functionality.

- Go's standard library is lightweight, with a limited set of built-in methods for working with strings, integers, arrays, and other data types.
- However, the standard library provides a set of packages that can be imported to access additional functionality.
- Each package in the standard library brings its own set of functions and values that can be used in Go programs.
- The standard library includes packages for working with strings, slices, and other common data structures.
- The `fmt` package is one example of a standard library package, providing functions like `printf()` and `sprintf()`.
- Other packages, such as the `strings` package, provide methods for manipulating and working with strings.
- The standard library packages are all available on the Go website, where you can find documentation and information about each package.
- To use a package, you need to import it into your Go program.

Strings Package
- The `strings` package functions do not modify the original string; they return a new string with the desired changes.
- The `strings` package provides a variety of functions for working with strings.
- Examples:
- `strings.Contains()`: Checks if a string contains a specific substring.
- `strings.ReplaceAll()`: Replaces all occurrences of a substring with another string.
- `strings.ToUpper()`: Converts a string to uppercase.
- `strings.Index()`: Finds the index of a substring within a string.
- `strings.Split()`: Splits a string into a slice of substrings based on a delimiter.


Sort Package
- The `sort` package provides functions for sorting slices of data, including slices of integers and slices of strings.
- Examples:
- `sort.Ints()`: Sorts a slice of integers in ascending order.
- `sort.Strings()`: Sorts a slice of strings in alphabetical order.
- `sort.Search()`: Performs a binary search on a sorted slice to find the index of a specific value.
- The sorting functions in the `sort` package modify the original slice; they do not return a new slice.
- The `sort.Ints()` function can be used to sort a slice of integers in-place.
- 
- The `sort.Search()` function can be used to find the index of a value in a sorted slice.
- If the value being searched for is not found in the sorted slice, `sort.Search()` returns the index where the value would be inserted to maintain the sorted order.

- The `sort.Strings()` function can be used to sort a slice of strings in alphabetical order.
- The `sort.SearchStrings()` function can be used to find the index of a string value in a sorted slice of strings.
