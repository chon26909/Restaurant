package utils_test

import (
	"restaurant/utils"
	"testing"
)

func TestGetPublicUrl(t *testing.T) {
	t.Run("test", func(t *testing.T) {

		type testCase struct {
			imageName string
			expected  string
		}

		testCases := []testCase{
			{
				imageName: "test-image",
				expected:  "test-image",
			},
			{
				imageName: "image1",
				expected:  "image1",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.imageName, func(t *testing.T) {
				url := utils.GetPublicUrl(tc.imageName)

				if url != tc.expected {
					t.Errorf("expected %s, got %s", tc.expected, url)
				}
			})
		}
	})
}
