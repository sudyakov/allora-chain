package rewards

import (
	"math/rand"
	"sort"

	alloraMath "github.com/allora-network/allora-chain/math"
)

// A structure to hold the original value and a random tiebreaker
type SortableTopicId struct {
	Value      TopicId
	Weight     *alloraMath.Dec
	Tiebreaker uint32
}

// Sorts the given slice of topics in descending order according to their corresponding return, using pseudorandom tiebreaker
// e.g. ([]uint64{1, 2, 3}, map[uint64]uint64{1: 2, 2: 2, 3: 3}, 0) -> [3, 1, 2] or [3, 2, 1]
func SortTopicsByWeightDescWithRandomTiebreaker(topicIds []TopicId, weights map[TopicId]*alloraMath.Dec, randSeed BlockHeight) []TopicId {
	// Convert the slice of Ts to a slice of SortableItems, each with a random tiebreaker
	r := rand.New(rand.NewSource(randSeed))
	items := make([]SortableTopicId, len(topicIds))
	for i, topicId := range topicIds {
		items[i] = SortableTopicId{topicId, weights[topicId], r.Uint32()}
	}

	// Sort the slice of SortableItems
	// If the values are equal, the tiebreaker will decide their order
	sort.Slice(items, func(i, j int) bool {
		if items[i].Value == items[j].Value {
			return items[i].Tiebreaker > items[j].Tiebreaker
		}
		return (*items[i].Weight).Gt(*items[j].Weight)
	})

	// Extract and print the sorted values to demonstrate the sorting
	sortedValues := make([]TopicId, len(topicIds))
	for i, item := range items {
		sortedValues[i] = item.Value
	}
	return sortedValues
}

// Returns a map of topicId to weights of the top N topics by weight in descending order
// It is assumed that topicIds is of a reasonable size, throttled by perhaps MaxTopicsPerBlock global param
func SkimTopTopicsByWeightDesc(weights map[TopicId]*alloraMath.Dec, N uint64, block BlockHeight) map[TopicId]*alloraMath.Dec {
	// Build a slice of topicIds
	topicIds := make([]TopicId, 0, len(weights))
	for topicId := range weights {
		topicIds = append(topicIds, topicId)
	}

	// Sort the topics by weight in descending order with a random tiebreaker
	sortedTopicIds := SortTopicsByWeightDescWithRandomTiebreaker(topicIds, weights, block)

	// Extract the top N values
	weightsOfTopN := make(map[TopicId]*alloraMath.Dec, N)
	numberToAdd := N
	if (uint64)(len(sortedTopicIds)) < N {
		numberToAdd = (uint64)(len(sortedTopicIds))
	}

	for i := uint64(0); i < numberToAdd; i++ {
		weightsOfTopN[sortedTopicIds[i]] = weights[sortedTopicIds[i]]
	}

	return weightsOfTopN
}