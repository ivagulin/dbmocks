package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite020 struct {
	SuiteBase
}

func (s *Suite020) TestFunc1()  {}
func (s *Suite020) TestFunc2()  {}
func (s *Suite020) TestFunc3()  {}
func (s *Suite020) TestFunc4()  {}
func (s *Suite020) TestFunc5()  {}
func (s *Suite020) TestFunc6()  {}
func (s *Suite020) TestFunc7()  {}
func (s *Suite020) TestFunc8()  {}
func (s *Suite020) TestFunc9()  {}
func (s *Suite020) TestFunc10() {}

func TestSuite020(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite020))
}
