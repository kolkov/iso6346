package iso6346

import "testing"

// Test some valid numbers
func TestValidNums(t *testing.T) {
	validNums := []string{"MRKU4976720", "TCNU6225360", "CMAU5110875", "HDMU6804878", "HDMU6676437"}
	for _, item := range validNums {
		if err := Validate(item); err != nil {
			if err == errValueNotIso6346 {
				t.Error("Valid number validated as invalid", item)
			} else if err != errValueNotIso6346 {
				t.Error("Not a digits")
			}
		}
	}
}

// Test some invalid numbers
func TestInvalidNums(t *testing.T) {
	invalidNums := []string{"MRKU4976721", "TCNU6225365", "CMAU5110870", "SEGU5304843", "458477121"}
	for _, item := range invalidNums {
		if err := Validate(item); err == nil {
			if err == errValueNotIso6346 {
				t.Error("Invalid number validated as valid", item)
			}
		} else if err == errValueLen11 {
			err.Error()
		} else if err != errValueNotIso6346 && err != errValueLen11 {
			t.Error("Other errors!")
		}
	}
}

// Test generating numbers
func TestIso6346(t *testing.T) {
	expectSignature := func(a string, b string) {
		c, err := Generate(a)
		if err != nil {
			t.Errorf("%v", err)
		} else if b != c {
			t.Errorf("for input " + a + " expected signature " + b + " but got " + c + " instead")
		}

		if err := Validate(c); err == nil {
			if err == errValueNotIso6346 {
				t.Errorf("Unable to validate signature that was generated")
			}
		}
	}

	expectSignature("CMAU511087", "CMAU5110875")
}

func BenchmarkIso6346(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Generate("MRKU497672")
	}
}
