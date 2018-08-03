package messaging

import (
	"time"

	"github.com/google/uuid"
)

var (
	// GenerateUUID creates a new random UUID or panics
	GenerateUUID = func() UUID {
		return UUID(uuid.New())
	}
)

type (
	// UUID is a 128 bit (16 byte) Universal Unique IDentifier as defined in RFC4122
	UUID [16]byte

	// Metadata is a container of metadata information
	Metadata map[string]interface{}

	// Message is a interface describing a message.
	// A message can be a command, domain event or some other type of message.
	Message interface {
		// UUID returns the identifier of this message
		UUID() UUID

		// CreatedAt returns the created time of the message
		CreatedAt() time.Time

		// Payload returns the payload of the message
		Payload() interface{}

		// Metadata return the message metadata
		Metadata() Metadata

		// WithAddedMetadata Returns new instance of the message with key and value added to metadata
		WithAddedMetadata(key string, value interface{}) Message
	}
)