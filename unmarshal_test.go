package njson

import (
	json2 "encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestUnmarshalInvalidJson(t *testing.T) {
	json := `
	BAD JSON %% @
	##
	}`

	type Name struct {
		First string `njson:"first"`
		Last  string `njson:"last"`
	}

	type User struct {
		Name    Name   `njson:"name"`
		Age     int    `njson:"age"`
		Friends []Name `njson:"friends"`
	}

	actual := User{}

	if err := Unmarshal([]byte(json), &actual); err == nil {
		t.Error("error should not be nil")
	}
}

func TestUnmarshalError(t *testing.T) {
	json := `
	{
        "name": {"first": "Mohamed", "last": "Shapan"},
        "age": 26,
        "friends": [
            {"first": "Asma", "age": 26},
            {"first": "Ahmed", "age": 25},
            {"first": "Mahmoud", "age": 30}
        ]
	}`

	if err := Unmarshal([]byte(json), nil); err == nil {
		t.Error("error should not be nil")
	}
}

func TestUnmarshalByValueError(t *testing.T) {
	json := `
	{
        "name": {"first": "Mohamed", "last": "Shapan"},
        "age": 26,
        "friends": [
            {"first": "Asma", "age": 26},
            {"first": "Ahmed", "age": 25},
            {"first": "Mahmoud", "age": 30}
        ]
	}`

	type Name struct {
		First string `njson:"first"`
		Last  string `njson:"last"`
	}

	type User struct {
		Name    Name   `njson:"name"`
		Age     int    `njson:"age"`
		Friends []Name `njson:"friends"`
	}

	actual := User{}

	if err := Unmarshal([]byte(json), actual); err == nil {
		t.Error("error should not be nil")
	}
}

func TestUnmarshalSmall(t *testing.T) {
	json := `
	{
        "name": {"first": "Mohamed", "last": "Shapan"},
        "age": 26,
        "friends": [
            {"first": "Asma", "age": 26},
            {"first": "Ahmed", "age": 25},
            {"first": "Mahmoud", "age": 30}
        ]
	}`

	type Name struct {
		First string `njson:"first"`
		Last  string `njson:"last"`
	}

	type User struct {
		Name    Name   `njson:"name"`
		Age     int    `njson:"age"`
		Friends []Name `njson:"friends"`
	}

	actual := User{}

	err := Unmarshal([]byte(json), &actual)
	if err != nil {
		t.Error(err)
	}

	var friends []Name
	friends = append(friends, Name{
		First: "Asma",
	})

	friends = append(friends, Name{
		First: "Ahmed",
	})

	friends = append(friends, Name{
		First: "Mahmoud",
	})

	expected := User{
		Name: Name{
			First: "Mohamed",
			Last:  "Shapan",
		},
		Age:     26,
		Friends: friends,
	}

	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Error(diff)
	}

}

