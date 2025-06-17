package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite058 struct {
	SuiteBase
}

func (s *Suite058) TestFunc1()  {}
func (s *Suite058) TestFunc2()  {}
func (s *Suite058) TestFunc3()  {}
func (s *Suite058) TestFunc4()  {}
func (s *Suite058) TestFunc5()  {}
func (s *Suite058) TestFunc6()  {}
func (s *Suite058) TestFunc7()  {}
func (s *Suite058) TestFunc8()  {}
func (s *Suite058) TestFunc9()  {}
func (s *Suite058) TestFunc10() {}

func TestSuite058(t *testing.T) {
	suite.Run(t, new(Suite058))
}
