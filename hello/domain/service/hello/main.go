package hello

import (
	helloEntity "example/hello/domain/entity/hello"
	helloRepository "example/hello/domain/repository/hello"
)

func GetHello() helloEntity.Hello {
	return helloRepository.GetHello()
}
