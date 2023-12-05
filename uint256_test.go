package nullable_test

import (
	"database/sql/driver"
	"encoding/json"
	"math/big"
	"testing"

	ethmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/m0t0k1ch1-go/bigutil/v2"

	"github.com/m0t0k1ch1-go/nullable"
	"github.com/m0t0k1ch1-go/nullable/internal/testutil"
)

func TestUint256NullableString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   nullable.Uint256
			out  nullable.String
		}{
			{
				"null",
				nullable.NewUint256(bigutil.Uint256{}, false),
				nullable.NewString("", false),
			},
			{
				"min",
				nullable.NewUint256(bigutil.MustBigIntToUint256(big.NewInt(0)), true),
				nullable.NewString("0x0", true),
			},
			{
				"max",
				nullable.NewUint256(bigutil.MustBigIntToUint256(ethmath.MaxBig256), true),
				nullable.NewString("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", true),
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				testutil.Equal(t, tc.out, tc.in.NullableString())
			})
		}
	})
}

func TestUint256Value(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   nullable.Uint256
			out  driver.Value
		}{
			{
				"null",
				nullable.NewUint256(bigutil.Uint256{}, false),
				nil,
			},
			{
				"min",
				nullable.NewUint256(bigutil.MustBigIntToUint256(big.NewInt(0)), true),
				[]byte{0x0},
			},
			{
				"max",
				nullable.NewUint256(bigutil.MustBigIntToUint256(ethmath.MaxBig256), true),
				[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				v, err := tc.in.Value()
				if err != nil {
					t.Fatal(err)
				}

				testutil.Equal(t, tc.out, v)
			})
		}
	})
}

func TestUint256Scan(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   any
			out  nullable.Uint256
		}{
			{
				"null",
				nil,
				nullable.NewUint256(bigutil.Uint256{}, false),
			},
			{
				"min",
				[]byte{0x0},
				nullable.NewUint256(bigutil.MustBigIntToUint256(big.NewInt(0)), true),
			},
			{
				"max",
				[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
				nullable.NewUint256(bigutil.MustBigIntToUint256(ethmath.MaxBig256), true),
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var n nullable.Uint256
				if err := n.Scan(tc.in); err != nil {
					t.Fatal(err)
				}

				testutil.Equal(t, tc.out.Valid, n.Valid)
				testutil.Equal(t, n.Uint256.BigInt().Cmp(tc.out.Uint256.BigInt()), 0)
			})
		}
	})
}

func TestUint256MarshalJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   nullable.Uint256
			out  []byte
		}{
			{
				"null",
				nullable.NewUint256(bigutil.Uint256{}, false),
				[]byte("null"),
			},
			{
				"min",
				nullable.NewUint256(bigutil.MustBigIntToUint256(big.NewInt(0)), true),
				[]byte(`"0x0"`),
			},
			{
				"max",
				nullable.NewUint256(bigutil.MustBigIntToUint256(ethmath.MaxBig256), true),
				[]byte(`"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"`),
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				b, err := json.Marshal(tc.in)
				if err != nil {
					t.Fatal(err)
				}

				testutil.Equal(t, tc.out, b)
			})
		}
	})
}

func TestUint256UnmarshalJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tcs := []struct {
			name string
			in   []byte
			out  nullable.Uint256
		}{
			{
				"null",
				[]byte("null"),
				nullable.NewUint256(bigutil.Uint256{}, false),
			},
			{
				"min",
				[]byte(`"0x0"`),
				nullable.NewUint256(bigutil.MustBigIntToUint256(big.NewInt(0)), true),
			},
			{
				"max",
				[]byte(`"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"`),
				nullable.NewUint256(bigutil.MustBigIntToUint256(ethmath.MaxBig256), true),
			},
		}

		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				var n nullable.Uint256
				if err := json.Unmarshal(tc.in, &n); err != nil {
					t.Fatal(err)
				}

				testutil.Equal(t, tc.out.Valid, n.Valid)
				testutil.Equal(t, n.Uint256.BigInt().Cmp(tc.out.Uint256.BigInt()), 0)
			})
		}
	})
}
