package scrubbers_test

import (
	"context"
	"testing"

	. "github.com/go-playground/assert/v2"
	"github.com/takt-corp/mold/scrubbers"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//

func TestEmails(t *testing.T) {
	scrub := scrubbers.New()
	email := "Takt.Engineering@takt.io"

	type Test struct {
		Email string `scrub:"emails"`
	}

	tt := Test{Email: email}
	err := scrub.Struct(context.Background(), &tt)
	Equal(t, err, nil)
	Equal(t, tt.Email, "<<scrubbed::email::sha1::16f0299fc5a801dc4f286fc1148d53bf2c28b306>>@takt.io")

	err = scrub.Field(context.Background(), &email, "emails")
	Equal(t, err, nil)
	Equal(t, email, "<<scrubbed::email::sha1::16f0299fc5a801dc4f286fc1148d53bf2c28b306>>@takt.io")

	var iface interface{}
	err = scrub.Field(context.Background(), &iface, "emails")
	Equal(t, err, nil)
	Equal(t, iface, nil)

	iface = "Takt.Engineering@takt.io"
	err = scrub.Field(context.Background(), &iface, "emails")
	Equal(t, err, nil)
	Equal(t, iface, "<<scrubbed::email::sha1::16f0299fc5a801dc4f286fc1148d53bf2c28b306>>@takt.io")
}

func TestText(t *testing.T) {
	scrub := scrubbers.New()
	name := "Takt Engineering"

	type Test struct {
		String string `scrub:"text"`
	}

	tt := Test{String: name}
	err := scrub.Struct(context.Background(), &tt)
	Equal(t, err, nil)
	Equal(t, tt.String, "<<scrubbed::text::sha1::5e084d3528084e014ddfece7c2e97c6c0ca2a660>>")

	err = scrub.Field(context.Background(), &name, "text")
	Equal(t, err, nil)
	Equal(t, name, "<<scrubbed::text::sha1::5e084d3528084e014ddfece7c2e97c6c0ca2a660>>")

	var iface interface{}
	err = scrub.Field(context.Background(), &iface, "text")
	Equal(t, err, nil)
	Equal(t, iface, nil)

	iface = "Takt Engineering"
	err = scrub.Field(context.Background(), &iface, "text")
	Equal(t, err, nil)
	Equal(t, iface, "<<scrubbed::text::sha1::5e084d3528084e014ddfece7c2e97c6c0ca2a660>>")

	// testing Text wrapped func
	name = "Takt Engineering"
	err = scrub.Field(context.Background(), &name, "name")
	Equal(t, err, nil)
	Equal(t, name, "<<scrubbed::name::sha1::5e084d3528084e014ddfece7c2e97c6c0ca2a660>>")
}
