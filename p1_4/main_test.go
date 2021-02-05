package main

import "testing"

func Test_isNdigit(t *testing.T) {
	if !isNdigit(9001, 4) {
		t.Fatal("9001 は4桁の数字です！")
	}

	if !isNdigit(123456, 6) {
		t.Fatal("123456 は6桁の数字です！")
	}

	if !isNdigit(1234567890, 10) {
		t.Fatal("1234567890 は10桁の数字です！")
	}
}

func Test_nDigit(t *testing.T) {
	r1 := nDigit(123, 2)
	if r1 != 2 {
		t.Fatalf("%d の %d 桁目の数値: %d", 123, 2, r1)
	}

	r2 := nDigit(12345, 5)
	if r2 != 1 {
		t.Fatalf("%d の %d 桁目の数値: %d", 12345, 5, r2)
	}

	r3 := nDigit(1234567890, 10)
	if r3 != 1 {
		t.Fatalf("%d の %d 桁目の数値: %d", 12345, 5, r2)
	}
}
