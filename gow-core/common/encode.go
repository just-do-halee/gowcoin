package common

type ByteEncoder interface {
	ToBytes() []byte
}

type HashEncoder interface {
	ToHash() Hash
}
