package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite041 struct {
	SuiteBase
}

func (s *Suite041) TestFunc1()  {}
func (s *Suite041) TestFunc2()  {}
func (s *Suite041) TestFunc3()  {}
func (s *Suite041) TestFunc4()  {}
func (s *Suite041) TestFunc5()  {}
func (s *Suite041) TestFunc6()  {}
func (s *Suite041) TestFunc7()  {}
func (s *Suite041) TestFunc8()  {}
func (s *Suite041) TestFunc9()  {}
func (s *Suite041) TestFunc10() {}

func TestSuite041(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite041))
}
