#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <locale.h>

// Função que verifica se uma linha corresponde a algum dos padrões fornecidos
int line_matches_patterns(const char *line, const char *patterns) {
    // Se o padrão for vazio ou "all", retorna 1 (corresponde a todas as linhas)
    if (patterns == NULL || strlen(patterns) == 0 || strcmp(patterns, "all") == 0) {
        return 1;  // Retorna 1 para imprimir todas as linhas
    }

    // Caso o padrão não seja vazio ou "all", busca o padrão na linha
    char *pattern_copy = strdup(patterns);
    if (!pattern_copy) {
        fprintf(stderr, "Erro: Falha ao alocar memória para padrões.\n");
        exit(1);
    }

    int match = 0;
    char *token = strtok(pattern_copy, "|");  // Suporta múltiplos padrões separados por '|'
    while (token != NULL) {
        if (strstr(line, token)) {
            match = 1;
            break;
        }
        token = strtok(NULL, "|");  // Avança para o próximo token
    }

    free(pattern_copy);  // Libera a memória alocada para a cópia do padrão
    return match;
}

// Função que remove a quebra de linha do final da string
void remove_newline(char *line) {
    size_t len = strlen(line);
    if (len > 0 && (line[len - 1] == '\n' || line[len - 1] == '\r')) {
        line[len - 1] = '\0';  // Substitui a quebra de linha por '\0'
    }
}

// Função principal que lê o arquivo e verifica as linhas conforme o padrão
void grep_simple_pattern(const char *filename, const char *patterns) {
    FILE *file = fopen(filename, "r");  // Abre o arquivo para leitura
    if (file == NULL) {
        perror("Erro ao abrir o arquivo");
        exit(1);  // Se o arquivo não for encontrado, o programa é encerrado
    }

    char *line = NULL;
    size_t len = 0;
    ssize_t read;

    // Se o padrão for "all", imprime todas as linhas
    if (patterns == NULL || strlen(patterns) == 0 || strcmp(patterns, "all") == 0) {
        while ((read = getline(&line, &len, file)) != -1) {
            remove_newline(line);  // Remove a quebra de linha no final de cada linha
            printf("%s\n", line);  // Imprime todas as linhas
        }
    } else {
        // Caso contrário, imprime apenas as linhas que correspondem ao padrão
        while ((read = getline(&line, &len, file)) != -1) {
            remove_newline(line);
            if (line_matches_patterns(line, patterns)) {
                printf("%s\n", line);  // Exibe a linha que corresponde ao padrão
            }
        }
    }

    free(line);  // Libera a memória alocada para armazenar a linha
    fclose(file);  // Fecha o arquivo após a leitura
}

int main(int argc, char *argv[]) {
    setlocale(LC_ALL, "");  // Suporte para UTF-8 no Windows

    // Verifica se os argumentos foram passados corretamente
    if (argc < 3) {
        printf("Uso: %s <arquivo> <padrao>\n", argv[0]);
        return 1;  // Se os argumentos não forem suficientes, exibe uma mensagem de erro
    }

    const char *filename = argv[1];
    const char *patterns = argv[2];

    grep_simple_pattern(filename, patterns);  // Chama a função que processa o arquivo
    return 0;  // Finaliza o programa
}
