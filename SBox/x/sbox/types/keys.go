package types

const (
	// ModuleName defines the module name
	ModuleName = "sbox"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_sbox"
)

var (
	ParamsKey = []byte("p_sbox")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
