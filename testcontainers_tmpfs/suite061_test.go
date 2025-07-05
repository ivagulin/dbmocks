package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite061 struct {
	SuiteBase
}

func (s *Suite061) TestFunc1()  {}
func (s *Suite061) TestFunc2()  {}
func (s *Suite061) TestFunc3()  {}
func (s *Suite061) TestFunc4()  {}
func (s *Suite061) TestFunc5()  {}
func (s *Suite061) TestFunc6()  {}
func (s *Suite061) TestFunc7()  {}
func (s *Suite061) TestFunc8()  {}
func (s *Suite061) TestFunc9()  {}
func (s *Suite061) TestFunc10() {}

func TestSuite061(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite061))
}
