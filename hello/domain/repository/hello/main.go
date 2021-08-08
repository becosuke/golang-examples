package hello

import (
	"example/hello/domain/entity/hello"
)

func GetHello() hello.Hello {
	return hello.Hello{Message: "hello"}
}
