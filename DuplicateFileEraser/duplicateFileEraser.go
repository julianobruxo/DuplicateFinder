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
	"time"

	"github.com/cheggaaa/pb/v3"
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
func excluirArquivos(arquivosDuplicados map[string][]string, numeroExcluir int) {
	for _, arquivos := range arquivosDuplicados {
		if numeroExcluir >= 0 && numeroExcluir < len(arquivos) {
			arquivoExcluir := arquivos[numeroExcluir]
			err := os.Remove(arquivoExcluir)
			if err != nil {
				fmt.Println("Erro ao excluir o arquivo:", err)
			} else {
				time.Sleep(1 * time.Second)
				fmt.Println("Arquivo excluído:", arquivoExcluir)
			}
		}
	}
}
func progressBar() { // barra de progresso
	count := 100

	// Configurar a barra de progresso
	bar := pb.StartNew(count)
	bar.Set("prefix", "Progresso ")
	bar.SetWidth(80)

	for i := 0; i < count; i++ {
		// Simular processamento
		time.Sleep(50 * time.Millisecond)

		// Atualizar a barra de progresso
		bar.Increment()
	}

	// Finalizar a barra de progresso
	bar.Finish()
}
func main() {
	for {
		fmt.Println("Digite o caminho do diretório/subdiretório a ser analisado (ou 's' para sair)")
		inputUsuario := bufio.NewReader(os.Stdin) //recebe o input do diretório
		dirRaiz, erro := inputUsuario.ReadString('\n')
		if erro != nil { // caso o caminho esteja errado ou não exista
			fmt.Println("Erro ao ler o caminho do diretório:", erro)
			time.Sleep(1 * time.Second)
			return // volta pro início da chamada
		}
		dirRaiz = strings.TrimSpace(dirRaiz) // formata o caminho do diretório
		dirRaiz = strings.ReplaceAll(dirRaiz, "\\", "\\\\")

		if dirRaiz == "s" || dirRaiz == "S" {
			time.Sleep(1 * time.Second)
			fmt.Println("Programa encerrado.")
			time.Sleep(1 * time.Second)
			return
		}

		arquivosDuplicados := findDupeFiles(dirRaiz) //faz a leitura dos duplicados
		if len(arquivosDuplicados) == 0 {            //caso o número de dupes não seja 0, segue a função
			fmt.Println("Não foram encontrados arquivos duplicados no diretório:", dirRaiz)
			time.Sleep(1 * time.Second)
		} else {
			fmt.Println("Arquivos duplicados encontrados no diretório:", dirRaiz)

			progressBar()

			for hash, arquivos := range arquivosDuplicados { //recebeu o mapa da função findDupeFiles(dirRaiz) e percorreu um range nele
				fmt.Println("Hash:", hash)
				for i, infoDoArquivo := range arquivos {
					fmt.Printf("%d: %s\n", i, infoDoArquivo)
				}
				fmt.Println()
			}

			fmt.Print("Digite o número do arquivo que deseja excluir (ou 'p' para pesquisar novamente): ")
			inputExcluir := bufio.NewReader(os.Stdin)
			respostaExcluir, _ := inputExcluir.ReadString('\n')
			respostaExcluir = strings.TrimSpace(respostaExcluir)

			if respostaExcluir == "p" || respostaExcluir == "P" {
				continue
			}

			numeroExcluir, erro := strconv.Atoi(respostaExcluir)
			if erro != nil {
				fmt.Println("Número inválido. Operação cancelada.")
				continue
			}

			excluirArquivos(arquivosDuplicados, numeroExcluir)
		}

		fmt.Println("Pesquisa concluída")
		time.Sleep(1 * time.Second)
	}
}
