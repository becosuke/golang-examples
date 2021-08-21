package collatz

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetFlags(log.Ldate | log.Ltime)
	if err := setup(); err != nil {
		log.Fatal(err)
	}
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func setup() error {
	return nil
}

func teardown() {
}

func TestCollatz(t *testing.T) {
	log.Print(Collatz(3 * 3))
	log.Print(Collatz(3 * 5))
	log.Print(Collatz(5 * 5))
	log.Print(Collatz(3 * 7))
	log.Print(Collatz(5 * 7))
	log.Print(Collatz(7 * 7))
}