func TestUnmarshalJsonNumber(t *testing.T) {
	json := `
	{
        "name": {"first": "Mohamed", "last": "Shapan"},
        "age": 26,
        "weight": "77",
        "friends": [
            {"first": "Asma", "age": 26},
            {"first": "Ahmed", "age": 25},
            {"first": "Mahmoud", "age": 30}
        ]
	}`

	type Name struct {
		First string `njson:"first"`
		Last  string `njson:"last"`
	}

	type User struct {
		Name    Name         `njson:"name"`
		Age     json2.Number `njson:"age"`
		Weight  json2.Number `njson:"weight"`
		Friends []Name       `njson:"friends"`
	}

	actual := User{}

	err := Unmarshal([]byte(json), &actual)
	if err != nil {
		t.Error(err)
	}

	var friends []Name
	friends = append(friends, Name{
		First: "Asma",
	})

	friends = append(friends, Name{
		First: "Ahmed",
	})

	friends = append(friends, Name{
		First: "Mahmoud",
	})

	expected := User{
		Name: Name{
			First: "Mohamed",
			Last:  "Shapan",
		},
		Age:     json2.Number("26"),
		Weight:  json2.Number("77"),
		Friends: friends,
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

func TestUnmarshalFlattenSlices(t *testing.T) {
	json := `
		[1, 2, 3, 4]
	`

	type Slice struct {
		Numbers1 []int `njson:"@flatten"`
	}

	actual := Slice{}

	err := Unmarshal([]byte(json), &actual)
	if err != nil {
		t.Error(err)
	}

	expected := Slice{
		Numbers1: []int{1, 2, 3, 4},
	}

	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Error(diff)
	}
}

func TestUnmarshalSlices(t *testing.T) {
	json := `
	{
		"top_slice": ["foo", "bar"],
		"int_matrix": [
			[101, 102, 103],
			[201, 202, 203],
			[301, 302, 303]
		],
		"int_3d_matrix": [
			[
				[
					101, 102, 103
				],
				[
					201, 202, 203
				],
				[
					301, 302, 303
				]
			],
			[
				[
					401, 402, 403
				],
				[
					501, 502, 503
				]
			],
			[
				[
					601, 602, 603
				]
			]

		],
		"string_matrix": [
			["string_1", "string_2"],
			["string_3", "string_4"]
		],
		"string_3d_matrix": [
			[
				["string_1", "string_2"]
			],
			[
				["string_3", "string_4"]
			],
			[
				["string_5", "string_6"]
			]
		]
	  }
	`

	type Slices struct {
		TopSlice       []string       `njson:"top_slice"`
		IntMatrix      [][]int        `njson:"int_matrix"`
		Int3DMatrix    [][][]int      `njson:"int_3d_matrix"`
		StringMatrix   [][]string     `njson:"string_matrix"`
		String3DMatrix [][][]string   `njson:"string_3d_matrix"`
		String4DMatrix [][][][]string `njson:"string_4d_matrix"`
	}

	actual := Slices{}

	err := Unmarshal([]byte(json), &actual)
	if err != nil {
		t.Error(err)
	}

	expected := Slices{
		TopSlice:       []string{"foo", "bar"},
		IntMatrix:      [][]int{{101, 102, 103}, {201, 202, 203}, {301, 302, 303}},
		Int3DMatrix:    [][][]int{{{101, 102, 103}, {201, 202, 203}, {301, 302, 303}}, {{401, 402, 403}, {501, 502, 503}}, {{601, 602, 603}}},
		StringMatrix:   [][]string{{"string_1", "string_2"}, {"string_3", "string_4"}},
		String3DMatrix: [][][]string{{{"string_1", "string_2"}}, {{"string_3", "string_4"}}, {{"string_5", "string_6"}}},
		String4DMatrix: [][][][]string{},
	}

	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Error(diff)
	}
}

