package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite047 struct {
	SuiteBase
}

func (s *Suite047) TestFunc1()  {}
func (s *Suite047) TestFunc2()  {}
func (s *Suite047) TestFunc3()  {}
func (s *Suite047) TestFunc4()  {}
func (s *Suite047) TestFunc5()  {}
func (s *Suite047) TestFunc6()  {}
func (s *Suite047) TestFunc7()  {}
func (s *Suite047) TestFunc8()  {}
func (s *Suite047) TestFunc9()  {}
func (s *Suite047) TestFunc10() {}

func TestSuite047(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite047))
}
