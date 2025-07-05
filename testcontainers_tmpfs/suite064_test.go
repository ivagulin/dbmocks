package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite064 struct {
	SuiteBase
}

func (s *Suite064) TestFunc1()  {}
func (s *Suite064) TestFunc2()  {}
func (s *Suite064) TestFunc3()  {}
func (s *Suite064) TestFunc4()  {}
func (s *Suite064) TestFunc5()  {}
func (s *Suite064) TestFunc6()  {}
func (s *Suite064) TestFunc7()  {}
func (s *Suite064) TestFunc8()  {}
func (s *Suite064) TestFunc9()  {}
func (s *Suite064) TestFunc10() {}

func TestSuite064(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite064))
}
