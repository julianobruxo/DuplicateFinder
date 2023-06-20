package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func findDuplicatedFiles(rootDir string) map[string][]string {
	fileHashes := make(map[string][]string) // Aqui estamos criando um mapa vazio chamado fileHashes.
	//Esse mapa será usado para armazenar as chaves (hashes) e os valores (nomes de arquivos duplicados) relacionados.
	duplicateFiles := make(map[string][]string) // essa variável armazena o mapacom a lista de nomes dos arquivos duplicados

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error { //Utilizamos a função filepath.Walk para percorrer todos os arquivos e pastas dentro do diretório especificado.
		//A função recebe o caminho raiz (rootDir) e uma função anônima ==>  func(path string, info os.FileInfo, err error) <== que será executada para cada arquivo ou pasta encontrado.
		//Ou seja, a cada arquivo lido pelo "info os.FileInfo", a func anonima executará as verificações IF abaixo
		if err != nil { // Aqui o compilador verifica se há algum erro. Caso err seja diferente de nil, ou seja, se houver um erro, o compilador para imediatamente.
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			hash := md5.New()
			if _, err := io.Copy(hash, file); err != nil {
				return err
			}

			fileHash := fmt.Sprintf("%x", hash.Sum(nil))
			fileHashes[fileHash] = append(fileHashes[fileHash], path)

			if len(fileHashes[fileHash]) > 1 {
				duplicateFiles[fileHash] = fileHashes[fileHash]
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Erro ao percorrer o diretório:", err)
		return nil
	}

	return duplicateFiles
}

func main() {
	fmt.Println("Digite o caminho do diretório a ser verificado:")
	reader := bufio.NewReader(os.Stdin)
	rootDir, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler o caminho do diretório:", err)
		return
	}

	rootDir = strings.TrimSpace(rootDir)
	rootDir = strings.ReplaceAll(rootDir, "\\", "\\\\")

	duplicates := findDuplicatedFiles(rootDir)
	if len(duplicates) == 0 {
		fmt.Println("Nenhum arquivo duplicado encontrado.")
		return
	}

	for hash, files := range duplicates {
		fmt.Println("Arquivos duplicados com o hash", hash+":")
		for _, file := range files {
			fmt.Println(file)
		}

	}
}
