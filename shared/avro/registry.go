package avro

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/hamba/avro"
	"github.com/riferrei/srclient"
)

func Encode(sr *srclient.SchemaRegistryClient, subject, schemaStr string, data interface{}) ([]byte, int, error) {
	schema, err := sr.CreateSchema(subject, schemaStr, srclient.Avro)
	if err != nil {
		schema, err = sr.GetLatestSchema(subject)
		if err != nil {
			return nil, 0, err
		}
	}

	id := schema.ID()

	parsed, err := avro.Parse(schemaStr)
	if err != nil {
		return nil, 0, err
	}

	bin, err := avro.Marshal(parsed, data)
	if err != nil {
		return nil, 0, err
	}

	var buf bytes.Buffer
	buf.WriteByte(0)
	binary.Write(&buf, binary.BigEndian, int32(id))
	buf.Write(bin)

	return buf.Bytes(), id, nil
}

func Decode(sr *srclient.SchemaRegistryClient, payload []byte) (map[string]interface{}, error) {
	if len(payload) < 5 {
		return nil, fmt.Errorf("invalid avro payload")
	}

	id := int(binary.BigEndian.Uint32(payload[1:5]))
	schema, err := sr.GetSchema(id)
	if err != nil {
		return nil, err
	}

	parsed, _ := avro.Parse(schema.Schema())

	var out map[string]interface{}
	err = avro.Unmarshal(parsed, payload[5:], &out)
	return out, err
}
