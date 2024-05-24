## Questionário

# 1. Explique o conceito de normalização e porque ele é usado.

Resposta: A normalização padroniza e valida os dados, a fim de eliminar redundâncias e inconsistências, garantindo assim a integridade dos dados. Possui 3 níveis de normalização.

- 1º forma normal: elimina dados duplicados em atributos

- 2º forma normal: remove colunas que não dependem da chave primária

- 3º forma normal: Elimina dependências transitivas. Ou seja, nenhum atributo não-chave deve depender de outro atribut não-chave, todos devem depender diretamente da chave primária.


# 11) O que são índices? Para que servem?

É um mecanismo para otimizar consultas, do qual melhora o acesso aos dados, fornecendo um caminho mais direto para os registros. 
Existem 3 tipos de índice:

- Chave primária: não permite duplicados e se refere ao identificador único da tabela.
- Comum: permite duplicados
- únicos: não se refere ao identificador da tabela, por isso não é primário. Porém, também não permite registros duplicados.