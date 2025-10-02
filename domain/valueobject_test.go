package domain

import "testing"

func TestID(t *testing.T) {

	t.Run("Should create a valid ID.", func(t *testing.T) {

		_, err := NewID()

		if err != nil {
			t.Errorf("I was expecting nil, but I got %s", err.Error())
		}
	})

}

func TestName(t *testing.T) {

	t.Run("Should create a valid name.", func(t *testing.T) {
		name := "Kiala Emanuel"

		_, err := NewName(name)

		if err != nil {
			t.Errorf("I was expecting nil, but I got %s", err.Error())
		}
	})

	t.Run("Should not create an empty name.", func(t *testing.T) {
		name := ""

		_, err := NewName(name)

		if err == nil {
			t.Error("I was expecting an error, but I got nil.")
		}
	})

	t.Run("Should not create an invalid name.", func(t *testing.T) {
		name := "Ki"

		_, err := NewName(name)

		if err == nil {
			t.Error("I was expecting an error, but I got nil.")
		}
	})

	t.Run("Should not create an invalid name with numbers.", func(t *testing.T) {
		name := "Kiala 123"

		_, err := NewName(name)

		if err == nil {
			t.Error("I was expecting an error, but I got nil.")
		}
	})

}

func TestEmail(t *testing.T) {

	t.Run("Should create a valid email.", func(t *testing.T) {
		email := "kiala@gmail.com"

		_, err := NewEmail(email)

		if err != nil {
			t.Errorf("I was expecting nil, but I got %s", err.Error())
		}
	})

	t.Run("Should not create an empty email.", func(t *testing.T) {
		email := ""

		_, err := NewEmail(email)

		if err == nil {
			t.Error("I was expecting an error, but I got nil.")
		}
	})

	t.Run("Should not create an invalid email.", func(t *testing.T) {
		email := "kialagmail"

		_, err := NewEmail(email)

		if err == nil {
			t.Error("I was expecting an error, but I got nil.")
		}
	})
}

func TestPassword(t *testing.T) {

	t.Run("Should create a valid password.", func(t *testing.T) {
		password := "Kiala001"

		_, err := NewPassword(password)

		if err != nil {
			t.Errorf("I was expecting nil, but I got %s", err.Error())
		}
	})

	t.Run("Should not create an invalid password.", func(t *testing.T) {
		password := "Kia"

		_, err := NewPassword(password)

		if err == nil {
			t.Error("I was expecting an error, but I got nil.")
		}
	})

	t.Run("Should not create an empty password.", func(t *testing.T) {
		password := ""

		_, err := NewPassword(password)

		if err == nil {
			t.Error("I was expecting an error, but I got nil.")
		}
	})
}
