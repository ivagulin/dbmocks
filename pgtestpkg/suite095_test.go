package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite095 struct {
	SuiteBase
}

func (s *Suite095) TestFunc1()  {}
func (s *Suite095) TestFunc2()  {}
func (s *Suite095) TestFunc3()  {}
func (s *Suite095) TestFunc4()  {}
func (s *Suite095) TestFunc5()  {}
func (s *Suite095) TestFunc6()  {}
func (s *Suite095) TestFunc7()  {}
func (s *Suite095) TestFunc8()  {}
func (s *Suite095) TestFunc9()  {}
func (s *Suite095) TestFunc10() {}

func TestSuite095(t *testing.T) {
	suite.Run(t, new(Suite095))
}
