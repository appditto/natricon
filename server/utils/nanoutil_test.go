package utils

import "testing"

func TestGenerateAddress(t *testing.T) {
	generated := GenerateAddress()
	if !ValidateAddress(generated) {
		t.Errorf("Invalid address %s", generated)
	}
}

func TestValidateAddress(t *testing.T) {
	// Valid
	valid := "nano_1zyb1s96twbtycqwgh1o6wsnpsksgdoohokikgjqjaz63pxnju457pz8tm3r"
	if !ValidateAddress(valid) {
		t.Errorf("Valid address test failed %s", valid)
	}
	// Invalid
	invalid := "nano_1zyb1s96twbtycqwgh1o6wsnpsksgdoohokikgjqjaz63pxnju457pz8tm3ra"
	if ValidateAddress(invalid) {
		t.Errorf("Valid address returned true when should have been false %s", invalid)
	}
	invalid = "nano_1zyb1s96twbtycqwgh1o6wsnpsksgdoohokikgjqjaz63pxnju457pz8tm3rb"
	if ValidateAddress(invalid) {
		t.Errorf("Valid address returned true when should have been false %s", invalid)
	}
}

func TestPKSha256(t *testing.T) {
	pk := "7fc9064e4d713af2afc73c1527334b665972eb57d65093a378a3e40dbb48ec43"
	hashed := PKSha256(pk, "123456789")
	expected := "fad674ab79c5615a0eb6af3fe763ea892c3bbb589268a2791cbbef9a71a51039"
	if hashed != expected {
		t.Errorf("Expected %s got %s", expected, hashed)
	}
	// Try different seed
	hashed = PKSha256(pk, "987654321")
	expected = "1b51a15f5f37e1a3e5674f445ec0730436cd3292f1cd3a2752307c75d3bb6a1b"
	if hashed != expected {
		t.Errorf("Expected %s got %s", expected, hashed)
	}
}

func TestAddressSha256(t *testing.T) {
	address := "nano_1zyb1s96twbtycqwgh1o6wsnpsksgdoohokikgjqjaz63pxnju457pz8tm3r"
	hashed := AddressSha256(address, "123456789")
	expected := "21562b601912ee6b0e6736f771a5079c33328c2e9d1663050d3205b838f5afb4"
	if hashed != expected {
		t.Errorf("Expected %s got %s", expected, hashed)
	}
	// Try different seed
	hashed = AddressSha256(address, "987654321")
	expected = "633ced6f006dffadbe8b04078cf8328359d39cc17dd1227605989fd7b00e3d04"
	if hashed != expected {
		t.Errorf("Expected %s got %s", expected, hashed)
	}
}

func TestRawToNano(t *testing.T) {
	// 1
	raw := "1000000000000000000000000000000"
	expected := 1.0
	converted, _ := RawToNano(raw)
	if converted != expected {
		t.Errorf("Expected %f but got %f", expected, converted)
	}
	// 1.000001
	raw = "1000001000000000000000000000000"
	expected = 1.000001
	converted, _ = RawToNano(raw)
	if converted != expected {
		t.Errorf("Expected %f but got %f", expected, converted)
	}
	// 1.0000019
	raw = "1000001900000000000000000000000"
	expected = 1.000001
	converted, _ = RawToNano(raw)
	if converted != expected {
		t.Errorf("Expected %f but got %f", expected, converted)
	}
	// 1234.123456789
	raw = "1234123456789000000000000000000000"
	expected = 1234.123456
	converted, _ = RawToNano(raw)
	if converted != expected {
		t.Errorf("Expected %f but got %f", expected, converted)
	}
	// Error
	raw = "1234NotANumber"
	expected = 1234.123456
	_, err := RawToNano(raw)
	if err == nil {
		t.Errorf("Expected error converting %s but didn't get one", raw)
	}
}
