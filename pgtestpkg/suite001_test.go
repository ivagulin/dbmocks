package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite001 struct {
	SuiteBase
}

func (s *Suite001) TestFunc1()  {}
func (s *Suite001) TestFunc2()  {}
func (s *Suite001) TestFunc3()  {}
func (s *Suite001) TestFunc4()  {}
func (s *Suite001) TestFunc5()  {}
func (s *Suite001) TestFunc6()  {}
func (s *Suite001) TestFunc7()  {}
func (s *Suite001) TestFunc8()  {}
func (s *Suite001) TestFunc9()  {}
func (s *Suite001) TestFunc10() {}

func TestSuite001(t *testing.T) {
	suite.Run(t, new(Suite001))
}
