package pgtestpkg

import (
	"context"
	"github.com/ivagulin/dbmocks/repo"
	"github.com/stretchr/testify/suite"
)

type SuiteBase struct {
	suite.Suite
	Repo   *repo.Repo
	stopDB func()
}

func (s *SuiteBase) SetupSuite() {
	s.Repo, s.stopDB = repo.NewTestRepo()
}

func (s *SuiteBase) TearDownSuite() {
	s.stopDB()
}

func (s *SuiteBase) TearDownTest() {
	s.Require().NoError(s.Repo.CleanTables(context.Background()))
}
