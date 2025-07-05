package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite078 struct {
	SuiteBase
}

func (s *Suite078) TestFunc1()  {}
func (s *Suite078) TestFunc2()  {}
func (s *Suite078) TestFunc3()  {}
func (s *Suite078) TestFunc4()  {}
func (s *Suite078) TestFunc5()  {}
func (s *Suite078) TestFunc6()  {}
func (s *Suite078) TestFunc7()  {}
func (s *Suite078) TestFunc8()  {}
func (s *Suite078) TestFunc9()  {}
func (s *Suite078) TestFunc10() {}

func TestSuite078(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite078))
}
