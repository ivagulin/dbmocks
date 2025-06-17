package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite085 struct {
	SuiteBase
}

func (s *Suite085) TestFunc1()  {}
func (s *Suite085) TestFunc2()  {}
func (s *Suite085) TestFunc3()  {}
func (s *Suite085) TestFunc4()  {}
func (s *Suite085) TestFunc5()  {}
func (s *Suite085) TestFunc6()  {}
func (s *Suite085) TestFunc7()  {}
func (s *Suite085) TestFunc8()  {}
func (s *Suite085) TestFunc9()  {}
func (s *Suite085) TestFunc10() {}

func TestSuite085(t *testing.T) {
	suite.Run(t, new(Suite085))
}
