package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	args := os.Args[1:]

	// Verifica se foi passado o nome do arquivo e ao menos um padrão
	if len(args) < 2 {
		fmt.Println("Erro: Informe um arquivo e ao menos um padrão")
		return
	}

	fileName := args[0]
	patterns := args[1:]

	// Abre o arquivo de entrada
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Erro ao acessar %s: %v\n", fileName, err)
		return
	}
	defer file.Close()

	// Um map para evitar imprimir a mesma linha mais de uma vez
	printedLines := make(map[string]bool)

	// Lê o arquivo linha por linha
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Verifica se a linha corresponde a algum dos padrões
		matches := false
		for _, pattern := range patterns {
			re, err := regexp.Compile(pattern)
			if err != nil {
				fmt.Printf("Erro ao processar regex: %v\n", err)
				continue
			}
			if re.MatchString(line) {
				matches = true
				break
			}
		}

		// Se a linha corresponde a algum padrão e ainda não foi impressa, imprima-a
		if matches && !printedLines[line] {
			fmt.Println(line)
			printedLines[line] = true
		}
	}

	// Verifica se ocorreu algum erro na leitura do arquivo
	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler o arquivo: %v\n", err)
	}
}
