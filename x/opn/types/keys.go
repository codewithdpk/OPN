package types

const (
	// ModuleName defines the module name
	ModuleName = "opn"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_opn"
)

var (
	ParamsKey = []byte("p_opn")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
