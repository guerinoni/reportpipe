package tests

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (s *IntegrationSuite) TestLoginWithInvalidBody() {
	cases := []string{
		`{}`,

		`{"email": ""}`,
		`{"email": "me@guerra.io"}`,
		`{"password": ""}`,
		`{"password": "password"}`,

		`{"email": "", "password": ""}`,
		`{"email": "me@guerra.io", "password": ""}`,
		`{"email": "", "password": "password"}`,
	}

	for _, body := range cases {
		r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/login", strings.NewReader(body))
		s.Require().NoError(err)
		resp, err := http.DefaultClient.Do(r)
		s.Require().NoError(err)
		s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	}
}

func (s *IntegrationSuite) TestLoginWithValidCredentialsButNoUserFound() {
	email := fmt.Sprintf("%s@%s.io", NewRandom(5), NewRandom(5))
	body := fmt.Sprintf(`{"email": "%s", "password": "password"}`, email)
	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/login", strings.NewReader(body))
	s.Require().NoError(err)

	resp, err := http.DefaultClient.Do(r)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusUnauthorized, resp.StatusCode)
	b, err := io.ReadAll(resp.Body)
	s.Require().NoError(err)
	s.Require().JSONEq(`{"errors":{"user":"not found"}}`, string(b))
}
