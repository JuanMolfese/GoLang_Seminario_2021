# GoLang_Seminario_2021
## TUDAI 2021
### Grupo Cristian Tisera y Juan Molfese

Objetivo: Crear una función que dada una cadena con un formato determinado genere una instancia de una estructura con los valores de los campos correspondientes al formato de la cadena.

Por ejemplo:  
**TX03ABC**

Deberá generar una estructura con los siguientes valores:

**{TX 3 ABC}**

Donde la estructura deberá tener una definicion del tipo:

type Result struct {

    Type    string     
    Value   string    
    Length  int
}

La cadena de caracteres tiene el siguiente formato:

los primeros **dos** caracteres son el tipo
los segundos **dos** caracteres son el largo del valor
los siguientes **N** caracteres son el valor, donde N es el valor del bullet anterior.

Entonces NN040001 equivale a:

Type = NN
Length = 04
Value = 0001

**Resolución:** 

Generamos un archivo en el que desarrollamos una función que dado un string, genera una estructura. Primero declaramos un struct de tipo "Chain", y dicha función toma el string y asigna la cantidad de caracteres definidos como constantes a cada atributo del struct (Type, Lenght), utilizando esas constantes para establecer indices de comienzo y final de caracteres a asignar a cada atributo.
Luego retornamos el resultado de dicha función y se lo asignamos a una variable del tipo de la estructura.
En el mismo archivo, generamos una función que analiza si el string recibido cumple con los requisitos necesarios para conformar una estructura de tipo "Chain".
Finalmente, generamos un archivo de testing el cual chequea las funciones implementadas utilizando un dataset, provisto por la catedra. El mismo evalua la coincidencia de los caracteres de cada atributo de la estructura, la coincidencia del largo de la cadena del atributo "Value", y que el error sea nulo o no en consecuencia.