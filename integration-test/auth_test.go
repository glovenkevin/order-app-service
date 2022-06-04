package integration_test

import (
	"net/http"
	"testing"

	. "github.com/Eun/go-hit"
)

func TestHTTPSuccessRegister(t *testing.T) {
	body := `{
		"name": "test",
		"email": "test@test.com",
		"password": "asdf",
		"phone_number": "08123456789",
		"is_blocked": false,
		"fcm_token": "asdf123"
	}`

	Test(t,
		Description("DoNormalRegister"),
		Post(basePathApi+"/auth/register"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusOK),
	)
}

func TestHTTPIncompleteParamRegister(t *testing.T) {
	body := `{
		"name": "test",
		"password": "asdf",
		"phone_number": "08123456789",
		"is_blocked": false
	}`

	Test(t,
		Description("DoNormalRegister"),
		Post(basePathApi+"/auth/register"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".message").Contains("Error:Field validation"),
	)
}

func TestHTTPDuplicateEmail(t *testing.T) {
	body := `{
		"name": "test",
		"email": "test@test.com",
		"password": "asdf",
		"phone_number": "08123456789",
		"is_blocked": false,
		"fcm_token": "asdf123"
	}`

	Test(t,
		Description("DoNormalRegister with duplicate email"),
		Post(basePathApi+"/auth/register"),
		Send().Headers("Content-Type").Add("application/json"),
		Send().Body().String(body),
		Expect().Status().Equal(http.StatusBadRequest),
		Expect().Body().JSON().JQ(".message").Contains("Email already exists"),
	)
}
