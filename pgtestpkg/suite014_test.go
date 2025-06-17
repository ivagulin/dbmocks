package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite014 struct {
	SuiteBase
}

func (s *Suite014) TestFunc1()  {}
func (s *Suite014) TestFunc2()  {}
func (s *Suite014) TestFunc3()  {}
func (s *Suite014) TestFunc4()  {}
func (s *Suite014) TestFunc5()  {}
func (s *Suite014) TestFunc6()  {}
func (s *Suite014) TestFunc7()  {}
func (s *Suite014) TestFunc8()  {}
func (s *Suite014) TestFunc9()  {}
func (s *Suite014) TestFunc10() {}

func TestSuite014(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite014))
}
