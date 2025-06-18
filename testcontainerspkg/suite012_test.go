package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite012 struct {
	SuiteBase
}

func (s *Suite012) TestFunc1()  {}
func (s *Suite012) TestFunc2()  {}
func (s *Suite012) TestFunc3()  {}
func (s *Suite012) TestFunc4()  {}
func (s *Suite012) TestFunc5()  {}
func (s *Suite012) TestFunc6()  {}
func (s *Suite012) TestFunc7()  {}
func (s *Suite012) TestFunc8()  {}
func (s *Suite012) TestFunc9()  {}
func (s *Suite012) TestFunc10() {}

func TestSuite012(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite012))
}
