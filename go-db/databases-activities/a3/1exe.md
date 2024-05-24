# 1. Devolver restaurante_id, nombre, barrio e tipo_cocina mas excluindo _id para um documento (o primeiro)
db.restaurants.find({},{_id:0, grados:0, direccion:0})

# 2. Devolver restaurant_id, nombre, barrio e tipo_cocina para os primeiros 3 restaurantes que contenham 'Bake' em alguma parte do seu nome.
db.restaurants.find({nombre: /Bake/}, {restaurant_id: 1, nombre: 1, barrio: 1, tipo_cocina: 1}).limit(3)

# 3. Contar os restaurantes de comida (tipo_cocina) china (Chinese) e tailandesa (Thai) do bairro (bairro) Bronx.
db.restaurants.count({tipo_cocina: {$in: ["Chinese", "Thai"]}, barrio: "Bronx"})