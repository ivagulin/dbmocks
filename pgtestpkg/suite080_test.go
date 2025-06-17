package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite080 struct {
	SuiteBase
}

func (s *Suite080) TestFunc1()  {}
func (s *Suite080) TestFunc2()  {}
func (s *Suite080) TestFunc3()  {}
func (s *Suite080) TestFunc4()  {}
func (s *Suite080) TestFunc5()  {}
func (s *Suite080) TestFunc6()  {}
func (s *Suite080) TestFunc7()  {}
func (s *Suite080) TestFunc8()  {}
func (s *Suite080) TestFunc9()  {}
func (s *Suite080) TestFunc10() {}

func TestSuite080(t *testing.T) {
	suite.Run(t, new(Suite080))
}
