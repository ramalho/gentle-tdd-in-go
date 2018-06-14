# Runes exercise playbook

## Sample codepoint entries

In Unicode version 10, `UnicodeData.txt` has 30,592 entries.

This is the entry for caracter `'A'`. The essential fields are the first and the second:
the hex code and the name of the character.

```
0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;
```

A sample of characters including one hyphenated name
(`"HYPHEN MINUS"`):
 
```
0027;APOSTROPHE;Po;0;ON;;;;;N;APOSTROPHE-QUOTE;;;;
0028;LEFT PARENTHESIS;Ps;0;ON;;;;;Y;OPENING PARENTHESIS;;;;
0029;RIGHT PARENTHESIS;Pe;0;ON;;;;;Y;CLOSING PARENTHESIS;;;;
002A;ASTERISK;Po;0;ON;;;;;N;;;;;
002B;PLUS SIGN;Sm;0;ES;;;;;N;;;;;
002C;COMMA;Po;0;CS;;;;;N;;;;;
002D;HYPHEN-MINUS;Pd;0;ES;;;;;N;;;;;
002E;FULL STOP;Po;0;CS;;;;;N;PERIOD;;;;
002F;SOLIDUS;Po;0;CS;;;;;N;SLASH;;;;
0030;DIGIT ZERO;Nd;0;EN;;0;0;0;N;;;;;
0031;DIGIT ONE;Nd;0;EN;;1;1;1;N;;;;;
0032;DIGIT TWO;Nd;0;EN;;2;2;2;N;;;;;
```

## Simple test example

```go
func TestParseInt(t *testing.T) {
	input := "2A"
	want := int64(42)
	got, _ := strconv.ParseInt(input, 16, 7)
	if got != want {
		t.Errorf("Parsing %q, got: %d, want: %d", input, got, want)
	}
}
```

## Test example with `Fatal` condition

When `t.Fatal` is called, the entire test function is aborted; no further checks are executed in that function.
Other test functions remaining will still run. 

```go
func TestParseInt(t *testing.T) {
	input := "2A"
	want := int64(42)
	got, err := strconv.ParseInt(input, 16, 7)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Parsing %q, got: %d, want: %q", input, got, want)
	}
}
```

## Table test example

```go
func TestSplitN(t *testing.T) {
	var testCases = []struct {
		text  string
		sep   string
		parts int
		want  []string
	}{
		{"A-B-C-D", "-", -1, []string{"A", "B", "C", "D"}},
		{"A-B-C-D", "|", -1, []string{"A-B-C-D"}},
		{"", "-", -1, []string{""}},
		{"A-B-C-D", "-", 0, []string{}},
		{"", "-", 0, []string{}},
		{"A-B-C-D", "-", 2, []string{"A", "B-C-D"}},
		{"A-B-C-D", "-", 3, []string{"A", "B", "C-D"}},
		{"A-B-C-D", "-", 8, []string{"A", "B", "C", "D"}},
	}
	for _, tc := range testCases {
		testName := fmt.Sprintf("%q,%q,%d", tc.text, tc.sep, tc.parts)
		t.Run(testName, func(t *testing.T) {
			got := strings.SplitN(tc.text, tc.sep, tc.parts)
			if !equalSlices(tc.want, got) {
				t.Errorf("got: %#v, want: %#v", got, tc.want)
			}
		})
	}
}

```
