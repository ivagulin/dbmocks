package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite046 struct {
	SuiteBase
}

func (s *Suite046) TestFunc1()  {}
func (s *Suite046) TestFunc2()  {}
func (s *Suite046) TestFunc3()  {}
func (s *Suite046) TestFunc4()  {}
func (s *Suite046) TestFunc5()  {}
func (s *Suite046) TestFunc6()  {}
func (s *Suite046) TestFunc7()  {}
func (s *Suite046) TestFunc8()  {}
func (s *Suite046) TestFunc9()  {}
func (s *Suite046) TestFunc10() {}

func TestSuite046(t *testing.T) {
	suite.Run(t, new(Suite046))
}
