package ateliersansia

func CompressRLE(s string) string {
	if len(s) == 0 {
		return ""
	}

	compress := ""
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			compress += string(s[i-1]) + itoa(count)
			count = 1
		}
	}
	// Ajouter le dernier caractère et son compteur
	compress += string(s[len(s)-1]) + itoa(count)

	return compress
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		result = string('0'+(n%10)) + result
		n /= 10
	}
	return result
}