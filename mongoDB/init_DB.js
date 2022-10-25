
/// Crear colección
// crear reglas de la collecion para que esto tenga sentido

/// crear collecion libros
db.createCollection(
    "libros",  // nombre de la colleccion           
    {
        validatokr: { // validar campos de la coleccion
            $jsonSchema: {
                bsonType: 'object',
                properties: {
                  titulo: {
                    bsonType: 'string',
                    description: 'nombre del libro'
                  },
                  sinopsis: {
                    bsonType: 'string',
                    description: 'resumen del libro'
                  },
                  autor: {
                    bsonType: 'array',
                    minItems: 1,
                    uniqueItems: true,
                    items: {
                      bsonType: 'objectId'
                    },
                    description: 'autor/autores del libro'
                  },
                  editorial: {
                    bsonType: 'string',
                    description: 'editorial del libro'
                  },
                  paginacion: {
                    bsonType: 'object',
                    required: [
                      'to',
                      'end'
                    ],
                    properties: {
                      to: {
                        bsonType: 'int',
                        description: 'editorial del libro'
                      },
                      end: {
                        bsonType: 'int',
                        description: 'editorial del libro'
                      }
                    }
                  },
                  origen: {
                    bsonType: 'object',
                    required: [
                      'nombre',
                      'url'
                    ],
                    description: '\'items\' must contain the stated fields.',
                    properties: {
                      nombre: {
                        bsonType: 'string',
                        description: 'editorial del libro'
                      },
                      url: {
                        bsonType: 'string',
                        description: 'editorial del libro'
                      }
                    }
                  },
                  path: {
                    bsonType: 'string',
                    description: 'Ruta de ubicacion del libro'
                  },
                  creado: {
                    bsonType: 'date',
                    description: 'Ruta de ubicacion del libro'
                  }
                }
              }
            
        }
    }
)

db.createCollection(
    "autor",  // nombre de la colleccion           
    {
        validator: { // validar campos de la coleccion
            $jsonSchema: { // esquema de la collectio de datoa a evaluar
                bsonType: "object",
                required: ["nombre"], // datos requeridos
                properties: { // propiedades de los campos
                    nombre: {
                        bsonType: "string",
                        description: "nombre del autor"
                    }
                }
            }
        }
    }
)
db.createCollection(
    "coleccion",  // nombre de la colleccion           
    {
        validator: { // validar campos de la coleccion
            $jsonSchema: { // esquema de la collectio de datoa a evaluar
                bsonType: "object",
                required: ["titulo", "libros", "path", "creado"], // datos requeridos
                properties: { // propiedades de los campos
                    titulo: {
                        bsonType: "string",
                        description: "nombre de la colllecion"
                    },
                    sipnosis: {
                        bsonType: "string",
                        description: "resumen del contenido"
                    },
                    libros: {
                        bsonType: "array",
                        minItems: 1, // each box of food color must have at least one color
                        uniqueItems: true,
                        items: {
                            bsonType: ["object"],
                        },
                        description: "lista de _ids de los libros"

                    }, path: {
                        bsonType: "string",
                        description: "ruta de la carpeta de loslibros"
                    },
                    creado: {
                        bsonType: "date",
                        description: "fecha de creacion del elemento"
                    },
                }
            }
        }
    }
)


