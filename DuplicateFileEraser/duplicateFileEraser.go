package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func findDupeFiles(dirRaiz string) map[string][]string {
	mapaComOHashDeCadaArquivo := make(map[string][]string)
	mapaComOsArquivosDuplicados := make(map[string][]string)

	erro := filepath.Walk(dirRaiz, func(path string, info fs.FileInfo, erro error) error {
		if erro != nil {
			return erro
		}
		if !info.IsDir() {
			infoDoArquivo, erro := os.Open(path)
			if erro != nil {
				return erro
			}

			defer infoDoArquivo.Close()

			varParaArmazenarHashCriado := md5.New()
			if _, erro := io.Copy(varParaArmazenarHashCriado, infoDoArquivo); erro != nil {
				return erro
			}
			hashDeCadaArquivo := fmt.Sprintf("%x", varParaArmazenarHashCriado.Sum(nil))
			mapaComOHashDeCadaArquivo[hashDeCadaArquivo] = append(mapaComOHashDeCadaArquivo[hashDeCadaArquivo], path)

			if len(mapaComOHashDeCadaArquivo[hashDeCadaArquivo]) > 1 {
				mapaComOsArquivosDuplicados[hashDeCadaArquivo] = mapaComOHashDeCadaArquivo[hashDeCadaArquivo]
			}
		}
		return nil
	})
	if erro != nil {
		fmt.Println("Erro ao percorrer o diretório:", erro)
		return nil
	}

	return mapaComOsArquivosDuplicados
}

func main() {
	fmt.Println("Digite o caminho do diretório/subdiretório a ser analisado")
	inputUsuario := bufio.NewReader(os.Stdin)
	dirRaiz, erro := inputUsuario.ReadString('\n')
	if erro != nil {
		fmt.Println("Erro ao ler o caminho do diretório", erro)
		return
	}
	dirRaiz = strings.TrimSpace(dirRaiz)
	dirRaiz = strings.ReplaceAll(dirRaiz, "\\", "\\\\")

	arquivosDuplicados := findDupeFiles(dirRaiz)
	if len(arquivosDuplicados) == 0 {
		fmt.Println("Não foram encontrados arquivos duplicados.")
		return
	}

	fmt.Println("Arquivos duplicados encontrados no diretório:", dirRaiz)
	for hash, arquivos := range arquivosDuplicados {
		fmt.Println("Hash:", hash)
		for i, infoDoArquivo := range arquivos {
			fmt.Printf("%d: %s\n", i, infoDoArquivo)
		}
		fmt.Println()
	}

	fmt.Print("Digite o número do arquivo que deseja excluir (ou 's' para sair): ")
	inputExcluir := bufio.NewReader(os.Stdin)
	respostaExcluir, _ := inputExcluir.ReadString('\n')
	respostaExcluir = strings.TrimSpace(respostaExcluir)

	if respostaExcluir == "s" || respostaExcluir == "S" {
		return
	}

	numeroExcluir, erro := strconv.Atoi(respostaExcluir)
	if erro != nil {
		fmt.Println("Número inválido. Operação cancelada.")
		return
	}

	excluirArquivos(arquivosDuplicados, numeroExcluir)
}

func excluirArquivos(arquivosDuplicados map[string][]string, numeroExcluir int) {
	for _, arquivos := range arquivosDuplicados {
		if numeroExcluir >= 0 && numeroExcluir < len(arquivos) {
			arquivoExcluir := arquivos[numeroExcluir]
			err := os.Remove(arquivoExcluir)
			if err != nil {
				fmt.Println("Erro ao excluir o arquivo:", err)
			} else {
				fmt.Println("Arquivo excluído:", arquivoExcluir)
			}
		}
	}
}
