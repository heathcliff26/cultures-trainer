package app

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt32Entry(t *testing.T) {
	t.Run("NewEntry", func(t *testing.T) {
		assert := assert.New(t)
		require := require.New(t)

		entry := NewInt32Entry("Test Label")
		require.NotNil(entry, "Should return a new Int32Entry")
		require.NotNil(entry.Entry, "Entry widget should not be nil")
		require.NotNil(entry.Checkbox, "Checkbox widget should not be nil")
		assert.NotNil(entry.canvas, "Should have a canvas object")

		assert.Equal(fmt.Sprintf("%d", entry.value), entry.Entry.Text, "Internal value should match text of entry")
		assert.Equal("Test Label", strings.TrimSpace(entry.Checkbox.Text), "Checkbox should have correct label")
	})
	t.Run("Set", func(t *testing.T) {
		assert := assert.New(t)

		entry := NewInt32Entry("Test Label")
		testValue := int32(42)
		entry.Set(testValue)

		assert.Equal(testValue, entry.value, "Internal value should be set correctly")
		assert.Equal(fmt.Sprintf("%d", testValue), entry.Entry.Text, "Entry text should reflect the set value")
	})
	t.Run("Get", func(t *testing.T) {
		assert := assert.New(t)

		entry := NewInt32Entry("Test Label")
		testValue := int32(100)
		entry.Set(testValue)

		assert.Equal(testValue, entry.Get(), "Retrieved value should match the set value")
	})
	t.Run("EntryValidator", func(t *testing.T) {
		assert := assert.New(t)

		entry := NewInt32Entry("Test Label")

		assert.NoError(entry.Entry.Validator("12345"), "Should accept valid int32 string")
		assert.Error(entry.Entry.Validator("abc"), "Should reject non-integer string")
		assert.Error(entry.Entry.Validator("2147483648"), "Should reject out-of-range int32 string")
	})
	t.Run("EntryOnChanged", func(t *testing.T) {
		assert := assert.New(t)

		entry := NewInt32Entry("Test Label")

		entry.Entry.SetText("invalid")
		assert.Equal(int32(0), entry.value, "Internal value should not change on invalid entry change")

		entry.Entry.SetText("2147483648")
		assert.Equal(int32(0), entry.value, "Internal value should not change on out-of-range entry change")

		entry.Entry.SetText("256")
		assert.Equal(int32(256), entry.value, "Internal value should update on valid entry change")

		entry.Entry.SetText("")
		assert.Equal(int32(0), entry.value, "Internal value should be zero when entry is cleared")
	})
}
