package utils

import "testing"

func TestGetHashWithValue(t *testing.T) {
	hash, err := GetHash("123456", "123456")
	if err != nil {
		t.Error(err.Error())
	}
	if hash == "" {
		t.Error("hash value can't be null or empty")
	}
}

func TestGetHashWithEmptySource(t *testing.T) {
	hash, err := GetHash("", "123456")
	if err != nil {
		if err.Error() != "source can't be null or empty" {
			t.Error("Source value is empty so that error value mismatch")
		}
	}
	if hash != "" {
		t.Error("hash should be null or empty")
	}
}

func TestGetHashWithEmptySalt(t *testing.T) {
	hash, err := GetHash("123456", "")
	if err != nil {
		if err.Error() != "salt can't be null or empty" {
			t.Error("Salt value is empty so that error value mismatch")
		}
	}
	if hash != "" {
		t.Error("hash value should be null or empty")
	}
}
