package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite1 struct {
	SuiteBase
}

func (s *Suite1) TestFunc1()  {}
func (s *Suite1) TestFunc2()  {}
func (s *Suite1) TestFunc3()  {}
func (s *Suite1) TestFunc4()  {}
func (s *Suite1) TestFunc5()  {}
func (s *Suite1) TestFunc6()  {}
func (s *Suite1) TestFunc7()  {}
func (s *Suite1) TestFunc8()  {}
func (s *Suite1) TestFunc9()  {}
func (s *Suite1) TestFunc10() {}

func TestSuite1(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite1))
}
