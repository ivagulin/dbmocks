package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite017 struct {
	SuiteBase
}

func (s *Suite017) TestFunc1()  {}
func (s *Suite017) TestFunc2()  {}
func (s *Suite017) TestFunc3()  {}
func (s *Suite017) TestFunc4()  {}
func (s *Suite017) TestFunc5()  {}
func (s *Suite017) TestFunc6()  {}
func (s *Suite017) TestFunc7()  {}
func (s *Suite017) TestFunc8()  {}
func (s *Suite017) TestFunc9()  {}
func (s *Suite017) TestFunc10() {}

func TestSuite017(t *testing.T) {
	suite.Run(t, new(Suite017))
}
