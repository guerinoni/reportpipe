package tests

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func (s *IntegrationSuite) TestSignUpWithInvalidBody() {
	cases := []string{
		`{}`,
		`{"email": ""}`,
		`{"email": "", "password": ""}`,
		`{"email": "", "password": "", "username": ""}`,

		`{"email": "me@guerra.io", "password": "", "username": ""}`,
		`{"email": "", "password": "password", "username": ""}`,
		`{"email": "", "password": "", "username": "Guerra"}`,

		`{"email": "me@guerra.io", "password": "pw", "username": ""}`,
		`{"email": "me@guerra.io", "password": "", "username": "guerra"}`,
		`{"email": "", "password": "password", "username": "Guerra"}`,
	}

	for _, body := range cases {
		s.T().Run(body, func(t *testing.T) {
			r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/signup", strings.NewReader(body))
			s.Require().NoError(err)
			resp, err := http.DefaultClient.Do(r)
			s.Require().NoError(err)
			s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
		})
	}
}

func (s *IntegrationSuite) TestSignUpWithValidBody() {
	username := NewRandom(5)
	email := fmt.Sprintf("%s@%s.io", NewRandom(5), NewRandom(5))
	s.signupUser(email, username)
}
