package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite037 struct {
	SuiteBase
}

func (s *Suite037) TestFunc1()  {}
func (s *Suite037) TestFunc2()  {}
func (s *Suite037) TestFunc3()  {}
func (s *Suite037) TestFunc4()  {}
func (s *Suite037) TestFunc5()  {}
func (s *Suite037) TestFunc6()  {}
func (s *Suite037) TestFunc7()  {}
func (s *Suite037) TestFunc8()  {}
func (s *Suite037) TestFunc9()  {}
func (s *Suite037) TestFunc10() {}

func TestSuite037(t *testing.T) {
	suite.Run(t, new(Suite037))
}
