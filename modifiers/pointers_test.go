package modifiers_test

import (
	"context"
	"testing"

	"github.com/takt-corp/mold/modifiers"
)

func TestNilEmptyTypeWithEmptyValue(t *testing.T) {
	emptyString := ""
	testValue := &emptyString

	conform := modifiers.New()
	err := conform.Field(context.Background(), &testValue, "nil_empty")

	if err != nil {
		t.Errorf("there was an error while conforming nil_empty: %v", err)
		return
	}

	if testValue != nil {
		t.Errorf("test value should be nil")
		return
	}
}

func TestNilEmptyTypeWithNonEmptyValue(t *testing.T) {
	nonEmptyString := "hello world"
	testValue := &nonEmptyString

	conform := modifiers.New()
	err := conform.Field(context.Background(), &testValue, "nil_empty")

	if err != nil {
		t.Errorf("there was an error while conforming nil_empty: %v", err)
		return
	}

	if testValue == nil {
		t.Errorf("test value should be nil")
		return
	}

	if *testValue != nonEmptyString {
		t.Errorf("test value did not equal original string")
		return
	}
}
