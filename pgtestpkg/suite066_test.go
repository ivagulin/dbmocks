package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite066 struct {
	SuiteBase
}

func (s *Suite066) TestFunc1()  {}
func (s *Suite066) TestFunc2()  {}
func (s *Suite066) TestFunc3()  {}
func (s *Suite066) TestFunc4()  {}
func (s *Suite066) TestFunc5()  {}
func (s *Suite066) TestFunc6()  {}
func (s *Suite066) TestFunc7()  {}
func (s *Suite066) TestFunc8()  {}
func (s *Suite066) TestFunc9()  {}
func (s *Suite066) TestFunc10() {}

func TestSuite066(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite066))
}
