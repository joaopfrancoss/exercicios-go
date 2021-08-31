package main

func comparador(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	letraQuantidade := make(map[string]int)
	letraQuantidade2 := make(map[string]int)

	for i := 0; i < len(a); i++ {
		letra := string(a[i])
		_, existe := letraQuantidade[letra]
		if existe {
			letraQuantidade[letra]++
		} else {
			letraQuantidade[letra] = 1
		}
	}

	for i := 0; i < len(b); i++ {
		letra := string(b[i])
		_, existe := letraQuantidade2[letra]
		if existe {
			letraQuantidade2[letra]++
		} else {
			letraQuantidade2[letra] = 1
		}
	}

	for chave := range letraQuantidade {
		if letraQuantidade[chave] != letraQuantidade2[chave] {
			return false
		}
	}
	return true
}
