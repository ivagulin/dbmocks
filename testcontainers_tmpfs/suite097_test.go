package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite097 struct {
	SuiteBase
}

func (s *Suite097) TestFunc1()  {}
func (s *Suite097) TestFunc2()  {}
func (s *Suite097) TestFunc3()  {}
func (s *Suite097) TestFunc4()  {}
func (s *Suite097) TestFunc5()  {}
func (s *Suite097) TestFunc6()  {}
func (s *Suite097) TestFunc7()  {}
func (s *Suite097) TestFunc8()  {}
func (s *Suite097) TestFunc9()  {}
func (s *Suite097) TestFunc10() {}

func TestSuite097(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite097))
}
