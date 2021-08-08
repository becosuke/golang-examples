package hello

import (
	"example/hello-gin/entity/hello"
)

func GetHello() hello.Hello {
	return hello.Hello{"hello"}
}
