package extpoint

type IDEncode interface {
	EncodeId(ns string, id uint64) (string, error)
}
