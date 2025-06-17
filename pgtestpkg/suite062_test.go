package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite062 struct {
	SuiteBase
}

func (s *Suite062) TestFunc1()  {}
func (s *Suite062) TestFunc2()  {}
func (s *Suite062) TestFunc3()  {}
func (s *Suite062) TestFunc4()  {}
func (s *Suite062) TestFunc5()  {}
func (s *Suite062) TestFunc6()  {}
func (s *Suite062) TestFunc7()  {}
func (s *Suite062) TestFunc8()  {}
func (s *Suite062) TestFunc9()  {}
func (s *Suite062) TestFunc10() {}

func TestSuite062(t *testing.T) {
	suite.Run(t, new(Suite062))
}
