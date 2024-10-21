package utils

import "testing"

func TestXxx(t *testing.T) {
	bashed, err := HashPassword("12345678")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(bashed)
}

// $2a$10$At5e.mrora18r8rjmAn3Je9Tsw2NKBEtqI.lzJ9XXYNFaKnsBT6uy
