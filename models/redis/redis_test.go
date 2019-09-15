package redis

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "The Lazy Traveler Suite")
}

var _ = Describe("models", func() {
	Context("##redis", func() {
		client := NewClient()
		It("should save value in redis cache", func() {
			_ = client.Put("key-123", "value", time.Minute)
			val, _ := client.Get("key-123")
			Expect(val).To(Equal("value"))
		})

		It("should return value from redis cache", func() {
			_ = client.Put("key-123", "value", time.Minute)
			val, _ := client.Get("key-123")
			Expect(val).To(Equal("value"))
		})

		It("should delete value in redis cache", func() {
			_ = client.Put("key-123", "value", time.Minute)
			val, _ := client.Get("key-123")
			Expect(val).To(Equal("value"))
			_ = client.Delete("key-123")
			val, _ = client.Get("key-123")
			Expect(val).To(Equal(""))
		})
	})
})
