package edn

import (
	. "github.com/bjeanes/go-edn/types"
	. "testing"
)

func parse(s string, t *T) (val Value) {
	val, err := ParseString(s)

	if err != nil {
		t.Errorf("Expected parsing %+v to succeed. %v", s, err)
		t.FailNow()
	}

	return
}

func assertValueEqual(actual, expected Value, t *T) {
	if !expected.Equals(actual) {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func TestDoesNotParseEmptyInput(t *T) {
	_, err := ParseString("")

	if err == nil {
		t.Error("Expected parsing empty input to fail")
	}
}

func TestParseList(t *T) {
	val := parse("()", t)

	if !new(List).Equals(val) {
		t.Errorf("Expected parsing \"()\" to return an empty list, got %+v", val)
	}

	str := `(() "abc" [] "def")`
	val = parse(str, t)

	l := new(List)
	l.Insert(new(List))
	l.Insert(String("abc"))
	l.Insert(Vector{})
	l.Insert(String("def"))

	if !l.Equals(val) {
		t.Errorf("Expected %v, got %v", l, val)
	}
}

func TestParseVector(t *T) {
	assertValueEqual(parse(`[]`, t), Vector{}, t)
	assertValueEqual(parse(`[[]]`, t), Vector{Vector{}}, t)
	assertValueEqual(parse(`[[("abc")] "def"]`, t),
		Vector{
			Vector{
				new(List).Insert(String("abc")),
			},
			String("def"),
		}, t,
	)
}

func TestParseMap(t *T) {
	_, err := ParseString("{:x}")

	if err.Error() != "Error: Map literal must contain an even number of forms" {
		t.Error("Expected parsing map with single item to fail")
	}

	assertValueEqual(parse(`{}`, t), Map{}, t)
	assertValueEqual(parse(`{:a :b}`, t), Map{Keyword("a"): Keyword("b")}, t)

}

func TestParseCharacter(t *T) {
	assertValueEqual(parse(`\r`, t), Character('r'), t)
	assertValueEqual(parse(`\\`, t), Character('\\'), t)
	assertValueEqual(parse(`\~`, t), Character('~'), t)
	assertValueEqual(parse(`\!`, t), Character('!'), t)
	assertValueEqual(parse(`\[`, t), Character('['), t)
	assertValueEqual(parse(`\5`, t), Character('5'), t)
	assertValueEqual(parse(`\space`, t), Character(' '), t)
	assertValueEqual(parse(`\return`, t), Character('\r'), t)
	assertValueEqual(parse(`\newline`, t), Character('\n'), t)
	assertValueEqual(parse(`\tab`, t), Character('\t'), t)
}

func TestParseString(t *T) {
	val := parse(`""`, t)
	if val == nil || !val.Equals(String("")) {
		t.Errorf("Expected \"\", got %v", val)
	}

	val = parse(`"abc"`, t)
	if val == nil || !val.Equals(String("abc")) {
		t.Errorf("Expected \"abc\", got %v", val)
	}
}

func TestParse(t *T) {
	expected := new(List).Insert(
		String("abc"),
		new(List).Insert(String("spaced")),
		Vector{
			String("vec"),
			new(List).Insert(
				String("an"),
				String("inner"),
				String("list"),
			),
		},
		new(List),
		Vector{},
		new(List),
		String(""),
		Vector{Character('x'), Character('\n')},
		Map{},
		Set{}.Insert(String("set")),
		Value(nil),
		Map{String("key"): String("value")},
		Map{
			String("key"):   String("value"),
			Keyword("key2"): new(List).Insert(Keyword("value")),
		},
		Keyword("key"),
	)

	actual := parse(`
	("abc" ( "spaced" )
	    ["vec"( "an"	"inner""list",)]
		(),[]()""[\x \newline]
		{}#{"set"} nil
		{"key" "value"}
		{"key" "value" :key2 (:value)}
		:key
	)
	`, t)

	assertValueEqual(actual, expected, t)
}
