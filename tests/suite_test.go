package tests

import (
	"context"
	"github.com/stretchr/testify/suite"
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
