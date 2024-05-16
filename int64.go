package nullable

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

// Int64 is a nullable int64.
type Int64 struct {
	sql.NullInt64
}

// NewInt64 returns a new Int64.
func NewInt64(i int64, valid bool) Int64 {
	return Int64{
		sql.NullInt64{
			Int64: i,
			Valid: valid,
		},
	}
}

// NewInt64FromPtr returns a new Int64 from a pointer.
func NewInt64FromPtr(i *int64) Int64 {
	if i == nil {
		return NewInt64(0, false)
	}

	return NewInt64(*i, true)
}

// MarshalJSON implements the json.Marshaler interface.
func (n Int64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Int64)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *Int64) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		n.Int64, n.Valid = 0, false

		return nil
	}

	if err := json.Unmarshal(b, &n.Int64); err != nil {
		return err
	}

	n.Valid = true

	return nil
}
