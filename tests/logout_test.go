package tests

import (
	"net/http"
	"strings"
)

func (s *IntegrationSuite) TestLogoutWithoutToken() {
	body := `{"token":"sometoken"}`
	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/logout", strings.NewReader(body))
	s.Require().NoError(err)
	resp, err := http.DefaultClient.Do(r)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusUnauthorized, resp.StatusCode)
}

func (s *IntegrationSuite) TestLogoutWithInvalidToken() {
	body := `{"token":"sometoken"}`
	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/logout", strings.NewReader(body))
	bearer := "Bearer sometoken"
	r.Header.Add("authorization", bearer)
	s.Require().NoError(err)
	resp, err := http.DefaultClient.Do(r)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusUnauthorized, resp.StatusCode)
}

func (s *IntegrationSuite) TestLogoutWithValidToken() {
	// signup user
	username := NewRandom(10)
	token := s.signupUser(username, username+"@example.com")

	// logout
	body := `{"token":"` + token + `"}`
	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/logout", strings.NewReader(body))
	bearer := "Bearer " + token
	r.Header.Add("authorization", bearer)
	s.Require().NoError(err)
	resp, err := http.DefaultClient.Do(r)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusOK, resp.StatusCode)

	// try to call an endpoint with the same token
	r, err = http.NewRequest(http.MethodPost, "http://localhost:8080/logout", strings.NewReader(body))
	bearer = "Bearer " + token
	r.Header.Add("authorization", bearer)
	s.Require().NoError(err)
	resp, err = http.DefaultClient.Do(r)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusUnauthorized, resp.StatusCode)
}
