package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func NoError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func MustReadFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	NoError(err)
	return data
}

func MustUnmarshalJson(data []byte, o interface{}) {
	NoError(json.Unmarshal(data, o))
}

func MustMarshalJson(o interface{}) []byte {
	data, err := json.Marshal(o)
	NoError(err)
	return data
}

func MustParseRedisURL(url string) *redis.Options {
	opt, err := redis.ParseURL(url)
	NoError(err)
	return opt
}

func ServeAddrFromConfig(c *Config) string {
	addr, port := c.ServeOn.Addr, c.ServeOn.Port

	if len(addr) == 0 && len(port) == 0 {
		return ""
	}
	if len(addr) == 0 {
		return ":" + port
	}

	if len(port) == 0 {
		return addr + ":0"
	}

	return addr + ":" + port
}

func MustMarshalProto(obj proto.Message) []byte {
	data, err := proto.Marshal(obj)
	NoError(err)
	return data
}

func MustUnmarshalProto(data []byte, o proto.Message) {
	NoError(proto.Unmarshal(data, o))
}

func MakeCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
