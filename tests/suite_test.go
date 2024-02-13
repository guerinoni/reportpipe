package tests

import (
	"context"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"net/http"
	"os"
	"reportpipe/internal"
	"testing"
	"time"
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
