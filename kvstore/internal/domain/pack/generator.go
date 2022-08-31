package pack

type Generator interface {
	GenerateKey() *Key
}
