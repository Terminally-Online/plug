package solver

import (
	"math/big"
	"os"
	"solver/solver/signature"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var (
	mockPlugs = signature.Plugs{
		Socket: common.HexToAddress("0x62180042606624f02d8a130da8a3171e9b33894d"),
		Plugs: []signature.Plug{{
			To:    common.HexToAddress("0x62180042606624f02d8a130da8a3171e9b33894d"),
			Data:  []byte{},
			Value: big.NewInt(0),
		}},
		Solver: []byte{},
		Salt:   []byte{},
	}
	mockLivePlugs = signature.LivePlugs{
		Plugs:      mockPlugs,
		Signature: []byte{},
	}
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func setupTest(_ *testing.T) (*Solver, func()) {
	solver := New()

	return solver, func() {}
}

func TestNewSolver(t *testing.T) {
	solver, cleanup := setupTest(t)
	defer cleanup()

	t.Run("should have solver", func(t *testing.T) {
		if solver == nil {
			t.Error("expected solver to be initialized")
		}
	})

	t.Run("should have intended salt", func(t *testing.T) {})
	t.Run("should have intended signature", func(t *testing.T) {})
}
