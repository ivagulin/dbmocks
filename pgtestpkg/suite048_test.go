package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite048 struct {
	SuiteBase
}

func (s *Suite048) TestFunc1()  {}
func (s *Suite048) TestFunc2()  {}
func (s *Suite048) TestFunc3()  {}
func (s *Suite048) TestFunc4()  {}
func (s *Suite048) TestFunc5()  {}
func (s *Suite048) TestFunc6()  {}
func (s *Suite048) TestFunc7()  {}
func (s *Suite048) TestFunc8()  {}
func (s *Suite048) TestFunc9()  {}
func (s *Suite048) TestFunc10() {}

func TestSuite048(t *testing.T) {
	suite.Run(t, new(Suite048))
}
