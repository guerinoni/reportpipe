package tests

import (
	"encoding/json"
	"net/http"
	"reportpipe/internal"
	"strings"
	"testing"
)

func (s *IntegrationSuite) TestSignInWithInvalidBody() {
	cases := []string{
		`{}`,
		`{"email": ""}`,
		`{"email": "", "password": ""}`,
		`{"email": "", "password": "", "name": ""}`,

		`{"email": "me@guerra.io", "password": "", "name": ""}`,
		`{"email": "", "password": "password", "name": ""}`,
		`{"email": "", "password": "", "name": "Guerra"}`,

		`{"email": "me@guerra.io", "password": "pw", "name": ""}`,
		`{"email": "me@guerra.io", "password": "", "name": "guerra"}`,
		`{"email": "", "password": "password", "name": "Guerra"}`,
	}

	for _, body := range cases {
		s.T().Run(body, func(t *testing.T) {
			r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/signin", strings.NewReader(body))
			s.Require().NoError(err)
			resp, err := http.DefaultClient.Do(r)
			s.Require().NoError(err)
			s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
		})
	}
}

func (s *IntegrationSuite) TestSignInWithValidBody() {
	body := `{"email": "me@guerra.io", "password": "password", "name": "Guerra"}`
	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/signin", strings.NewReader(body))
	s.Require().NoError(err)

	resp, err := http.DefaultClient.Do(r)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusCreated, resp.StatusCode)

	var v internal.SignInResponse
	s.NoError(json.NewDecoder(resp.Body).Decode(&v))
	s.NotEmpty(v.Token)
	s.NotEmpty(v.Email)
	s.NotEmpty(v.Name)
}
