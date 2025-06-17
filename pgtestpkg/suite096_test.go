package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite096 struct {
	SuiteBase
}

func (s *Suite096) TestFunc1()  {}
func (s *Suite096) TestFunc2()  {}
func (s *Suite096) TestFunc3()  {}
func (s *Suite096) TestFunc4()  {}
func (s *Suite096) TestFunc5()  {}
func (s *Suite096) TestFunc6()  {}
func (s *Suite096) TestFunc7()  {}
func (s *Suite096) TestFunc8()  {}
func (s *Suite096) TestFunc9()  {}
func (s *Suite096) TestFunc10() {}

func TestSuite096(t *testing.T) {
	suite.Run(t, new(Suite096))
}
