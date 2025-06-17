package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite063 struct {
	SuiteBase
}

func (s *Suite063) TestFunc1()  {}
func (s *Suite063) TestFunc2()  {}
func (s *Suite063) TestFunc3()  {}
func (s *Suite063) TestFunc4()  {}
func (s *Suite063) TestFunc5()  {}
func (s *Suite063) TestFunc6()  {}
func (s *Suite063) TestFunc7()  {}
func (s *Suite063) TestFunc8()  {}
func (s *Suite063) TestFunc9()  {}
func (s *Suite063) TestFunc10() {}

func TestSuite063(t *testing.T) {
	suite.Run(t, new(Suite063))
}
