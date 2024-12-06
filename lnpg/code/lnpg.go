package main

import (
	"bufio"  // Para leitura de arquivos linha a linha
	"fmt"    // Para exibir mensagens no console
	"os"     // Para manipular arquivos e acessar argumentos de linha de comando
	"regexp" // Para trabalhar com expressões regulares
)

func main() {
	// Verifica se há argumentos suficientes
	if len(os.Args) < 3 {
		fmt.Println("Erro: Informe um arquivo e ao menos um padrão")
		return
	}

	// Pega o nome do arquivo e os padrões da linha de comando
	fileName := os.Args[1]
	patterns := os.Args[2:]

	// Abre o arquivo para leitura
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo %s: %v\n", fileName, err)
		return
	}
	defer file.Close()

	// Lê o arquivo e processa as linhas
	err = processFile(file, patterns)
	if err != nil {
		fmt.Printf("Erro ao processar o arquivo: %v\n", err)
	}
}

// processFile lê o arquivo linha a linha e verifica cada padrão
func processFile(file *os.File, patterns []string) error {
	// Mapa para garantir que cada linha seja impressa no máximo uma vez
	alreadyPrinted := make(map[string]bool)

	// Lê o arquivo linha por linha
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Verifica se a linha corresponde a algum dos padrões
		if matchesAnyPattern(line, patterns) {
			// Imprime a linha apenas se ainda não foi impressa
			if !alreadyPrinted[line] {
				fmt.Println(line)
				alreadyPrinted[line] = true
			}
		}
	}

	// Retorna erro, se ocorrer, na leitura do arquivo
	return scanner.Err()
}

// matchesAnyPattern verifica se uma linha corresponde a pelo menos um dos padrões
func matchesAnyPattern(line string, patterns []string) bool {
	for _, pattern := range patterns {
		// Compila o padrão para regex
		re, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Printf("Erro ao compilar regex '%s': %v\n", pattern, err)
			continue
		}

		// Verifica se a linha corresponde ao padrão
		if re.MatchString(line) {
			return true
		}
	}
	return false
}
