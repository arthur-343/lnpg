#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <errno.h>
#include <locale.h>

// Função que verifica se uma linha corresponde a algum dos padrões fornecidos
int line_matches_patterns(const char *line, const char *patterns) {
    // Cria uma cópia do padrão para manipulação (a string de padrões não será modificada)
    char *pattern_copy = strdup(patterns);
    if (!pattern_copy) {
        fprintf(stderr, "Erro: Falha ao alocar memória para padrões.\n");
        exit(1);  // Se não conseguir alocar memória, termina o programa
    }

    int match = 0;
    char *token = strtok(pattern_copy, "|"); // Divide o padrão por '|'
    while (token != NULL) {
        // Verifica se o padrão (token) existe na linha
        if (strstr(line, token)) {
            match = 1;
            break; // Se encontrar uma correspondência, não precisa verificar os outros tokens
        }
        token = strtok(NULL, "|"); // Avança para o próximo token
    }

    free(pattern_copy); // Libera a memória alocada para a cópia do padrão
    return match;
}

// Função que remove a quebra de linha do final da string, se existir
void remove_newline(char *line) {
    size_t len = strlen(line);
    // Se a última posição for '\n' ou '\r' (dependendo do sistema operacional), remove
    if (len > 0 && (line[len - 1] == '\n' || line[len - 1] == '\r')) {
        line[len - 1] = '\0'; // Substitui o '\n' ou '\r' por '\0' (terminador de string)
    }
}

// Função principal que lê o arquivo e verifica as linhas conforme o padrão
void grep_simple_pattern(const char *filename, const char *patterns) {
    // Verificação para garantir que o padrão não seja vazio
    if (strlen(patterns) == 0) {
        fprintf(stderr, "Erro: O padrão fornecido não pode ser vazio.\n");
        exit(1); // Se o padrão for vazio, encerra o programa com erro
    }

    char full_path[1024];
    snprintf(full_path, sizeof(full_path), "%s", filename);  // Obtém o caminho completo do arquivo

    printf("Tentando abrir o arquivo: %s\n", full_path);

    FILE *file = fopen(full_path, "r");  // Tenta abrir o arquivo em modo leitura
    if (file == NULL) {
        // Se não conseguir abrir o arquivo, verifica o erro e fornece uma mensagem apropriada
        if (errno == ENOENT) {
            fprintf(stderr, "Erro: Arquivo '%s' não encontrado.\n", full_path);
        } else if (errno == EACCES) {
            fprintf(stderr, "Erro: Permissão negada para abrir o arquivo '%s'.\n", full_path);
        } else {
            fprintf(stderr, "Erro desconhecido ao abrir o arquivo '%s'.\n", full_path);
        }
        return;
    }

    // Verifica se o arquivo está vazio
    if (fgetc(file) == EOF) {
        fprintf(stderr, "Erro: O arquivo '%s' está vazio.\n", full_path);
        fclose(file);  // Fecha o arquivo antes de retornar
        return;
    }
    rewind(file); // Retorna o ponteiro de leitura ao início do arquivo

    char *line = NULL;
    size_t len = 0;
    ssize_t read;
    int match_found = 0;
    int match_count = 0;

    // Lê o arquivo linha por linha
    while ((read = getline(&line, &len, file)) != -1) {
        // Remove a quebra de linha ('\n' ou '\r\n') que pode estar presente no final da linha
        remove_newline(line);

        // Verifica se a linha corresponde a algum dos padrões fornecidos
        if (line_matches_patterns(line, patterns)) {
            printf("%s\n", line);  // Exibe a linha correspondente ao padrão
            match_found = 1;  // Marca que foi encontrada uma correspondência
            match_count++;  // Conta o número de correspondências
        }
    }
    free(line); // Libera a memória alocada para a linha
    fclose(file);  // Fecha o arquivo após a leitura

    // Exibe o número total de correspondências
    if (match_found) {
        printf("\nForam encontradas %d linha(s) que correspondem ao padrão.\n", match_count);
        exit(0);  // Sucesso
    } else {
        printf("\nNenhuma linha corresponde ao padrão fornecido.\n");
        exit(1);  // Nenhuma correspondência
    }
}

int main(int argc, char *argv[]) {
    setlocale(LC_ALL, "");  // Suporte ao UTF-8 no Windows

    // Verifica se os argumentos foram passados corretamente
    if (argc < 3) {
        printf("Uso: %s <arquivo> <padrao>\n", argv[0]);
        return 1;
    }

    // Chama a função grep_simple_pattern para buscar as linhas no arquivo
    grep_simple_pattern(argv[1], argv[2]);
    return 0;
}
