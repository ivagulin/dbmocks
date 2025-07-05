package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite036 struct {
	SuiteBase
}

func (s *Suite036) TestFunc1()  {}
func (s *Suite036) TestFunc2()  {}
func (s *Suite036) TestFunc3()  {}
func (s *Suite036) TestFunc4()  {}
func (s *Suite036) TestFunc5()  {}
func (s *Suite036) TestFunc6()  {}
func (s *Suite036) TestFunc7()  {}
func (s *Suite036) TestFunc8()  {}
func (s *Suite036) TestFunc9()  {}
func (s *Suite036) TestFunc10() {}

func TestSuite036(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite036))
}
