package main

import (
	"bufio"         // pacote que recebe Input and Output(io)em Buffer
	"crypto/md5"    // calcula e cria um hash md5 para o arquivo especificado em diretório
	"fmt"           // formatação
	"io"            //pacote que recebe Input and Output(io)
	"io/fs"         //pacote que estende as operações de entrada/saída (io) para incluir funcionalidades específicas do sistema de arquivos (fs = file system)
	"os"            // pacote referente ao operational system
	"path/filepath" // pacote utilizado para lidar com manipulação de caminhos de arquivos e diretórios
	"strings"       // pacote para manipular strings
)

func findDupeFiles(dirRaiz string) map[string][]string {
	mapaComOHashDeCadaArquivo := make(map[string][]string)   // esse mapa vazio vai receber os hashes de cada arquivo lido no dir
	mapaComOsArquivosDuplicados := make(map[string][]string) // esse mapa vazio vai listar os nomes (chaves) dos arquivos que forem encontrados em duplicidade; caso não haja nenhum(erro), retornará mensagem de erro

	erro := filepath.Walk(dirRaiz, func(path string, info fs.FileInfo, erro error) error { //path representa o caminho no sistema do dirRaiz
		if erro != nil {
			return erro //verifica se ocorreu algum erro durante o percurso do diretório. Se houver um erro, ele é retornado imediatamente,
			//encerrando o percurso. Se não houver erro, a função continua a execução normalmente.
		}
		if !info.IsDir() { // se essa info a obtida não for (!) um dir, é um arquivo e pode ser aberto (os.Open)
			infoDoArquivo, erro := os.Open(path) // a var infoDoArquivo armazena a info obtida ao percorrer o filepath
			if erro != nil {
				return erro
			}

			defer infoDoArquivo.Close()

			varParaArmazenarHashCriado := md5.New() //essa fução calculae armazena o hash para cada arquivo lido na var
			if _, erro := io.Copy(varParaArmazenarHashCriado, infoDoArquivo); erro != nil {
				return erro
			}
			hashDeCadaArquivo := fmt.Sprintf("%x", varParaArmazenarHashCriado.Sum(nil)) // imprime os hashes de cada arquivo
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
	for hash, arquivos := range arquivosDuplicados {
		fmt.Println("Arquivos duplicados com o hash", hash+":")
		for _, infoDoArquivo := range arquivos {
			fmt.Println(infoDoArquivo)
		}
	}

}
