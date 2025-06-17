package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite089 struct {
	SuiteBase
}

func (s *Suite089) TestFunc1()  {}
func (s *Suite089) TestFunc2()  {}
func (s *Suite089) TestFunc3()  {}
func (s *Suite089) TestFunc4()  {}
func (s *Suite089) TestFunc5()  {}
func (s *Suite089) TestFunc6()  {}
func (s *Suite089) TestFunc7()  {}
func (s *Suite089) TestFunc8()  {}
func (s *Suite089) TestFunc9()  {}
func (s *Suite089) TestFunc10() {}

func TestSuite089(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite089))
}
