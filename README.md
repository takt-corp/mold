# Mold

This project has been forked from [go-playground/mold](https://github.com/go-playground/mold) and customized with additional validation functions. Mold is a package that helps modify, clean, and default struct values.

## Modifiers
These functions modify the data in-place.

| Name  | Description  |
|-------|--------------|
| default | Sets the provided default value only if the data is equal to it's default datatype value. |
| trim | Trims space from the data. |
| ltrim | Trims spaces from the left of the data provided in the params. |
| rtrim | Trims spaces from the right of the data provided in the params. |
| tprefix | Trims a prefix from the value using the provided param value. |
| tsuffix | Trims a suffix from the value using the provided param value. |
| lcase | lowercases the data. |
| ucase | Uppercases the data. |
| snake | Snake Cases the data. |
| camel | Camel Cases the data. |
| title | Title Cases the data. |
| ucfirst | Upper cases the first character of the data. |
| strip_alpha | Strips all ascii characters from the data. |
| strip_num | Strips all ascii numeric characters from the data. |
| strip_alpha_unicode | Strips all unicode characters from the data. |
| strip_num_unicode | Strips all unicode numeric characters from the data. |
| strip_punctuation | Strips all ascii punctuation from the data. |
| nil_empty | Sets the value of an empty (defined as the zero value) to nil. |


## Scrubbers
These functions obfuscate the specified types within the data for pii purposes.

| Name  | Description  |
|-------|--------------|
| emails | Scrubs multiple emails from data. |
| email | Scrubs the data from and specifies the sha name of the same name. |
| text | Scrubs the data from and specifies the sha name of the same name. |
| name | Scrubs the data from and specifies the sha name of the same name. |
| fname | Scrubs the data from and specifies the sha name of the same name. |
| lname | Scrubs the data from and specifies the sha name of the same name. |

## Usage

### Installation
Use go get.

```shell
go get -u github.com/takt-corp/mold
```

### Example

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"
	"github.com/takt-corp/mold/modifiers"
	"github.com/takt-corp/mold/scrubbers"
)

// This example is centered around a form post, but doesn't have to be
// just trying to give a well rounded real life example.

// <form method="POST">
//   <input type="text" name="Name" value="joeybloggs"/>
//   <input type="text" name="Age" value="3"/>
//   <input type="text" name="Gender" value="Male"/>
//   <input type="text" name="Address[0].Name" value="26 Here Blvd."/>
//   <input type="text" name="Address[0].Phone" value="9(999)999-9999"/>
//   <input type="text" name="Address[1].Name" value="26 There Blvd."/>
//   <input type="text" name="Address[1].Phone" value="1(111)111-1111"/>
//   <input type="text" name="active" value="true"/>
//   <input type="submit"/>
// </form>

var (
	conform  = modifiers.New()
	scrub    = scrubbers.New()
	validate = validator.New()
	decoder  = form.NewDecoder()
)

// Address contains address information
type Address struct {
	Name  string `mod:"trim" validate:"required"`
	Phone string `mod:"trim" validate:"required"`
}

// User contains user information
type User struct {
	Name    string            `mod:"trim"      validate:"required"              scrub:"name"`
	Age     uint8             `                validate:"required,gt=0,lt=130"`
	Gender  string            `                validate:"required"`
	Email   string            `mod:"trim"      validate:"required,email"        scrub:"emails"`
	Address []Address         `                validate:"required,dive"`
	Active  bool              `form:"active"`
	Misc    map[string]string `mod:"dive,keys,trim,endkeys,trim"`
}

func main() {
	// this simulates the results of http.Request's ParseForm() function
	values := parseForm()

	var user User

	// must pass a pointer
	err := decoder.Decode(&user, values)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Decoded:%+v\n\n", user)

	// great not lets conform our values, after all a human input the data
	// nobody's perfect
	err = conform.Struct(context.Background(), &user)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Conformed:%+v\n\n", user)

	// that's better all those extra spaces are gone
	// let's validate the data
	err = validate.Struct(user)
	if err != nil {
		log.Panic(err)
	}

	// ok now we know our data is good, let's do something with it like:
	// save to database
	// process request
	// etc....

	// ok now I'm done working with my data
	// let's log or store it somewhere
	// oh wait a minute, we have some sensitive PII data
	// let's make sure that's de-identified first
	err = scrub.Struct(context.Background(), &user)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Scrubbed:%+v\n\n", user)
}

// this simulates the results of http.Request's ParseForm() function
func parseForm() url.Values {
	return url.Values{
		"Name":             []string{"  joeybloggs  "},
		"Age":              []string{"3"},
		"Gender":           []string{"Male"},
		"Email":            []string{"Dean.Karn@gmail.com  "},
		"Address[0].Name":  []string{"26 Here Blvd."},
		"Address[0].Phone": []string{"9(999)999-9999"},
		"Address[1].Name":  []string{"26 There Blvd."},
		"Address[1].Phone": []string{"1(111)111-1111"},
		"active":           []string{"true"},
		"Misc[  b4  ]":     []string{"  b4  "},
	}
}

```

## Security

If you identity a security vulnerability or concern with this repository, reach out directly to [security@takt.io](mailto:security@takt.io) immediatley with the details.
