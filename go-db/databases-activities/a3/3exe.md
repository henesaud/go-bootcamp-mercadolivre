# 1. Quais são os 3 principais tipos de culinária (cuisine) que podemos encontrar nos dados? Googlear "mongodb group by field, count it and sort it", Ver a etapa limit do pipeline de agregação.
db.restaurants.aggregate([
    {
        $group: {
            _id: "$tipo_cocina",
            count: { $sum: 1 }
        }
    },
    {
        $sort: { count: -1 }
    },
    {
        $limit: 3
    },
    {
        $project: {
            _id: 0,
            tipo_cocina: "$_id",
            count: 1
        }
    }
])

# 2. Quais são os bairros mais desenvolvidos gastronomicamente? Calcular a média ($avg) da pontuação (grades.score) por bairro; considerando restaurantes com mais de três avaliações. Classifique os bairros com a melhor pontuação.
db.restaurants.aggregate([
    {
        $match: {
            $expr: { $gt: [ { $size: "$grados" }, 3 ] }
        }
    },
    {
        $unwind: "$grados"
    },
    {
        $group: {
            _id: "$barrio",
            avgScore: { $avg: "$grados.puntaje" }
        }
    },
    {
        $sort: {
            avgScore: -1
        }
    }
])

# 3. Uma pessoa querendo comer está na longitude -73.93414657 e na latitude 40.82302903. Quais opções você tem em um raio de 500 metros?
db.restaurants.createIndex({ "direccion.coord": "2dsphere" })
db.restaurants.find({
    "direccion.coord": {
        $near: {
            $geometry: {
                type: "Point",
                coordinates: [-73.93414657, 40.82302903]
            },
            $maxDistance: 500
        }
    }
})