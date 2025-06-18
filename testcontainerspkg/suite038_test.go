package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite038 struct {
	SuiteBase
}

func (s *Suite038) TestFunc1()  {}
func (s *Suite038) TestFunc2()  {}
func (s *Suite038) TestFunc3()  {}
func (s *Suite038) TestFunc4()  {}
func (s *Suite038) TestFunc5()  {}
func (s *Suite038) TestFunc6()  {}
func (s *Suite038) TestFunc7()  {}
func (s *Suite038) TestFunc8()  {}
func (s *Suite038) TestFunc9()  {}
func (s *Suite038) TestFunc10() {}

func TestSuite038(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite038))
}
