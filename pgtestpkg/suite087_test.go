package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite087 struct {
	SuiteBase
}

func (s *Suite087) TestFunc1()  {}
func (s *Suite087) TestFunc2()  {}
func (s *Suite087) TestFunc3()  {}
func (s *Suite087) TestFunc4()  {}
func (s *Suite087) TestFunc5()  {}
func (s *Suite087) TestFunc6()  {}
func (s *Suite087) TestFunc7()  {}
func (s *Suite087) TestFunc8()  {}
func (s *Suite087) TestFunc9()  {}
func (s *Suite087) TestFunc10() {}

func TestSuite087(t *testing.T) {
	suite.Run(t, new(Suite087))
}
