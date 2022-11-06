package data

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumber_BytesSize(t *testing.T) {
	num := Number(15.5)
	require.Equal(t, Int64Size, num.BytesSize())
}

func TestNumber_Serialize(t *testing.T) {
	num := Number(15.5)
	expected := []byte{
		0b01000000,
		0b00101111,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00000000,
	}

	result, err := num.Serialize()
	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestNumber_Deserialize(t *testing.T) {
	raw := []byte{
		0b01000000,
		0b00101111,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00000000,
		0b00000000,
	}
	expected := Number(15.5)

	var result Number
	err := result.Deserialize(raw)

	require.Nil(t, err)
	require.Equal(t, expected, result)
}
