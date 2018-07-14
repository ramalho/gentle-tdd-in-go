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

## More sample lines

```
002D;HYPHEN-MINUS;Pd;0;ES;;;;;N;;;;;
002E;FULL STOP;Po;0;CS;;;;;N;PERIOD;;;;
002F;SOLIDUS;Po;0;CS;;;;;N;SLASH;;;;
0030;DIGIT ZERO;Nd;0;EN;;0;0;0;N;;;;;
0031;DIGIT ONE;Nd;0;EN;;1;1;1;N;;;;;
0032;DIGIT TWO;Nd;0;EN;;2;2;2;N;;;;;
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

## Table test with subtests

```go
func TestMatch(t *testing.T) {
		testCases := []struct {
				query string
				name  string
				want  bool
		}{
				{"BLACK", "BLACK CHESS KING", true},
				{"CHESS BLACK", "WHITE CHESS KING", false},
				{"CHESS BLACK", "BLACK CHESS KING", true},
		}
		for _, tc := range testCases {
				t.Run(fmt.Sprintf("%v in %v", tc.query, tc.name), func(t *testing.T) {
						queryTerms := strset.MakeFromText(tc.query)
						if got := match(queryTerms, tc.name); got != tc.want {
								t.Errorf("got %v; want %v", got, tc.want)
						}
				})
		}
}

```

## Building a buffer from a string

```go
data := strings.NewReader(dataSample)
```

## Use of testify

Importing:

```go
import "github.com/stretchr/testify/assert"
```

Simple assertion:

```go
	expected := []CharName{
		{'0', "DIGIT ZERO"},
		{'1', "DIGIT ONE"},
		{'2', "DIGIT TWO"},
	}
	assert.Equal(t, got, expected)
```

## Iterating over text file or buffer

```go
// Filter returns CharName records when query matches line in UnicodeData.txt
func Filter(data io.Reader, query string) []CharName {
	queryTerms := strset.MakeFromText(strings.ToUpper(query))
	scanner := bufio.NewScanner(data)
	result := []CharName{}
	for scanner.Scan() {
		charName := parseLine(scanner.Text())
		if match(queryTerms, charName.name) {
			result = append(result, charName)
		}
	}
	return result
}
```