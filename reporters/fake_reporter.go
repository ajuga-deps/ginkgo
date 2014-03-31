package reporters

import (
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

//FakeReporter is useful for testing purposes
type FakeReporter struct {
	Config config.GinkgoConfigType

	BeginSummary         *types.SuiteSummary
	SpecWillRunSummaries []*types.SpecSummary
	SpecSummaries        []*types.SpecSummary
	EndSummary           *types.SuiteSummary
}

func NewFakeReporter() *FakeReporter {
	return &FakeReporter{
		SpecWillRunSummaries: make([]*types.SpecSummary, 0),
		SpecSummaries:        make([]*types.SpecSummary, 0),
	}
}

func (fakeR *FakeReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
	fakeR.Config = config
	fakeR.BeginSummary = summary
}

func (fakeR *FakeReporter) SpecWillRun(specSummary *types.SpecSummary) {
	fakeR.SpecWillRunSummaries = append(fakeR.SpecWillRunSummaries, specSummary)
}

func (fakeR *FakeReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	fakeR.SpecSummaries = append(fakeR.SpecSummaries, specSummary)
}

func (fakeR *FakeReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	fakeR.EndSummary = summary
}