package outbound

type IDGenerator interface {
	NewID() string
}
