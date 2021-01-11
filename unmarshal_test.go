package njson

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnmarshalSmall(t *testing.T) {
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

	actual := User{}

	err := Unmarshal([]byte(json), &actual)
	if err != nil {
		t.Error(err)
	}

	expected := User{
		Name:    "Shapan",
		Age:     26,
		Friends: []string{"Asma", "Ahmed", "Mahmoud"},
	}

	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Error(diff)
	}

}

func TestUnmarshallBasicTypes(t *testing.T) {
	json := `
	{
        "name": "Shapan",
		"number": -42,
		"number_8": 8,
		"number_16": -16,
		"number_32": 32,
		"number_64": -64,
		"unsigned_number": 142,
		"unsigned_number_8": 18,
		"unsigned_number_16": 116,
		"unsigned_number_32": 132,
		"unsigned_number_64": 164,
		"has_email": true,
		"unicode": 8984,
		"percentage_32": 32.51,
		"percentage_64": 64.5248,
        "friends": [
            {"name": "Asma", "age": 26},
            {"name": "Ahmed", "age": 25},
            {"name": "Mahmoud", "age": 30}
        ]
	}`

	type Types struct {
		Name     string `njson:"name"`
		Number   int    `njson:"number"`
		Number8  int8   `njson:"number_8"`
		Number16 int16  `njson:"number_16"`
		Number32 int32  `njson:"number_32"`
		Number64 int64  `njson:"number_64"`
		// UnsignedNumber uint   `njson:"unsigned_number"`
		// UnsignedNumber8  uint8   `njson:"unsigned_number_8"`
		// UnsignedNumber16 uint16  `njson:"unsigned_number_16"`
		// UnsignedNumber32 uint32  `njson:"unsigned_number_32"`
		// UnsignedNumber64 uint64  `njson:"unsigned_number_64"`
		HasEmail     bool    `njson:"has_email"`
		Unicode      rune    `njson:"unicode"`
		Percentage32 float32 `njson:"percentage_32"`
		Percentage64 float64 `njson:"percentage_64"`
	}

	actual := Types{}

	err := Unmarshal([]byte(json), &actual)
	if err != nil {
		t.Error(err)
	}

	expected := Types{
		Name:     "Shapan",
		Number:   -42,
		Number8:  8,
		Number16: -16,
		Number32: 32,
		Number64: -64,
		// UnsignedNumber: 142,
		// UnsignedNumber8:  18,
		// UnsignedNumber16: 116,
		// UnsignedNumber32: 132,
		// UnsignedNumber64: 164,
		HasEmail:     true,
		Unicode:      8984,
		Percentage32: 32.51,
		Percentage64: 64.5248,
	}

	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Error(diff)
	}

}

func TestUnmarshalComplex(t *testing.T) {

	json := `
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
	`

	type User struct {
		FirstName              string   `njson:"name.first"`
		LastName               string   `njson:"name.last"`
		Age                    int      `njson:"age"`
		NumberOfChildren       int      `njson:"children.#"`
		Children               []string `njson:"children"`
		Friends                []string `njson:"friends.#.last"`
		NumberOfFriendNetworks []int    `njson:"friends.#.nets.#"`
		//FavoriteMovie    string   `njson:"fav.movie"`
	}

	actual := User{}

	err := Unmarshal([]byte(json), &actual)
	if err != nil {
		t.Error(err)
	}

	expected := User{
		FirstName:              "Tom",
		LastName:               "Anderson",
		Age:                    37,
		NumberOfChildren:       3,
		Children:               []string{"Sara", "Alex", "Jack"},
		Friends:                []string{"Murphy", "Craig", "Murphy"},
		NumberOfFriendNetworks: []int{3, 2, 2},
	}

	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Error(diff)
	}

}

func TestUnmarshalMoreComplex(t *testing.T) {

	json := `
	{
		"top_struct": {
			"first": "Tom", 
			"last": "Anderson",
			"child_struct": {
				"first_child": "Child 1",
				"another_child_struct": {
					"first_child_of_first_child": "The first child"
				},
				"a_child_list": ["six", "five", "four"]
			},
			"a_list": [5, 4, 3, 2, 1]
		},
		"list_of_lists": [
			[101, 102, 103],
			[201, 202, 203],
			[301, 302, 303]
		]
		"list_of_structs": [
		  {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
		  {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
		  {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
		]
	  }
	`

	type MoreComplex struct {
		FirstChild         string   `njson:"top_struct.child_struct.first_child"`
		FirstChildOfChild  string   `njson:"top_struct.child_struct.another_child_struct.first_child_of_first_child"`
		ListInStruct       []int    `njson:"top_struct.a_list"`
		AChildListInStruct []string `njson:"top_struct.child_struct.a_child_list"`
		//ListOfLists        [][]int  `njson:"list_of_lists"`
		ListOfStructs       []string `njson:"list_of_structs"`
		ListOfNetsInStructs []string `njson:"list_of_structs.#.nets"`
		// ListOfNetLists      [][]string `njson:"list_of_structs.#.nets.#.*"`
	}

	actual := MoreComplex{}

	err := Unmarshal([]byte(json), &actual)
	if err != nil {
		t.Error(err)
	}

	expected := MoreComplex{
		FirstChild:         "Child 1",
		FirstChildOfChild:  "The first child",
		ListInStruct:       []int{5, 4, 3, 2, 1},
		AChildListInStruct: []string{"six", "five", "four"},
		//ListOfLists:        [][]int{{101, 102, 103}, {201, 202, 203}, {301, 302, 303}},
		ListOfStructs: []string{
			`{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]}`,
			`{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]}`,
			`{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}`,
		},
		ListOfNetsInStructs: []string{
			`["ig", "fb", "tw"]`,
			`["fb", "tw"]`,
			`["ig", "tw"]`,
		},
		// ListOfNetLists: [][]string{{"ig", "fb", "tw"}, {"fb", "tw"}, {"ig", "tw"}},
	}

	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Error(diff)
	}

}
