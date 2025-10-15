package ateliersansia

func CompressRLE(s string) string {
	var a [3]int
	var compress string
	for i := 0; i < len(s); i++{
		if string(s[i]) == "a"{
			a[0] += 1
		} else if string(s[i]) == "b"{
			a[1] += 1
		} else if string(s[i]) == "c"{
			a[2] += 1
		}
	}
	compress = compress + "a" + string(rune(a[0])+48) + "b" + string(rune(a[1])+48) + "c" + string(rune(a[2])+48)
	return compress
}