func TestUnmarshallMaps(t *testing.T) {

	json := `
	{
		"map1": {
			"key1": "value1",
			"key2": "value2"
		},
		"map2": {
			"key1": 1,
			"key2": 2,
			"key3": 3
		},
		"map3": {
			"key1": {
				"k1": 1,
				"k2": 2
			},
			"key2": {
				"k3": 3,
				"k4": 4
			}
		},
		"map4": {
			"key1": {
				"k1": "v1",
				"k2": "v2"
			},
			"key2": {
				"k1": "v3",
				"k2": "v4"
			},
			"key3": {
				"k1": "v5",
				"k2": "v6"
			}
		}
	}
	`

	type S1 struct {
		K1 string
		K2 string
	}

	type Maps struct {
		Map1 map[string]string         `njson:"map1"`
		Map2 map[string]int            `njson:"map2"`
		Map3 map[string]map[string]int `njson:"map3"`
		Map4 map[string]*S1            `njson:"map4"`
	}

	actual := Maps{}

	err := Unmarshal([]byte(json), &actual)
	if err != nil {
		t.Error(err)
	}

	expected := Maps{
		Map1: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
		Map2: map[string]int{
			"key1": 1,
			"key2": 2,
			"key3": 3,
		},
		Map3: map[string]map[string]int{
			"key1": {
				"k1": 1,
				"k2": 2,
			},
			"key2": {
				"k3": 3,
				"k4": 4,
			},
		},
		Map4: map[string]*S1{
			"key1": {
				K1: "v1",
				K2: "v2",
			},
			"key2": {
				K1: "v3",
				K2: "v4",
			},
			"key3": {
				K1: "v5",
				K2: "v6",
			},
		},
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
		],
		"time_1": "2021-01-11T23:56:51.141Z",
		"time_2": "2021-01-11T23:56:51.141+01:00",
		"time_3": "2021-01-11T23:56:51.141-01:00"
	  }
	`

	type User struct {
		FirstName              string    `njson:"name.first"`
		LastName               string    `njson:"name.last"`
		Age                    int       `njson:"age"`
		NumberOfChildren       int       `njson:"children.#"`
		Children               []string  `njson:"children"`
		Friends                []string  `njson:"friends.#.last"`
		NumberOfFriendNetworks []int     `njson:"friends.#.nets.#"`
		Time1                  time.Time `njson:"time_1"`
		Time2                  time.Time `njson:"time_2"`
		Time3                  time.Time `njson:"time_3"`
		//FavoriteMovie    string   `njson:"fav.movie"`
	}

	actual := User{}

	err := Unmarshal([]byte(json), &actual)
	if err != nil {
		t.Error(err)
	}

	t1, _ := time.Parse(time.RFC3339, "2021-01-11T23:56:51.141Z")
	t2, _ := time.Parse(time.RFC3339, "2021-01-11T23:56:51.141+01:00")
	t3, _ := time.Parse(time.RFC3339, "2021-01-11T23:56:51.141-01:00")

	expected := User{
		FirstName:              "Tom",
		LastName:               "Anderson",
		Age:                    37,
		NumberOfChildren:       3,
		Children:               []string{"Sara", "Alex", "Jack"},
		Friends:                []string{"Murphy", "Craig", "Murphy"},
		NumberOfFriendNetworks: []int{3, 2, 2},
		Time1:                  t1,
		Time2:                  t2,
		Time3:                  t3,
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
		],
		"list_of_structs": [
		  {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
		  {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
		  {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
		]
	  }
	`

	type MoreComplex struct {
		FirstChild          string     `njson:"top_struct.child_struct.first_child"`
		FirstChildOfChild   string     `njson:"top_struct.child_struct.another_child_struct.first_child_of_first_child"`
		ListInStruct        []int      `njson:"top_struct.a_list"`
		AChildListInStruct  []string   `njson:"top_struct.child_struct.a_child_list"`
		ListOfLists         [][]int    `njson:"list_of_lists"`
		ListOfStructs       []string   `njson:"list_of_structs"`
		ListOfNetsInStructs []string   `njson:"list_of_structs.#.nets"`
		ListOfNetLists      [][]string `njson:"list_of_structs.#.nets"`
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
		ListOfLists:        [][]int{{101, 102, 103}, {201, 202, 203}, {301, 302, 303}},
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
		ListOfNetLists: [][]string{{"ig", "fb", "tw"}, {"fb", "tw"}, {"ig", "tw"}},
	}

	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Error(diff)
	}

}

func TestUnmarshalJson(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		json := `
		{
			"name": {"first": "Mohamed", "last": "Shapan"},
			"age": 26,
			"friends": [
				{"first": "Asma", "age": 26},
				{"first": "Ahmed", "age": 25},
				{"first": "Mahmoud", "age": 30}
			]
		}`

		type Name struct {
			First string `njson:"first"`
			Last  string `njson:"last"`
		}

		type User struct {
			Name    Name   `njson:"name"`
			Age     int    `json:"age"`
			Friends []Name `json:"friends"`
		}

		actual := User{}

		err := Unmarshal([]byte(json), &actual)
		if err != nil {
			t.Error(err)
		}

		var friends []Name
		friends = append(friends, Name{
			First: "Asma",
		})

		friends = append(friends, Name{
			First: "Ahmed",
		})

		friends = append(friends, Name{
			First: "Mahmoud",
		})

		expected := User{
			Name: Name{
				First: "Mohamed",
				Last:  "Shapan",
			},
			Age:     26,
			Friends: friends,
		}

		diff := cmp.Diff(expected, actual)
		if diff != "" {
			t.Error(diff)
		}
	})

	t.Run("fail", func(t *testing.T) {
		json := `
		{
			"name": {"first": "Mohamed", "last": "Shapan"},
			"age": 26,
			"friends": [
				{"first": "Asma", "age": 26},
				{"first": "Ahmed", "age": 25},
				{"first": "Mahmoud", "age": 30}
			]
		}`

		type Name struct {
			First string `njson:"first"`
			Last  string `njson:"last"`
		}

		type User struct {
			Name    string `json:"name.first"`
			Age     int    `json:"age"`
			Friends []Name `json:"friends"`
		}

		actual := User{}

		if err := Unmarshal([]byte(json), &actual); err == nil {
			t.Error("error should not be nil")
		}
	})
}
