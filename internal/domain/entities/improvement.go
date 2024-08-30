package entities

const PROMPT_IMPROVE_CODE = `
    Mejora los siguientes códigos fuentes para seguir las mejores prácticas : \n%s\n,
    los cuales están en %s. Quiero que tu respuesta esté en el siguiente formato
    
    @@@@NOMBRE ARCHIVO
    CONTENIDO CÓDIGO MEJORADO
    #############################################

    No agregues espacios en blanco después de la última línea de código.
    No coloques comentarios en el código a menos que sea absolutamente necesario.
    No menciones sugerencias ni acciones dentro del código.
`

const PROMPT_ADD_TEST_COVERAGE = `
    Agrega pruebas de cobertura a los siguientes códigos fuentes: \n%s\n,
    los cuales están en %s. Quiero que tu respuesta esté en el siguiente formato
    
    @@@@NOMBRE ARCHIVO_TEST
    CONTENIDO PRUEBAS DE COBERTURA
    #############################################

    No agregues espacios en blanco después de la última línea de código.
    No coloques comentarios en el código a menos que sea absolutamente necesario.
    No menciones sugerencias ni acciones dentro del código.
`
const PROMPT_MIGRACION = `
    Migra los siguientes codigos fuentes : \n%s\n los cuales estan en %s, debes migrar esto a %s,
    quiero que tu respuesta este en el siguiente formato:

    @@@@NOMBRE ARCHIVO
    CONTENIDO CODIGO MIGRADO
    #############################################

    No agregue espacios en blanco despues de la ultima linea de codigo, ya que esto puede generar errores en la migracion.
    No coloques comentarios en el codigo, ya que esto puede generar errores en la migracion.
    No menciones sugerencias ni acciones dentro del codigo, si quieres realizarlo comentalo en formato de comentario.
`
