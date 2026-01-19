package trainer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConst(t *testing.T) {
	t.Run("LengthsMatch", func(t *testing.T) {
		assert := assert.New(t)

		assert.Equal(len(StorageLocations), ResourceCount)
		assert.Equal(len(StorageIndexes), ResourceCount)
	})
	t.Run("IndexMapping", func(t *testing.T) {
		assert := assert.New(t)

		for i, name := range StorageLocations {
			index, exists := StorageIndexes[name]
			assert.True(exists, "StorageIndexes should contain key for %s", name)
			assert.Equal(i, index, "StorageIndexes[%s] should be %d", name, i)
		}
	})
	t.Run("CategoriesAreComplete", func(t *testing.T) {
		assert := assert.New(t)

		allItems := make(map[int]int, ResourceCount)

		checkCategory := func(items []string) {
			for _, name := range items {
				index, exists := StorageIndexes[name]
				assert.True(exists, "StorageIndexes should contain key for %s", name)
				allItems[index]++
			}
		}

		checkCategory(CategoryNahrung)
		checkCategory(CategoryBauwaren)
		checkCategory(CategoryResourcen)
		checkCategory(CategoryWaffen)
		checkCategory(CategoryBonus)
		checkCategory(CategorySonstiges)

		for i, count := range allItems {
			assert.Equal(1, count, "Storage item %s appears %d times in categories", StorageLocations[i], count)
		}

		assert.Equal(ResourceCount, len(allItems), "Not all storage items are covered in categories")
	})
}
