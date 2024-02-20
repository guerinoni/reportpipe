package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"reportpipe/internal"

	"github.com/stretchr/testify/suite"
)

type IntegrationSuite struct {
	suite.Suite

	cancelCtx context.CancelFunc
}

func Test(t *testing.T) {
	suite.Run(t, new(IntegrationSuite))
}

func (s *IntegrationSuite) SetupTest() {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancelCtx = cancel

	go func() {
		err := internal.Run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr)
		s.Require().NoError(err)
	}()

	for {
		r, err := http.NewRequest("GET", "http://127.0.0.1:8080/health", nil)
		s.Require().NoError(err)

		s.T().Log("Waiting for server to be ready")
		resp, err := http.DefaultClient.Do(r)
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()
			break
		}

		time.Sleep(250 * time.Millisecond)
	}
}

func (s *IntegrationSuite) TearDownTest() {
	s.cancelCtx()

	for {
		r, err := http.NewRequest("GET", "http://127.0.0.1:8080/health", nil)
		s.Require().NoError(err)

		resp, err := http.DefaultClient.Do(r)
		if err != nil || resp.StatusCode != http.StatusOK {
			break
		}

		resp.Body.Close()

		time.Sleep(250 * time.Millisecond)
	}
}

var defaultLettersForRandom = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// NewRandom generates a random string in the given length n from the given runes
func NewRandom(n int, from ...rune) string {
	if len(from) == 0 {
		from = defaultLettersForRandom
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = from[rand.Intn(len(from))]
	}
	return string(b)
}

// signupUser is a helper function to create a user.
func (s *IntegrationSuite) signupUser(username string, email string) (token string) {
	s.T().Helper()

	body := fmt.Sprintf(`{"email": "%s", "password": "password", "username": "%s"}`, email, username)
	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/signup", strings.NewReader(body))
	s.Require().NoError(err)

	resp, err := http.DefaultClient.Do(r)
	s.Require().NoError(err)
	s.Require().Equal(http.StatusCreated, resp.StatusCode)

	defer resp.Body.Close()

	var v internal.SignUpResponse
	s.NoError(json.NewDecoder(resp.Body).Decode(&v))
	s.NotEmpty(v.Token)
	s.NotEmpty(v.Email)
	s.NotEmpty(v.Username)

	token = v.Token
	return
}
