package redisdb

import (
	"encoding/json"
	"testing"

	. "github.com/franela/goblin"
)

const url = "localhost:6379"
const password = ""

func TestRedisDB(t *testing.T) {
	g := Goblin(t)
	g.Describe("Test Redis", func() {
		redis, err := Initial(url, password)
		g.It("Error Should be nil", func() {
			g.Assert(err).Equal(nil)
		})
		g.It("Should Set Key (string)", func() {
			err := redis.Set("test", "Hello")
			g.Assert(err).Equal(nil)
		})
		g.It("Should Get Value (string)", func() {
			result, err := redis.Get("test")
			g.Assert(err).Equal(nil)
			g.Assert(result).Equal("Hello")
		})
		g.It("Should Set Key (struct)", func() {
			type Test struct {
				Message string
			}
			test := Test{
				Message: "Hello",
			}
			bt, _ := json.Marshal(test)
			err := redis.Set("testStruct", string(bt))
			g.Assert(err).Equal(nil)
		})
		g.It("Should Get Value (struct)", func() {
			type Test struct {
				Message string
			}
			testResult := Test{}
			result, err := redis.Get("testStruct")
			g.Assert(err).Equal(nil)
			err = json.Unmarshal([]byte(result), &testResult)
			g.Assert(err).Equal(nil)
		})
	})
}
