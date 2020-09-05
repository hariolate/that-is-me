package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
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

func MustUnmarshal(data []byte, o interface{}) {
	NoError(json.Unmarshal(data, o))
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

func MustMarshalProto(obj interface{}) []byte {
	data, err := proto.Marshal(obj.(proto.Message))
	NoError(err)
	return data
}

func MustUnmarshalProto(data []byte, o interface{}) {
	NoError(proto.Unmarshal(data, o.(proto.Message)))
}

func MustMarshalAnyProto(msg proto.Message) *any.Any {
	res, err := ptypes.MarshalAny(msg)
	NoError(err)
	return res
}

func MustUnmarshalAnyProto(data *any.Any, msg proto.Message) {
	err := ptypes.UnmarshalAny(data, msg)
	NoError(err)
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
