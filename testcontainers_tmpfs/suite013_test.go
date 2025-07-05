package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite013 struct {
	SuiteBase
}

func (s *Suite013) TestFunc1()  {}
func (s *Suite013) TestFunc2()  {}
func (s *Suite013) TestFunc3()  {}
func (s *Suite013) TestFunc4()  {}
func (s *Suite013) TestFunc5()  {}
func (s *Suite013) TestFunc6()  {}
func (s *Suite013) TestFunc7()  {}
func (s *Suite013) TestFunc8()  {}
func (s *Suite013) TestFunc9()  {}
func (s *Suite013) TestFunc10() {}

func TestSuite013(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite013))
}
