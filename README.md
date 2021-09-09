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
