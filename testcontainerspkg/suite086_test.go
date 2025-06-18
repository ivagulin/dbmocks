package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite086 struct {
	SuiteBase
}

func (s *Suite086) TestFunc1()  {}
func (s *Suite086) TestFunc2()  {}
func (s *Suite086) TestFunc3()  {}
func (s *Suite086) TestFunc4()  {}
func (s *Suite086) TestFunc5()  {}
func (s *Suite086) TestFunc6()  {}
func (s *Suite086) TestFunc7()  {}
func (s *Suite086) TestFunc8()  {}
func (s *Suite086) TestFunc9()  {}
func (s *Suite086) TestFunc10() {}

func TestSuite086(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite086))
}
