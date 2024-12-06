package main

import (
	"bufio"  // Para leitura de arquivos linha a linha
	"fmt"    // Para exibir mensagens no console
	"os"     // Para manipular arquivos e acessar argumentos de linha de comando
	"regexp" // Para trabalhar com expressões regulares
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

	// Para cada padrão fornecido
	for _, pattern := range patterns {
		// Se o padrão for " ", ajusta para corresponder a todas as linhas
		if pattern == " " {
			pattern = ".*" // Regex que corresponde a qualquer linha
		}

		// Compila o regex do padrão
		re, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Printf("Erro ao processar regex para padrão %s: %v\n", pattern, err)
			continue
		}

		// Usamos um mapa para garantir que cada linha seja impressa no máximo uma vez para o padrão
		printed := make(map[string]bool)

		// Reseta o scanner para percorrer o arquivo novamente para cada padrão
		file.Seek(0, 0)
		scanner := bufio.NewScanner(file)

		// Lê o arquivo linha por linha para esse padrão
		for scanner.Scan() {
			line := scanner.Text()

			// Verifica se a linha corresponde ao padrão atual
			if re.MatchString(line) {
				// Se a linha ainda não foi impressa, imprime e marca como impressa
				if !printed[line] {
					fmt.Println(line)
					printed[line] = true
				}
			}
		}

		// Verifica se ocorreu algum erro na leitura do arquivo
		if err := scanner.Err(); err != nil {
			fmt.Printf("Erro ao ler o arquivo: %v\n", err)
		}
	}
}
