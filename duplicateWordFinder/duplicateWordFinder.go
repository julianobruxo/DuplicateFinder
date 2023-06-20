package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findRepeatedWords(texto string) { // nessa linha, a função findRepeatedWords recebe um valor chamado texto do tipo string)
	inputText := strings.Fields(texto)   // a variável inputText recebe o pacote strings.Fields que, por sua vez, receberá o valor "texto" e irá separar o texto em palavras
	findRepeated := make(map[string]int) // a variável findRepeated recebeu um map com um slice de strings, e irá retornar um inteiro, que será o número de palavras dentro do valor "texto"

	for _, word := range inputText { // o laço for percorre cada palavra do slice inputText (que são as palavras separadas do texto) usando a construção range. O _ é usado para ignorar o índice da iteração, já que não estamos usando ele nesse caso.
		findRepeated[word]++ //para cada palavra word, a linha findRepeated[word]++ incrementa o valor associado à chave word no mapa findRepeated. Isso é feito usando a sintaxe map[chave]++, que incrementa o valor associado à chave em 1.
	} // Basicamente, esse trecho está contando a quantidade de ocorrências de cada palavra no texto. Cada palavra é usada como chave do mapa findRepeated, e o valor associado a essa chave é incrementado a cada ocorrência encontrada.

	if len(findRepeated) == len(inputText) { // retorna o número de chaves em uma varável / slice ou de caracteres em uma string
		fmt.Println("Não foram encontradas palavras repetidas")
	} else {
		for word, counter := range findRepeated { //
			if counter > 1 {
				fmt.Printf("A palavra `%s` aparece %d vezes no texto inserido. \n", word, counter)
			}
		}

	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Digite o texto:")
	texto, _ := reader.ReadString('\n')
	fmt.Println("Palavras repetidas encontradas:")
	findRepeatedWords(texto)
}
