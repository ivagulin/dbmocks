package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite016 struct {
	SuiteBase
}

func (s *Suite016) TestFunc1()  {}
func (s *Suite016) TestFunc2()  {}
func (s *Suite016) TestFunc3()  {}
func (s *Suite016) TestFunc4()  {}
func (s *Suite016) TestFunc5()  {}
func (s *Suite016) TestFunc6()  {}
func (s *Suite016) TestFunc7()  {}
func (s *Suite016) TestFunc8()  {}
func (s *Suite016) TestFunc9()  {}
func (s *Suite016) TestFunc10() {}

func TestSuite016(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite016))
}
