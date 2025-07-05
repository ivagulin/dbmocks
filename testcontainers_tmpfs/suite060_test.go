package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite060 struct {
	SuiteBase
}

func (s *Suite060) TestFunc1()  {}
func (s *Suite060) TestFunc2()  {}
func (s *Suite060) TestFunc3()  {}
func (s *Suite060) TestFunc4()  {}
func (s *Suite060) TestFunc5()  {}
func (s *Suite060) TestFunc6()  {}
func (s *Suite060) TestFunc7()  {}
func (s *Suite060) TestFunc8()  {}
func (s *Suite060) TestFunc9()  {}
func (s *Suite060) TestFunc10() {}

func TestSuite060(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite060))
}
