## Questionário

# 1) Como é chamado um JOIN em um banco de dados?

Resposta: Você faz a função entre duas tabelas durante a query.
Por exemplo: select * from users inner join departaments
ON users.department_id = departments.id


# 2) Nomeie e explique 2 tipos de JOIN.

Resposta:

- INNER JOIN: pega as informações de ambas as tabelas, e só traz a linha apenas se
satisfazer a condição do ON.

- LEFT JOIN: pega as informações de ambas as tabelas, mas torna a tabela do JOIN opcional.
Ou seja, traz a linha mesmo se não satisfazer a condição do ON.

# 3) Para que serve o GROUP BY?

Resposta: Agrupar os resultados de acordo com as colunas indicadas

# 4) Para que serve o HAVING?

Resposta: É como se fosse o WHERE, mas afeta os grupos obtidos pelo GROUP BY.

# 5) Dados os diagramas a seguir, indique a qual tipo de JOIN cada um corresponde:

Resposta:

- O primeiro corresponde ao inner join

- O segundo corresponde ao left join

# 6) Escreva uma consulta genérica para cada um dos diagramas abaixo: 

Resposta:

-- primeiro
select * from tableA 
right join tableaB ON tableA.tableb_id = tableaB.id


-- segundo
SELECT * FROM tableA 
LEFT JOIN tableB ON tableA.tableb_id = tableB.id

UNION

SELECT * FROM tableA 
RIGHT JOIN tableB ON tableA.tableb_id = tableB.id