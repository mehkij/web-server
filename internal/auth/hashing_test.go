package auth

import "testing"

func TestPasswordHashing(t *testing.T) {
	passwords := []string{
		"aReallyStrongPassword2024$",
		"123456789",
		"anormalpassword",
	}

	for _, password := range passwords {
		_, err := HashPassword(password)
		if err != nil {
			t.Fatalf("Error hashing password: %s\n", err)
		}
	}
}

func TestComparison(t *testing.T) {
	passwords := []string{
		"aReallyStrongPassword2024$",
		"123456789",
		"anormalpassword",
	}

	for _, password := range passwords {
		hashedPass, err := HashPassword(password)
		if err != nil {
			t.Fatalf("Error hashing password: %s\n", err)
		}
		err = CheckPasswordHash(password, hashedPass)
		if err != nil {
			t.Fatalf("Error comparing password to hash\n: %s", err)
		}
	}
}
