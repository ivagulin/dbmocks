package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite018 struct {
	SuiteBase
}

func (s *Suite018) TestFunc1()  {}
func (s *Suite018) TestFunc2()  {}
func (s *Suite018) TestFunc3()  {}
func (s *Suite018) TestFunc4()  {}
func (s *Suite018) TestFunc5()  {}
func (s *Suite018) TestFunc6()  {}
func (s *Suite018) TestFunc7()  {}
func (s *Suite018) TestFunc8()  {}
func (s *Suite018) TestFunc9()  {}
func (s *Suite018) TestFunc10() {}

func TestSuite018(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite018))
}
