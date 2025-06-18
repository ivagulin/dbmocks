package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite002 struct {
	SuiteBase
}

func (s *Suite002) TestFunc1()  {}
func (s *Suite002) TestFunc2()  {}
func (s *Suite002) TestFunc3()  {}
func (s *Suite002) TestFunc4()  {}
func (s *Suite002) TestFunc5()  {}
func (s *Suite002) TestFunc6()  {}
func (s *Suite002) TestFunc7()  {}
func (s *Suite002) TestFunc8()  {}
func (s *Suite002) TestFunc9()  {}
func (s *Suite002) TestFunc10() {}

func TestSuite002(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite002))
}
