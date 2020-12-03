package extpoint

type IdGen interface {
    NextUniqueId(ns string) (uint64, error)
}
