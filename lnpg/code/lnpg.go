package main

import (
	"bufio"  // Pacote para leitura linha por linha de arquivos.
	"fmt"    // Pacote para formatação e impressão de mensagens no terminal.
	"os"     // Pacote para manipulação de arquivos e argumentos da linha de comando.
	"regexp" // Pacote para trabalhar com expressões regulares.
)

func main() {
	// Recebe todos os argumentos passados na linha de comando, exceto o nome do programa
	args := os.Args[1:]

	// Verifica se ao menos o nome do arquivo e o padrão de busca foram fornecidos
	if len(args) < 2 {
		// Se não foram fornecidos ambos, exibe uma mensagem de erro
		fmt.Printf("Erro: É necessário fornecer o nome do arquivo e o padrão de busca\n")
		// Informa o formato correto para o uso do programa
		fmt.Printf("Uso: %s <arquivo> <padrao>\n", os.Args[0])
		return // Sai do programa, já que os parâmetros são obrigatórios
	}

	// O primeiro argumento é o nome do arquivo
	fileName := args[0]
	// O segundo argumento é o padrão de busca (expressão regular ou "all")
	pattern := args[1]

	// Tenta abrir o arquivo com o nome fornecido
	file, err := os.Open(fileName)
	// Se ocorrer um erro ao abrir o arquivo, exibe a mensagem de erro e sai
	if err != nil {
		fmt.Printf("Erro ao acessar %s: %v\n", fileName, err)
		return // Sai do programa após o erro
	}
	// Garante que o arquivo será fechado após a leitura
	defer file.Close()

	// Compila a expressão regular do padrão fornecido
	re, err := regexp.Compile(pattern)
	// Se a expressão regular não for válida, exibe a mensagem de erro e sai
	if err != nil {
		fmt.Printf("Erro ao processar regex para padrão %s: %v\n", pattern, err)
		return // Sai do programa após o erro
	}

	// Cria um scanner para ler o arquivo linha por linha
	scanner := bufio.NewScanner(file)
	// Lê o arquivo linha por linha
	for scanner.Scan() {
		// Obtém a linha atual do arquivo
		line := scanner.Text()

		// Verifica se o padrão é "all"
		if pattern == "all" {
			// Imprime todas as linhas
			fmt.Println(line)
		} else if re.MatchString(line) {
			// Se a linha corresponder ao padrão, imprime a linha
			fmt.Println(line)
		}
	}

	// Verifica se houve algum erro durante a leitura do arquivo
	err = scanner.Err()
	if err != nil {
		// Se houver erro de leitura, exibe a mensagem de erro
		fmt.Printf("Erro ao ler o arquivo: %v\n", err)
	}
}