# 1) Quantas coleções tem o banco de dados?
1

# 2) Quantos documentos em cada coleção? Quanto ocupa cada coleção?
db.restaurants.countDocuments()
  25359
db.restaurants.dataSize()
  11140976

# 3) Quantos índices em cada coleção? Quanto espaço os índices de cada coleção ocupam?
1 índice/collection.
266.2 kb

# 4) Traga um documento de exemplo de cada coleção. db.collection.find(...).pretty() nos dá um formato mais legível.
Resposta: 

 {
    _id: ObjectId('5eb3d668b31de5d588f42a75'),
    direccion: {
      edificio: '2055',
      coord: [ -74.1321, 40.61266000000001 ],
      calle: 'Victory Boulevard',
      codigo_postal: '10314'
    },
    barrio: 'Staten Island',
    tipo_cocina: 'American',
    grados: [
      {
        date: ISODate('2014-11-06T00:00:00.000Z'),
        grado: 'B',
        puntaje: 25
      },
      {
        date: ISODate('2014-05-06T00:00:00.000Z'),
        grado: 'B',
        puntaje: 20
      },
      {
        date: ISODate('2013-01-26T00:00:00.000Z'),
        grado: 'A',
        puntaje: 13
      },
      {
        date: ISODate('2011-12-17T00:00:00.000Z'),
        grado: 'A',
        puntaje: 7
      }
    ],
    nombre: "Schaffer'S Tavern",
    restaurante_id: '40371771'
  }


# 5) Para cada coleção, liste os campos de nível raíz (ignore campos em documentos aninhados) e seus tipos de dados.
var firstDocument = db.restaurants.findOne();

Object.keys(firstDocument).map(fieldName => ({
    Field: fieldName,
    Type: typeof firstDocument[fieldName]
}));