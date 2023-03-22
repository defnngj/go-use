package base

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	tearDownAll := setUpAll()

	code := m.Run()

	tearDownAll() // you cannot use defer tearDownAll()
	os.Exit(code)
}

func setUpAll() func() {
	log.Printf("============setUpAll============")
	return func() {
		log.Printf("============tearDownAll============")
	}
}

func setUp(t *testing.T) func(t *testing.T) {
	log.Printf("-----------setUp-----------")

	return func(t *testing.T) {
		log.Printf("-----------tearDown-----------")
	}
}

func TestAdditon1(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown(t)
	log.Printf("test: TestAdditon1")
	num1 := 1
	num2 := 2
	result := 3
	ret := Addition(num1, num2)
	if ret != result {
		t.Errorf("Addition(%d, %d) != %d", num1, num2, result)
	}

}

func TestAdditon2(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown(t)
	log.Printf("test: TestAdditon1")
	num1 := 10
	num2 := 22
	result := 32
	ret := Addition(num1, num2)
	if ret != result {
		t.Errorf("Addition(%d, %d) != %d", num1, num2, result)
	}

}
