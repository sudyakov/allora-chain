package msgserver_test

import (
	"math"

	"github.com/allora-network/allora-chain/x/emissions/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *KeeperTestSuite) TestInsertInferencesAndQuery() {
	ctx, msgServer := s.ctx, s.msgServer
	require := s.Require()
	s.CreateOneTopic()

	// Mock setup for inferences
	inferences := []*types.Inference{
		{TopicId: 0, Worker: sdk.AccAddress(PKS[0].Address()).String(), Value: 2200},
		{TopicId: 0, Worker: sdk.AccAddress(PKS[1].Address()).String(), Value: 2100},
		{TopicId: 2, Worker: sdk.AccAddress(PKS[2].Address()).String(), Value: 12},
	}

	// Call the InsertInferences function to test writes
	processInferencesMsg := &types.MsgInsertInferences{
		Inferences: inferences,
	}
	_, err := msgServer.InsertInferences(ctx, processInferencesMsg)
	require.NoError(err, "Processing Inferences should not fail")

	/*
	 * Inferences over threshold should be returned
	 */
	// Ensure low ts for topic 1
	var topicId = uint64(0)
	inferenceBlock := int64(0x16)

	// _, err = msgServer.SetLatestInferencesTimestamp(ctx, inferencesMsg)
	err = s.emissionsKeeper.UpdateTopicEpochLastEnded(ctx, topicId, inferenceBlock)
	require.NoError(err, "Setting latest inference timestamp should not fail")

	allInferences, err := s.emissionsKeeper.GetLatestInferencesFromTopic(ctx, uint64(0))
	require.Equal(len(allInferences), 1)
	for _, inference := range allInferences {
		require.Equal(len(inference.Inferences.Inferences), 2)
	}
	require.NoError(err, "Inferences over ts threshold should be returned")

	/*
	 * Inferences under threshold should not be returned
	 */
	inferenceBlock = math.MaxInt64

	err = s.emissionsKeeper.UpdateTopicEpochLastEnded(ctx, topicId, inferenceBlock)
	require.NoError(err)

	allInferences, err = s.emissionsKeeper.GetLatestInferencesFromTopic(ctx, uint64(1))

	require.Equal(len(allInferences), 0)
	require.NoError(err, "Inferences under ts threshold should not be returned")
}