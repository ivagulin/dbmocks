package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite008 struct {
	SuiteBase
}

func (s *Suite008) TestFunc1()  {}
func (s *Suite008) TestFunc2()  {}
func (s *Suite008) TestFunc3()  {}
func (s *Suite008) TestFunc4()  {}
func (s *Suite008) TestFunc5()  {}
func (s *Suite008) TestFunc6()  {}
func (s *Suite008) TestFunc7()  {}
func (s *Suite008) TestFunc8()  {}
func (s *Suite008) TestFunc9()  {}
func (s *Suite008) TestFunc10() {}

func TestSuite008(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite008))
}
