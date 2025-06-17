package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite098 struct {
	SuiteBase
}

func (s *Suite098) TestFunc1()  {}
func (s *Suite098) TestFunc2()  {}
func (s *Suite098) TestFunc3()  {}
func (s *Suite098) TestFunc4()  {}
func (s *Suite098) TestFunc5()  {}
func (s *Suite098) TestFunc6()  {}
func (s *Suite098) TestFunc7()  {}
func (s *Suite098) TestFunc8()  {}
func (s *Suite098) TestFunc9()  {}
func (s *Suite098) TestFunc10() {}

func TestSuite098(t *testing.T) {
	suite.Run(t, new(Suite098))
}
