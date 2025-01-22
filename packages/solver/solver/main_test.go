package solver

import (
	"math/big"
	"os"
	"solver/types"
	"testing"
)

var (
	mockPlug = types.Plug{
		Socket: "0x62180042606624f02d8a130da8a3171e9b33894d",
		Plugs: []types.Transaction{{
			To:    "0x",
			Data:  "0x",
			Value: *big.NewInt(0),
		}},
		Solver: "0x62180042606624f02d8a130da8a3171e9b33894d",
		Salt:   "1234567890",
	}
	mockPlugs = types.Plugs{
		Plug:      mockPlug,
		Signature: "0x",
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
