package tests

import (
	"temporal/demo/versioning/workflow"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/api/workflowservicemock/v1"
	"go.temporal.io/sdk/worker"
)

type replayTestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
	service  *workflowservicemock.MockWorkflowServiceClient
}

func TestReplayTestSuite(t *testing.T) {
	s := new(replayTestSuite)
	suite.Run(t, s)
}

func (s *replayTestSuite) SetupTest() {
	s.mockCtrl = gomock.NewController(s.T())
	s.service = workflowservicemock.NewMockWorkflowServiceClient(s.mockCtrl)
}

func (s *replayTestSuite) TearDownTest() {
	s.mockCtrl.Finish() // assert mock’s expectations
}

func (s *replayTestSuite) TestReplayFromInitialVersion() {
	replayer := worker.NewWorkflowReplayer()

	replayer.RegisterWorkflow(workflow.CustomerWorkflow)

	err := replayer.ReplayWorkflowHistoryFromJSONFile(nil, "initversionhistory.json")
	require.NoError(s.T(), err)
}

func (s *replayTestSuite) TestReplayFromVersion1() {
	replayer := worker.NewWorkflowReplayer()

	replayer.RegisterWorkflow(workflow.CustomerWorkflow)

	err := replayer.ReplayWorkflowHistoryFromJSONFile(nil, "v1history.json")
	require.NoError(s.T(), err)
}