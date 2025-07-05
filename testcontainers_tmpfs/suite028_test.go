package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite028 struct {
	SuiteBase
}

func (s *Suite028) TestFunc1()  {}
func (s *Suite028) TestFunc2()  {}
func (s *Suite028) TestFunc3()  {}
func (s *Suite028) TestFunc4()  {}
func (s *Suite028) TestFunc5()  {}
func (s *Suite028) TestFunc6()  {}
func (s *Suite028) TestFunc7()  {}
func (s *Suite028) TestFunc8()  {}
func (s *Suite028) TestFunc9()  {}
func (s *Suite028) TestFunc10() {}

func TestSuite028(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite028))
}
