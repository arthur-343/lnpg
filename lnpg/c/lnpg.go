package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// removeNewline remove espaços desnecessários e quebras de linha
func removeNewline(line string) string {
	return strings.TrimSpace(line)
}

// matchRegex verifica se uma linha corresponde a um ou mais padrões regex
func matchRegex(line string, patterns []string) (bool, error) {
	// Verifica cada padrão
	for _, pattern := range patterns {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return false, fmt.Errorf("erro ao compilar expressão regular: %w", err)
		}
		if re.MatchString(line) {
			return true, nil // Encontra uma correspondência
		}
	}
	return false, nil // Nenhuma correspondência
}

// grepSimplePattern realiza a busca no arquivo fornecido
func grepSimplePattern(filename string, patterns []string, caseInsensitive, inverse bool) bool {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo %s: %v\n", filename, err)
		return false
	}
	defer file.Close()

	// Adiciona suporte a case insensitive diretamente no padrão
	if caseInsensitive {
		for i := range patterns {
			patterns[i] = "(?i)" + patterns[i]
		}
	}

	scanner := bufio.NewScanner(file)
	matchesFound := false

	for scanner.Scan() {
		line := scanner.Text()
		processedLine := removeNewline(line)

		matches, err := matchRegex(processedLine, patterns)
		if err != nil {
			fmt.Printf("Erro ao processar regex na linha: %s, erro: %v\n", line, err)
			continue
		}

		if (!matches && inverse) || (matches && !inverse) {
			fmt.Println(line)
			matchesFound = true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler o arquivo %s: %v\n", filename, err)
	}

	return matchesFound
}

// searchInDirectory realiza busca recursiva em diretórios
func searchInDirectory(dirPath string, patterns []string, recursive, caseInsensitive, inverse bool) bool {
	dir, err := os.Open(dirPath)
	if err != nil {
		fmt.Printf("Erro ao abrir o diretório %s: %v\n", dirPath, err)
		return false
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Printf("Erro ao listar conteúdos do diretório %s: %v\n", dirPath, err)
		return false
	}

	matchesFound := false
	for _, file := range files {
		if file.Name() == "." || file.Name() == ".." {
			continue
		}

		filePath := dirPath + "/" + file.Name()

		if file.IsDir() && recursive {
			matchesFound = searchInDirectory(filePath, patterns, recursive, caseInsensitive, inverse) || matchesFound
		} else if !file.IsDir() {
			matchesFound = grepSimplePattern(filePath, patterns, caseInsensitive, inverse) || matchesFound
		}
	}
	return matchesFound
}

// main processa os argumentos de entrada e decide o que executar
func main() {
	// Processa os argumentos
	caseInsensitive := false
	inverse := false
	recursive := false
	patterns := []string{}

	// Processa os argumentos para padrões e flags
	for i := 2; i < len(os.Args); i++ {
		if strings.HasPrefix(os.Args[i], "-") {
			switch os.Args[i] {
			case "-r":
				recursive = true
			case "-i":
				caseInsensitive = true
			case "-v":
				inverse = true
			}
		} else {
			patterns = append(patterns, os.Args[i])
		}
	}

	// Verifica se temos padrões de busca
	if len(patterns) == 0 {
		fmt.Println("Por favor, forneça pelo menos um padrão para busca.")
		return
	}

	// Determina se é um arquivo ou diretório
	fileOrDir := os.Args[1]
	fileInfo, err := os.Stat(fileOrDir)
	if err != nil {
		fmt.Printf("Erro ao acessar %s: %v\n", fileOrDir, err)
		return
	}

	matchesFound := false
	if fileInfo.IsDir() {
		matchesFound = searchInDirectory(fileOrDir, patterns, recursive, caseInsensitive, inverse)
	} else {
		matchesFound = grepSimplePattern(fileOrDir, patterns, caseInsensitive, inverse)
	}

	// Define o código de saída
	if matchesFound {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
