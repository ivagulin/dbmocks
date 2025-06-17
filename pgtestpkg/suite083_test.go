package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite083 struct {
	SuiteBase
}

func (s *Suite083) TestFunc1()  {}
func (s *Suite083) TestFunc2()  {}
func (s *Suite083) TestFunc3()  {}
func (s *Suite083) TestFunc4()  {}
func (s *Suite083) TestFunc5()  {}
func (s *Suite083) TestFunc6()  {}
func (s *Suite083) TestFunc7()  {}
func (s *Suite083) TestFunc8()  {}
func (s *Suite083) TestFunc9()  {}
func (s *Suite083) TestFunc10() {}

func TestSuite083(t *testing.T) {
	suite.Run(t, new(Suite083))
}
