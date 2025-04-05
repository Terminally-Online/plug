package streams

import "solver/internal/solver/signature"

type Filter struct {
	Categories []string
	ChainIDs   []int64
}

func ApplyFilter(_ signature.Plugs) bool {
	// Check if we are interest in any of the protocols/actions in this plug
	// for _, plug := range plugs.Plugs {
	//		TODO: Filter down to the protocol and action names.
	// }

	return true
}
