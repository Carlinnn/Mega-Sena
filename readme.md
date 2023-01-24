# Mega Sena

## O que é esse gerador?
Um gerador de números da Mega-Sena é um programa ou dispositivo que gera automaticamente seis números aleatórios dentro do intervalo de 1 a 60. Esses números gerados são utilizados como aposta na loteria Mega-Sena, na qual os jogadores escolhem seis números para tentar acertar a combinação vencedora. O objetivo de um gerador de números da Mega-Sena é fornecer uma forma fácil e rápida de escolher números para apostar, sem precisar pensar muito ou seguir algum tipo de estratégia.

## Parte técnica
A função "main" é executada quando o programa é iniciado.
Dentro da função "main", a função "generateMegaSenaNumbers" é chamada e armazena o resultado em "numbers". Essa função gera seis números aleatórios sem repetição e os retorna como um slice de inteiros.
Depois disso, a função "sort.Ints" é usada para ordenar os números gerados em ordem crescente.
Por fim, a função "fmt.Println" é usada para imprimir os números gerados e ordenados na tela.
A função "generateMegaSenaNumbers" funciona da seguinte maneira:

Usando a função "rand.Seed" com a data e hora atual, é iniciado um gerador de números aleatórios.
Dentro de um loop, enquanto a quantidade de números gerados for menor que 6, é gerado um número aleatório entre 1 e 60.
Em seguida, é verificado se o número gerado já existe no slice de números gerados anteriormente, usando outro loop.
Se o número gerado ainda não existir no slice, é adicionado ao mesmo.
O loop continua até que sejam gerados 6 números diferentes.
Ao final, a função retorna o slice com os 6 números gerados.