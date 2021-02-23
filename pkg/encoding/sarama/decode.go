package sarama

import (
	"github.com/golang/protobuf/proto"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/broker/sarama"
	encoding "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/encoding/proto"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/errors"
)

// Unmarshal a sarama message into a protobuffer
func Unmarshal(msg *sarama.Msg, pb proto.Message) error {
	// Unmarshal Sarama message to Envelope
	err := encoding.Unmarshal(msg.Value(), pb)
	if err != nil {
		return errors.FromError(err).SetComponent(component)
	}
	return nil
}