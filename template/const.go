
// --------------------
str := []string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
str := []string {"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}

func GetAlphabetMap() map[string]int {
	m := make(map[string]int)
	alphabets := []string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	for i:=0; i < 26; i++ {
		m[alphabets[i]] = 0
	}
	return m
}
