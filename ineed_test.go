package i_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xchapter7x/ineed"
)

var _ = Describe("given ineed package", func() {
	Describe("given Want() func", func() {

		Describe("given want.ToMap() method", func() {
			Context("when given a name -> interface{} combination", func() {
				It("then it should register the mapping to the dependencies list", func() {
					deps := i.Want()
					Ω(len(deps)).Should(Equal(0))
					deps.ToMap("something", "to somthing else")
					Ω(len(deps)).Should(Equal(1))
				})
			})
		})

		Describe("given want.ToUse() method", func() {
			Context("when given an interface{} object", func() {
				It("then it should add it to the registered dependencies list", func() {
					deps := i.Want()
					Ω(len(deps)).Should(Equal(0))
					deps.ToUse("to somthing else")
					Ω(len(deps)).Should(Equal(1))
				})
			})
		})

		Describe("given want.Get() method", func() {
			Context("when called with the name of a registered dependency", func() {
				controlValue := "to somthing else"
				controlName := "something"
				deps := i.Want()
				deps.ToMap(controlName, controlValue)

				It("then it should return the interface{} of the dependency", func() {
					Ω(deps.Get(controlName)).Should(Equal(controlValue))
				})
			})
		})

		Describe("given want.CastInto() method", func() {
			Context("when called on a pointer to a struct", func() {
				var testObject = new(struct {
					RandomName string
				})
				controlValue := "to somthing else"
				deps := i.Want()
				deps.ToUse(controlValue)

				It("then it should set all public fields in the struct to a registered dependency of a matching type", func() {
					deps.CastInto(testObject)
					Ω(testObject.RandomName).Should(Equal(controlValue))
				})
			})
		})

		Describe("given want.MapInto() method", func() {
			Context("when called on a pointer to a struct", func() {
				var testObject = new(struct {
					Something string
				})
				controlValue := "to somthing else"
				controlName := "Something"
				deps := i.Want()
				deps.ToMap(controlName, controlValue)

				It("then is should set all public fields whos name has a corresbonding record in the dependency repo", func() {
					deps.MapInto(testObject)
					Ω(testObject.Something).Should(Equal(controlValue))
				})
			})
		})
	})
})
