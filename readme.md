# NJson
NJSON is a Go package that Unmarshal/Decode nested JSON data depend on provided path

## Installation
```bash
$ go get -u github.com/m7shapan/njson
```

## Example

```go
package main

import (
	"fmt"

	"github.com/m7shapan/njson"
)

func main() {
    json := `
	{
        "name": {"first": "Mohamed", "last": "Shapan"},
        "age": 26,
        "friends": [
            {"name": "Asma", "age": 26},
            {"name": "Ahmed", "age": 25},
            {"name": "Mahmoud", "age": 30}
        ]
	}`

	type User struct {
		Name    string   `njson:"name.last"`
		Age     int      `njson:"age"`
		Friends []string `njson:"friends.#.name"`
	}

	u := User{}

	err := njson.Unmarshal([]byte(json), &u)
	if err != nil {
		// do anything
	}

	fmt.Printf("%+v\n", u) // {Name:Shapan Age:26 Friends:[Asma Ahmed Mahmoud]}
}
```

## Path Syntax
A path is a series of keys separated by a dot. A key may contain special wildcard characters '*' and '?'. To access an array value use the index as the key. To get the number of elements in an array or to access a child path, use the '#' character. The dot and wildcard characters can be escaped with '\'.
```json
{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}
```

```
"name.last"          >> "Anderson"
"age"                >> 37
"children"           >> ["Sara","Alex","Jack"]
"children.#"         >> 3
"children.1"         >> "Alex"
"child*.2"           >> "Jack"
"c?ildren.0"         >> "Sara"
"fav\.movie"         >> "Deer Hunter"
"friends.#.first"    >> ["Dale","Roger","Jane"]
"friends.1.last"     >> "Craig"
```

## TODOs
- [x] Add test cases 
- [ ] Improve `map` type Unmarshal/Decode performance
- [ ] Improve `struct` type Unmarshal/Decode performance

## Contact
Mohamed Shapan [@m7shapan](https://twitter.com/M7Shapan)

## Oh, Thanks!
By the way i want to thank Josh Baker for his package [gjson](https://github.com/tidwall/gjson)
