# Lista enlazadas con golang

## Introducción
una lista enlazada consta de un núero de nodos con dos componentes o campos, un dato o valor que puede ser de cualquier tipo y un apuntador de tipo nodo que apunta al siguiente nodo.
las listas enlazadas se clasifican en:
* Lista simplemente enlazada
* Lista doblemente enlazada
* Lista circular simplemente enlazada
* Lista circular doblemente enlazada

**nota:** En esta librería solo se trataran con listas simplemente enlazada, pero en un futuro se agregaran las demas.

## Pre-requisitos
golang 1.11+
ejecutado exitosamente en golang1.14.10 sobre peppermintOS linux/386

## Instalación

```golang
go get github.com/ing-thejis/list
```

## Ejemplo
```golang

package main

import(
	"fmt"
	"github.com/thejis/go-lsq/list"
)

func main(){
	lista := list.NewList() //Crea una lista nueva y vacia
	lista.Insert(1) //Inserta elementos a la lista de cualquier tipo (int, string, struct, etc)
	lista.Insert(2)
  lista.Insert(3)
	lista.Show()  //muestra elementos de la lista
	fmt.Println(lista.Search(2))  //busca e impreme en pantalla el elemento de valor 2 de la lista
	lista.Remove(3) //Remueve o elimina el elementos con valor 3 de la lista
}
```

# Autor
***The JIS*** <ing.jesithgarcia@gmail.com>
