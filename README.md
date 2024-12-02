### Passo Zero

Como na maioria das linguagens de programação, começamos com índice zero!

Para este passo, deixarei você configurar seu IDE/editor de escolha e a linguagem de programação de escolha. Depois disso, aqui está o que eu gostaria que você fizesse para estar pronto para testar sua solução.

Baixe o seguinte texto: [https://www.gutenberg.org/cache/epub/132/pg132.txt](https://www.gutenberg.org/cache/epub/132/pg132.txt) e salve-o como `test.txt`. Depois disso, baixe os dados de teste adicionais do meu Dropbox aqui e descompacte-os, você deve ter o seguinte em seu diretório de teste:

```sh
% tree
.
├── rockbands.txt
├── symbols.txt
├── test-subdir
│   └── BFS1985.txt
└── test.txt
```

Se você estiver no Windows, pode usar GoW ou GitBash para obter um console bash, ou se você for um usuário do PowerShell, tenho certeza de que você é capaz de traduzir os comandos! Bem melhor do que eu! 😇

### Passo 1

Neste passo, seu objetivo é implementar suporte para uma expressão vazia. Uma expressão vazia corresponde a todas as linhas, então o comando `grep "" test.txt` escreverá todas as linhas do arquivo `test.txt`.

Implemente sua solução, execute o comando e verifique se todas as linhas são escritas. Você pode automatizar este teste com as ferramentas de linha de comando Unix:

```sh
 grep " " test.txt | diff test.txt -
%
```

Isso mostra que a saída do seu `grep` é igual ao conteúdo do arquivo `test.txt` porque não há diferença na saída. Se o seu `grep` não estiver correto para este passo, haverá uma saída de diferença.

### Passo 2

Neste passo, seu objetivo é corresponder a um padrão simples de uma letra e retornar o código de saída correto para o shell. Quando um padrão é correspondido, o `grep` deve retornar o código de saída zero e não zero quando nenhum padrão é correspondido.

```sh
 grep J rockbands.txt
Judas Priest
Bon Jovi
Junkyard
```

Você pode verificar o código de retorno usando `echo $?`.

### Passo 3

Neste passo, seu objetivo é percorrer uma árvore de diretórios, ou seja, suportar a opção de linha de comando `-r`.

Então, seu caso de teste para este passo é:

```sh
 grep -r Nirvana *
rockbands.txt:Nirvana
test-subdir/BFS1985.txt:Since Bruce Springsteen, Madonna, way before Nirvana
test-subdir/BFS1985.txt:On the radio was Springsteen, Madonna, way before Nirvana
test-subdir/BFS1985.txt:And bring back Springsteen, Madonna, way before Nirvana
test-subdir/BFS1985.txt:Bruce Springsteen, Madonna, way before Nirvana
```

### Passo 4

Neste passo, seu objetivo é implementar a opção `-v`. Isso inverte a busca, excluindo qualquer resultado que corresponda. Se não gostamos de Madonna, podemos fazer isso:

```sh
 grep -r Nirvana * | grep -v Madonna
rockbands.txt:Nirvana
```

Encontrando todos os primeiros resultados que não incluem Madonna.

### Passo 5

Neste passo, seu objetivo é suportar `\d` e `\w` no padrão de busca. Seus significados são:

- `\d` - um dígito.
- `\w` - um caractere de palavra.

Use os seguintes dois casos de teste para verificar sua implementação:

```sh
 grep "\d" test-subdir/BFS1985.txt
Her dreams went out the door when she turned 24
There was U2 and Blondie, and music still on MTV
'Cause she's still preoccupied with 19, 19, 1985, 1985
There was U2 and Blondie, and music still on MTV
'Cause she's still preoccupied with 19, 19, 1985
There was U2 and Blondie, and music still on MTV
'Cause she's still preoccupied with 1985
There was U2 and Blondie, and music still on MTV
'Cause she's still preoccupied with 19, 19, 1985

 grep "\w" symbols.txt
pound
dollar
```

### Passo 6

Neste passo, seu objetivo é implementar suporte para correspondência `^` no início de uma linha e `$` no final.

Você pode testar com:

```sh
 grep ^A rockbands.txt
AC/DC
Aerosmith
Accept
April Wine
Autograph
```

e:

```sh
 grep na$ rockbands.txt
Nirvana
```

### Passo Final

Neste passo, seu objetivo é suportar o argumento opcional de linha de comando `-i`, para que você suporte busca insensível a maiúsculas e minúsculas:

```sh
 grep A rockbands.txt | wc -l
8

 grep -i A rockbands.txt | wc -l
58
```

Uma vez que você obtenha o resultado acima, parabéns! Você conseguiu, dê um tapinha nas costas, trabalho bem feito!





