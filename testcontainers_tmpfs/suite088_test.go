package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite088 struct {
	SuiteBase
}

func (s *Suite088) TestFunc1()  {}
func (s *Suite088) TestFunc2()  {}
func (s *Suite088) TestFunc3()  {}
func (s *Suite088) TestFunc4()  {}
func (s *Suite088) TestFunc5()  {}
func (s *Suite088) TestFunc6()  {}
func (s *Suite088) TestFunc7()  {}
func (s *Suite088) TestFunc8()  {}
func (s *Suite088) TestFunc9()  {}
func (s *Suite088) TestFunc10() {}

func TestSuite088(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite088))
}
