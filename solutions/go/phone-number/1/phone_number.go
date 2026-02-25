package phonenumber
import (
    "regexp"
    "fmt"
)

func Number(phoneNumber string) (string, error) {    
 	// 1. Remove all non-digit characters
	reg := regexp.MustCompile(`\D`)
	cleaned := reg.ReplaceAllString(phoneNumber, "")

	// 2. Remove country code '1' if it starts an 11-digit number
	reCountry := regexp.MustCompile(`^1(\d{10})$`)
	if reCountry.MatchString(cleaned) {
		cleaned = reCountry.ReplaceAllString(cleaned, "$1")
	}

	// 3. Check final length
	if len(cleaned) != 10 {
		return "", fmt.Errorf("invalid length: must be 10 digits (or 11 with country code 1)")
	}

	// 4. Validate NANP NXX rules (N must be 2-9)
	// Area Code (index 0) and Exchange Code (index 3)
	if cleaned[0] < '2' || cleaned[3] < '2' {
		return "", fmt.Errorf("invalid NANP format: area and exchange codes cannot start with 0 or 1")
	}

	return cleaned, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phone, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return phone[:3], nil
}

func Format(phoneNumber string) (string, error) {
	phone, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return fmt.Sprintf("(%s) %s-%s", phone[:3], phone[3:6], phone[6:]), nil
}
