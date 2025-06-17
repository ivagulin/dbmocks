package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite073 struct {
	SuiteBase
}

func (s *Suite073) TestFunc1()  {}
func (s *Suite073) TestFunc2()  {}
func (s *Suite073) TestFunc3()  {}
func (s *Suite073) TestFunc4()  {}
func (s *Suite073) TestFunc5()  {}
func (s *Suite073) TestFunc6()  {}
func (s *Suite073) TestFunc7()  {}
func (s *Suite073) TestFunc8()  {}
func (s *Suite073) TestFunc9()  {}
func (s *Suite073) TestFunc10() {}

func TestSuite073(t *testing.T) {
	suite.Run(t, new(Suite073))
}
