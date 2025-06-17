package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite074 struct {
	SuiteBase
}

func (s *Suite074) TestFunc1()  {}
func (s *Suite074) TestFunc2()  {}
func (s *Suite074) TestFunc3()  {}
func (s *Suite074) TestFunc4()  {}
func (s *Suite074) TestFunc5()  {}
func (s *Suite074) TestFunc6()  {}
func (s *Suite074) TestFunc7()  {}
func (s *Suite074) TestFunc8()  {}
func (s *Suite074) TestFunc9()  {}
func (s *Suite074) TestFunc10() {}

func TestSuite074(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite074))
}
