
// --------------------
str := []string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}

func GetAlphabetMap() map[string]int {
	m := make(map[string]int)
	alphabets := []string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	for i:=0; i < 26; i++ {
		m[alphabets[i]] = 0
	}
	return m
}