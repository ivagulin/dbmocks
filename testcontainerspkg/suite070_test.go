package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite070 struct {
	SuiteBase
}

func (s *Suite070) TestFunc1()  {}
func (s *Suite070) TestFunc2()  {}
func (s *Suite070) TestFunc3()  {}
func (s *Suite070) TestFunc4()  {}
func (s *Suite070) TestFunc5()  {}
func (s *Suite070) TestFunc6()  {}
func (s *Suite070) TestFunc7()  {}
func (s *Suite070) TestFunc8()  {}
func (s *Suite070) TestFunc9()  {}
func (s *Suite070) TestFunc10() {}

func TestSuite070(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite070))
}
