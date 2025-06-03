# leetcodegrind
## two sum
Recebe array de números e um número alvo, deve percorrer o array e encontrar e retornar os índices dos dois números que somados são igual ao número alvo.
Resolvido com map como auxiliar para guardar os números percorridos e com variável complemento para identificar qual o número que falta para atingir o alvo em cada iteração.

## add two numbers
Recebe duas linked lists, tendo cada nó um digito, representando um número ao contrário, e deve somar cada dígito de cada ll para retornar outra ll com o resultado final contendo os dígitos em cada nó.
2 -> 4 -> 3 representa 342
Resolvido usando duas lls, uma dummy que inicia vazia e representa o ponto inicial e outra current que representa a lista de retorno real em que será adicionada a soma dos valores, e usando variável carry para passar a dezena para a próxima casa decimal caso a soma dê acima de 9.
