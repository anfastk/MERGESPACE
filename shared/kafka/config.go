package kafka

import (
	"github.com/twmb/franz-go/pkg/kgo"
	"time"
)

func DefaultKafkaClientOptions(brokers []string) []kgo.Opt {
	return []kgo.Opt{
		kgo.SeedBrokers(brokers...),
		kgo.RequiredAcks(kgo.AllISRAcks()),
		kgo.ProducerLinger(10 * time.Millisecond),
		kgo.BlockRebalanceOnPoll(),
		kgo.FetchMaxWait(500 * time.Millisecond),
	}
}
