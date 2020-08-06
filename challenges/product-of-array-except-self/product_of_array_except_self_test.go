package product_of_array_except_self_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aikin/go-algorithms/challenges/product-of-array-except-self"
)

var _ = Describe("ProductOfArrayExceptSelf", func() {
	Context("Product of array except self general logic", func() {
		When("nums [1, 2, 3, 4]", func() {
			It("should return result: [24, 12, 8, 6], err: nil", func() {

				nums := []int{1, 2, 3, 4}

				result , err := ProductExceptSelf(nums)

				Ω(err).Should(BeNil())
				Ω(result).Should(Equal([]int {24, 12, 8, 6}))
			})
		})

		When("nums [2, 4, 8, 5, 6]", func() {
			It("should return result: [960, 480, 240, 384, 320], err: nil", func() {

				nums := []int{2, 4, 8, 5, 6}

				result , err := ProductExceptSelf(nums)

				Ω(err).Should(BeNil())
				Ω(result).Should(Equal([]int {960, 480, 240, 384, 320}))
			})
		})
	})
})