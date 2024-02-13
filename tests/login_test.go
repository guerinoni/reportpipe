package tests

import (
	"io"
	"net/http"
	"strings"
)

func (s *IntegrationSuite) TestLoginWithNoBody() {
	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/login", nil)
	s.Require().NoError(err)

	resp, err := http.DefaultClient.Do(r)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
}

func (s *IntegrationSuite) TestLoginWithInvalidEmail() {
	body := `{"email": "", "password": "password"}`
	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/login", strings.NewReader(body))
	s.Require().NoError(err)

	resp, err := http.DefaultClient.Do(r)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	b, err := io.ReadAll(resp.Body)
	s.Require().NoError(err)
	s.Require().JSONEq(`{"errors":{"email":"is required"}}`, string(b))
}

func (s *IntegrationSuite) TestLoginWithInvalidPassword() {
	body := `{"email": "me@guerra.io", "password": ""}`
	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/login", strings.NewReader(body))
	s.Require().NoError(err)

	resp, err := http.DefaultClient.Do(r)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	b, err := io.ReadAll(resp.Body)
	s.Require().NoError(err)
	s.Require().JSONEq(`{"errors":{"password":"is required"}}`, string(b))
}

func (s *IntegrationSuite) TestLoginWithEmptyEmailAndPassword() {
	body := `{"email": "", "password": ""}`
	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/login", strings.NewReader(body))
	s.Require().NoError(err)

	resp, err := http.DefaultClient.Do(r)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusBadRequest, resp.StatusCode)
	b, err := io.ReadAll(resp.Body)
	s.Require().NoError(err)
	s.Require().JSONEq(`{"errors":{"email":"is required","password":"is required"}}`, string(b))
}

func (s *IntegrationSuite) TestLoginWithValidCredentialsButNoUserFound() {
	body := `{"email": "me@guerra.io", "password": "password"}`
	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/login", strings.NewReader(body))
	s.Require().NoError(err)

	resp, err := http.DefaultClient.Do(r)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusUnauthorized, resp.StatusCode)
	b, err := io.ReadAll(resp.Body)
	s.Require().NoError(err)
	s.Require().JSONEq(`{"errors":{"user":"not found"}}`, string(b))
}
