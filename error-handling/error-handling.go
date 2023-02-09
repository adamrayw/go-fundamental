package main

import (
	"errors"
	"fmt"
)

func pembagian(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Pembagian tidak boleh 0")
	}

	return a / b, nil
}

func main() {
	a := 10
	b := 3

	// memanggil fungsi pembagian
	c, err := pembagian(a, b)

	if err != nil {
		// jika terjadi error
		fmt.Println(err.Error())
	} else {
		// jika tidak error, tampilkan hasil pembagian
		fmt.Println(c)
	}

}
