package producer

import (
	"bytes"
	"context"
	"encoding/binary"
	"time"

	"github.com/anfastk/MERGESPACE/shared/kafka"
	"github.com/hamba/avro"
	"github.com/riferrei/srclient"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client   *kgo.Client
	sr       *srclient.SchemaRegistryClient
	schema   string
	subject  string
	topic    string
	schemaID int
	codec    avro.Schema
}

func New(brokers []string, srURL, topic, schemaStr string) (*Producer, error) {
	client, err := kgo.NewClient(kafka.DefaultKafkaClientOptions(brokers)...)
	if err != nil {
		return nil, err
	}
	sr := srclient.CreateSchemaRegistryClient(srURL)
	subject := topic + "-value"

	schema, err := sr.CreateSchema(subject, schemaStr, srclient.Avro)
	if err != nil {
		return nil, err
	}

	parsed, err := avro.Parse(schemaStr)
	if err != nil {
		return nil, err
	}

	return &Producer{
		client:   client,
		sr:       sr,
		topic:    topic,
		subject:  subject,
		schemaID: schema.ID(),
		codec:    parsed,
	}, nil
}

/* func (p *Producer) Publish(ctx context.Context, key []byte, event interface{}) error {
	payload, _, err := avro.Encode(p.sr, p.subject, p.schema, event)
	if err != nil {
		return err
	}

	rec := &kgo.Record{
		Topic: p.topic,
		Key:   key,
		Value: payload,
	}

	res := p.client.ProduceSync(ctx, rec)
	return res.FirstErr()
}
*/

func (p *Producer) Publish(ctx context.Context, key []byte, event interface{}) error {
	bin, err := avro.Marshal(p.codec, event)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	buf.WriteByte(0)
	binary.Write(&buf, binary.BigEndian, int32(p.schemaID))
	buf.Write(bin)

	rec := &kgo.Record{
		Topic: p.topic,
		Key:   key,
		Value: buf.Bytes(),
	}

	produceCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return p.client.ProduceSync(produceCtx, rec).FirstErr()
}
