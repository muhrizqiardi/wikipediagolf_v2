package testutil

import (
	"errors"
	"io"
	"net/url"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/schema"
)

func ToFormUrlencoded(src interface{}) (io.Reader, error) {
	dst := url.Values{}
	if err := schema.NewEncoder().Encode(src, dst); err != nil {
		return nil, err
	}
	result := dst.Encode()
	return strings.NewReader(result), nil
}

func MustToFormUrlencoded(src interface{}) io.Reader {
	b, err := ToFormUrlencoded(src)
	if err != nil {
		return nil
	}

	return b
}

func CompareError(t *testing.T, exp, got error) bool {
	t.Helper()
	result := errors.Is(exp, got)
	if !result {
		t.Errorf("exp: %v; got: %v", exp, got)
	}
	return result
}

func AssertError(t *testing.T, err error) bool {
	t.Helper()
	if err == nil {
		t.Error("exp error; got nil")
		return false
	}

	return true
}

func AssertNoError(t *testing.T, err error) bool {
	t.Helper()
	if err != nil {
		t.Error("exp nil; got error:", err)
		return false
	}

	return true
}

func AssertEqual[T comparable](t *testing.T, exp, got T) {
	t.Helper()
	if exp != got {
		t.Errorf("exp %v; got %v", exp, got)
	}
}

func AssertInequal[T comparable](t *testing.T, exp, got T) {
	t.Helper()
	if exp == got {
		t.Errorf("exp got != %v; got %v", exp, got)
	}
}

func AssertNotNil(t *testing.T, got any) {
	t.Helper()
	if got == nil {
		t.Errorf("exp not nil; got %v", got)
	}
}

func AssertNil(t *testing.T, got any) {
	t.Helper()
	if got == nil {
		t.Errorf("exp not nil; got %v", got)
	}
}

func AssertEqualCMP(t *testing.T, exp, got interface{}, opts ...cmp.Option) bool {
	t.Helper()
	if diff := cmp.Diff(exp, got, opts...); diff != "" {
		t.Errorf("not equal (-exp +got):\n%s", diff)
		return false
	}

	return true
}

func AssertNotEqualCMP(t *testing.T, exp, got interface{}, opts ...cmp.Option) bool {
	t.Helper()
	if diff := cmp.Diff(exp, got, opts...); diff == "" {
		t.Error("expected not equal, got equal")
		return false
	}

	return true
}
