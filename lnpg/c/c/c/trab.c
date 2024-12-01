#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <locale.h>
#include <stdbool.h>

// Função que verifica se uma linha corresponde a algum dos padrões fornecidos
int line_matches_patterns(const char *line, const char *patterns) {
    if (patterns == NULL || strlen(patterns) == 0 || strcmp(patterns, "all") == 0) {
        return 1;
    }

    char *pattern_copy = strdup(patterns);
    if (!pattern_copy) {
        fprintf(stderr, "Erro: Falha ao alocar memória para padrões.\n");
        exit(1);
    }

    int match = 0;
    char *token = strtok(pattern_copy, "|");
    while (token != NULL) {
        if (strstr(line, token)) {
            match = 1;
            break;
        }
        token = strtok(NULL, "|");
    }

    free(pattern_copy);
    return match;
}

// Função que remove a quebra de linha do final da string
void remove_newline(char *line) {
    size_t len = strlen(line);
    if (len > 0 && (line[len - 1] == '\n' || line[len - 1] == '\r')) {
        line[len - 1] = '\0';
    }
}

// Função para realizar a busca normal ou invertida
void grep_pattern(const char *filename, const char *patterns, bool invert_search) {
    FILE *file = fopen(filename, "r");
    if (file == NULL) {
        perror("Erro ao abrir o arquivo");
        exit(1);
    }

    char *line = NULL;
    size_t len = 0;
    ssize_t read;

    while ((read = getline(&line, &len, file)) != -1) {
        remove_newline(line);
        int match = line_matches_patterns(line, patterns);
        if ((match && !invert_search) || (!match && invert_search)) {
            printf("%s\n", line);
        }
    }

    free(line);
    fclose(file);
}

int main(int argc, char *argv[]) {
    setlocale(LC_ALL, "");

    if (argc < 3) {
        printf("Uso: %s <arquivo> <padrao> [-v]\n", argv[0]);
        return 1;
    }

    const char *filename = argv[1];
    const char *patterns = argv[2];

    bool invert_search = false;

    // Verifica se o argumento "-v" foi passado
    if (argc == 4 && strcmp(argv[3], "-v") == 0) {
        invert_search = true;
    }

    grep_pattern(filename, patterns, invert_search);

    return 0;
}
