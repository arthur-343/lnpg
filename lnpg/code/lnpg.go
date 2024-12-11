package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	args := os.Args[1:]

	// Verifica se ambos os parâmetros foram fornecidos
	if len(args) < 2 {
		fmt.Printf("Erro: É necessário fornecer o nome do arquivo e o padrão de busca\n")
		fmt.Printf("Uso: %s <arquivo> <padrao>\n", os.Args[0])
		return
	}

	fileName := args[0]
	pattern := args[1]

	// Abre o arquivo
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Erro ao acessar %s: %v\n", fileName, err)
		return
	}
	defer file.Close()

	// Compila o regex do padrão
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Erro ao processar regex para padrão %s: %v\n", pattern, err)
		return
	}

	// Lê o arquivo linha por linha
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Verifica se a linha corresponde ao padrão
		if re.MatchString(line) {
			fmt.Println(line)
		}
	}

	// Verifica se houve erros na leitura do arquivo
	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler o arquivo: %v\n", err)
	}
}
