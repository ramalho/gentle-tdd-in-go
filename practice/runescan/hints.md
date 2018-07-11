# Hints for a `runescan` practice

## Example test

```go
func Example() {
	main()
	// Output:
	// Please provide one or more words to search.
}
```


## Sample `UnicodeData.txt` line

```
005A;LATIN CAPITAL LETTER Z;Lu;0;L;;;;;N;;;;007A;
```


## A struct to represent a character and its name

```go
type CharName struct {
	char rune
	name string
}
```

## Convert hexadecimal code string to rune

```go
code, err := strconv.ParseInt(hexaStr, 16, 32)
char := rune(code)
```
