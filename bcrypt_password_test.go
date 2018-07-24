package go_password

import "testing"

func TestSaltIsGenerated(t *testing.T) {
	left, err := NewBcryptPassword("test")
	if err != nil {
		t.Error("error while creating password")
	}
	right, err := NewBcryptPassword("test")
	if err != nil {
		t.Error("error while creating password")
	}

	if left.GetSalt() == right.GetSalt() {
		t.Error("left and right salt are the same")
	}
}

func TestHashIsGenerated(t *testing.T) {
	left, err := NewBcryptPassword("test")
	if err != nil {
		t.Error("error while creating password")
	}
	right, err := NewBcryptPassword("test")
	if err != nil {
		t.Error("error while creating password")
	}

	if left.GetHash() == right.GetHash() {
		t.Error("left and right hash are the same")
	}
}

func TestPasswordMayBeCompared(t *testing.T) {
	password, err := NewBcryptPassword("test")
	if err != nil {
		t.Error("error while creating password")
	}
	if !password.Compare("test") {
		t.Error("password coult not be compared correctly")
	}
}

func TestCustomSaltIsUsed(t *testing.T) {
	left, err := NewBcryptPasswordWithGivenSalt("test", "salt")
	if err != nil {
		t.Error("error while creating password")
	}
	right, err := NewBcryptPasswordWithGivenSalt("test", "salt")
	if err != nil {
		t.Error("error while creating password")
	}

	if left.GetSalt() != right.GetSalt() {
		t.Error("left and right salt are not the same")
	}
}

func TestBcryptHasSaltedHash(t *testing.T) {
	left, err := NewBcryptPasswordWithGivenSalt("test", "salt")
	if err != nil {
		t.Error("error while creating password")
	}
	right, err := NewBcryptPasswordWithGivenSalt("test", "salt")
	if err != nil {
		t.Error("error while creating password")
	}

	if left.GetHash() == right.GetHash() {
		t.Error("left and right hash are the same")
	}
}
