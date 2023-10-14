package session

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alexedwards/scs/v2"
)

func TestSession_InitSession(t *testing.T) {
	c := &Session{
		CookieLifetime: "100",
		CookiePersist:  "true",
		CookieName:     "brimstoneesan",
		CookieDomain:   "localhost",
		SessionType:    "cookie",
	}

	var sm *scs.SessionManager

	ses := c.InitSession()

	var sessKind reflect.Kind
	var sessType reflect.Type

	reflectValue := reflect.ValueOf(ses)

	for reflectValue.Kind() == reflect.Ptr || reflectValue.Kind() == reflect.Interface {
		fmt.Println("For Loop: ", reflectValue.Kind(), reflectValue.Type(), reflectValue)
		sessKind = reflectValue.Kind()
		sessType = reflectValue.Type()

		reflectValue = reflectValue.Elem()
	}

	if !reflectValue.IsValid() {
		t.Error("invalid type of kind; kind:", reflectValue.Kind(), ", type : ", reflectValue.Type())
	}

	if sessKind != reflect.ValueOf(sm).Kind() {
		t.Error("wrong kind returned. Expected : ", reflect.ValueOf(sm).Kind(), ", got : ", sessKind)
	}

	if sessType != reflect.ValueOf(sm).Type() {
		t.Error("wrong type returned. Expected : ", reflect.ValueOf(sm).Type(), ", got : ", sessType)
	}
}
