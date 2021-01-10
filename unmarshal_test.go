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

func TestUnmarshal2(t *testing.T) {

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
