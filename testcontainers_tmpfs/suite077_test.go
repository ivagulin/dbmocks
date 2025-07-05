package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite077 struct {
	SuiteBase
}

func (s *Suite077) TestFunc1()  {}
func (s *Suite077) TestFunc2()  {}
func (s *Suite077) TestFunc3()  {}
func (s *Suite077) TestFunc4()  {}
func (s *Suite077) TestFunc5()  {}
func (s *Suite077) TestFunc6()  {}
func (s *Suite077) TestFunc7()  {}
func (s *Suite077) TestFunc8()  {}
func (s *Suite077) TestFunc9()  {}
func (s *Suite077) TestFunc10() {}

func TestSuite077(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite077))
}
