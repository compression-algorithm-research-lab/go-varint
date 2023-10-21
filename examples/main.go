package main

import (
	"fmt"
	"github.com/compression-algorithm-research-lab/go-varint"
)

func main() {

	encode := varint.Encode(uint64(1))
	fmt.Println(encode)

}
