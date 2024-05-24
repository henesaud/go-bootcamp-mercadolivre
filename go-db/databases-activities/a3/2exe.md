# 1. Traga 3 restaurantes que receberem pelo menos uma classificação de grau 'A' com pontuação maior que 20. A mesma classificação deve atender às duas condições simultaneamente
db.restaurants.find({
    "grados": {
        $elemMatch: {
            "grado": 'A',
            "puntaje": { $gt: 20 }
        }
    }
}).limit(3)

# 2. Quantos documentos estão faltando coordenadas geográficas? Em outras palavras, verifique se o tamanho de direccion.coord é 0 e contar.
db.restaurants.count({
    "direccion.coord": { $size: 0 }
})

# 3. Devolver nombre, barrio, tipo_cocina e grados para os 3 primeiros restaurantes; de cada documento apenas a última avaliação.
db.restaurants.aggregate([
    {
        $project: {
            _id: 0,
            nombre: 1,
            barrio: 1,
            tipo_cocina: 1,
            grados: {
                $slice: ["$grados", -1]
            }
        }
    },
    {
        $limit: 3
    }
])