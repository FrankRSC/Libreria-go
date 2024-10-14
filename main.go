package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Book struct {
	Title     string
	Autor     string
	CreatedAt time.Time
}

type Library struct {
	name    string
	address string
	books   []Book
}

type LibraryUpdate struct {
	name    *string
	address *string
}

type DataBase struct {
	Libraries []*Library
}

// aqui agregaremos los libros
//* es un puntero a memoria
//& es el la direccion en la memoria

// * referenciar
// & direccion dela memoria
// * desreferenciar

func main() {
	reader := bufio.NewReader(os.Stdin)
	dataBase := &DataBase{}

	for {
		// Seleccionar la libreria y en base a la libreria hacer las otras opciones, no seleccionar la libreria en cada opcion.
		fmt.Println("----------------------------------------------------")
		fmt.Println("- 1 - Crear libreria")
		fmt.Println("- 2 - Agregar libro a libreria")
		fmt.Println("- 3 - Ver libros")
		fmt.Println("- 4 - Editar un libro")
		fmt.Println("- 5 - Buscar libro")
		fmt.Println("- 6 - Eliminar un libro de una libreria")
		fmt.Println("- 7 - Eliminar una libreria")
		fmt.Println("- 8 - Ver librerias")
		fmt.Println("Para salir - salir")
		fmt.Println("----------------------------------------------------")

		option, _ := reader.ReadString('\n')
		option = DeleteNewLine(option) // eliminar salto de linea
		option = strings.ToLower(option)

		if option == "salir" {
			fmt.Println("Saliendo...")
			break
		}
		switch option {
		//crear libro
		case "1":

			var name, address string

			fmt.Println("Ingresa el nombre de la libreria:")
			name, _ = reader.ReadString('\n')
			name = DeleteNewLine(name) // Eliminar salto de línea

			fmt.Println("Ingresa la direccion:")
			address, _ = reader.ReadString('\n')
			address = DeleteNewLine(address)

			newLibrary, err := NewLibrary(name, address)
			if err != nil {
				fmt.Println(err)
			}

			dataBase.Libraries = append(dataBase.Libraries, newLibrary)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(dataBase.Libraries)

			//agregar libro a libreria
		case "2":
			var bookName, Autor, lib string
			if len(dataBase.Libraries) == 0 {
				fmt.Println("No hay librerias disponibles")
				break
			}
			fmt.Println("A que libreria deseas agregar el libro?")
			for i, library := range dataBase.Libraries {
				fmt.Printf("%v - %s\n", i, library.name)
			}
			lib, _ = reader.ReadString('\n')
			lib = DeleteNewLine(lib)
			nlib, err := strconv.Atoi(lib)

			if err != nil {
				fmt.Println(err)
			}

			library := dataBase.Libraries[nlib]
			fmt.Println("Ingresa el nombre del libro:")
			bookName, _ = reader.ReadString('\n')
			bookName = DeleteNewLine(bookName)
			fmt.Println("Ingresa el nombre del Autor:")
			Autor, _ = reader.ReadString('\n')
			Autor = DeleteNewLine(Autor)

			errB := AddBook(bookName, Autor, time.Now(), library)
			if errB != nil {
				fmt.Println(errB)
			}

		case "3":
			// mostrar las librerias
			// seleccionar libreria y mostrar los libros
			fmt.Println("Selecciona la libreria de la que deseas ver los libros")
			for i, library := range dataBase.Libraries {
				fmt.Printf("%v - %s\n", i, library.name)
			}
			libIndex, _ := reader.ReadString('\n')
			libIndex = DeleteNewLine(libIndex)
			nlib, err := strconv.Atoi(libIndex)
			if err != nil {
				fmt.Println(err)
			}

			books := dataBase.Libraries[nlib].books
			for _, book := range books {
				fmt.Println("------------------------------------------")
				fmt.Printf("Nombre del libro %s\n", book.Title)
				fmt.Printf("Nombre del Autor %s\n", book.Autor)
				fmt.Printf("Creado en %v\n", book.CreatedAt.Format("2006-01-02"))
			}

		case "4":
			// editar un libro
			fmt.Println("Selecciona la libreria de la que deseas ver los libros")
			for i, library := range dataBase.Libraries {
				fmt.Printf("%v - %s\n", i, library.name)
			}
			libIndex, _ := reader.ReadString('\n')
			libIndex = DeleteNewLine(libIndex)
			nlib, err := strconv.Atoi(libIndex)
			if err != nil {
				fmt.Println(err)
			}

			books := dataBase.Libraries[nlib].books
			for ib, book := range books {
				fmt.Printf("----------------- %v -----------------------\n", ib)

				// Obtener el valor y el tipo del struct
				val := reflect.ValueOf(book)
				typ := reflect.TypeOf(book)

				// Iterar sobre los campos del struct Book
				for i := 0; i < val.NumField(); i++ {
					fieldName := typ.Field(i).Name         // Nombre del campo
					fieldValue := val.Field(i).Interface() // Valor del campo

					// Verificar si el campo es de tipo time.Time para formatear la fecha
					switch v := fieldValue.(type) {
					case time.Time:
						// Si fieldValue es de tipo time.Time
						fmt.Printf("%s: %s\n", fieldName, v.Format("2006-01-02"))
					case string:
						// Si fieldValue es de tipo string
						fmt.Printf("%s: %s\n", fieldName, v)
					default:
						// Cualquier otro tipo
						fmt.Printf("%s: %v\n", fieldName, v)
					}
				}

			}
			fmt.Println("Que libro deseas editar?")
			bookIndex, _ := reader.ReadString('\n')
			bookIndex = DeleteNewLine(bookIndex)
			nbook, err := strconv.Atoi(bookIndex)
			if err != nil {
				fmt.Println(err)
			}
			editBook := &books[nbook]

			// Obtener el valor y el tipo del struct
			val2 := reflect.ValueOf(editBook).Elem()
			typ3 := reflect.TypeOf(editBook).Elem()

			// Iterar sobre los campos del struct Book
			for i := 0; i < val2.NumField(); i++ {
				fieldName := typ3.Field(i).Name         // Nombre del campo
				fieldValue := val2.Field(i).Interface() // Valor del campo

				// Verificar si el campo es de tipo time.Time para formatear la fecha
				switch v := fieldValue.(type) {
				case time.Time:
					// Si fieldValue es de tipo time.Time
					fmt.Printf("%s: %s\n", fieldName, v.Format("2006-01-02"))
				case string:
					// Si fieldValue es de tipo string
					fmt.Printf("%s: %s\n", fieldName, v)
				default:
					// Cualquier otro tipo
					fmt.Printf("%s: %v\n", fieldName, v)
				}
			}
			fmt.Println("Que elemento del libro deseas editar?")
			updateField, _ := reader.ReadString('\n')
			updateField = DeleteNewLine(updateField)
			field := val2.FieldByName(updateField)

			fmt.Println("Ingresa el nuevo valor")
			newValue, _ := reader.ReadString('\n')
			newValue = DeleteNewLine(newValue)

			fmt.Println(field)
			if field.IsValid() && field.CanSet() {
				field.SetString(newValue)
			}

		case "5":
			foundBooks := []Book{}
			ch := make(chan Book)

			fmt.Println("Selecciona la libreria de la que deseas buscar")
			for i, library := range dataBase.Libraries {
				fmt.Printf("%v - %s\n", i, library.name)
			}
			libIndex, _ := reader.ReadString('\n')
			libIndex = DeleteNewLine(libIndex)
			nlib, err := strconv.Atoi(libIndex)
			if err != nil {
				fmt.Println(err)
			}

			books := dataBase.Libraries[nlib].books

			fmt.Println("Buscar por nombre: ")
			searchString, _ := reader.ReadString('\n')
			searchString = DeleteNewLine(searchString)

			for _, book := range books {
				//go routine Concurrencia
				go func(b Book) {
					if strings.Contains(b.Title, searchString) {
						ch <- b
					} else {
						ch <- Book{}
					}
				}(book)
			}

			for range books {
				result := <-ch
				if result.Title != "" {
					foundBooks = append(foundBooks, result)
				}
			}
			if len(foundBooks) > 0 {
				for i, book := range foundBooks {
					fmt.Printf("--------------- %v ---------------\n", i)
					fmt.Printf("Nombre del libro %v\n", book.Title)
					fmt.Printf("Autor del libro %v\n", book.Title)
				}
			} else {
				fmt.Println("----------------------------")
				fmt.Println("No se encontraron libros")
				fmt.Println("----------------------------")
			}

		case "6":

			fmt.Println("Selecciona la libreria de la que desea borrar un libro")
			for i, library := range dataBase.Libraries {
				fmt.Printf("%v - %s\n", i, library.name)
			}
			libIndex, _ := reader.ReadString('\n')
			libIndex = DeleteNewLine(libIndex)
			nlib, err := strconv.Atoi(libIndex)
			if err != nil {
				fmt.Println(err)
			}
			books := dataBase.Libraries[nlib].books

			for ib, book := range books {
				fmt.Printf("--------------------%v---------------------\n", ib)
				fmt.Printf("Nombre del libro %s\n", book.Title)
				fmt.Printf("Nombre del Autor %s\n", book.Autor)
				fmt.Printf("Creado en %v\n", book.CreatedAt.Format("2006-01-02"))
			}
			fmt.Println("Que libro deseas borrar: ")
			selectBook, _ := reader.ReadString('\n')
			selectBook = DeleteNewLine(selectBook)
			iSelectedBook, err := strconv.Atoi(selectBook)
			if err != nil {
				fmt.Println(err)
			}

			books = append(books[:iSelectedBook], books[iSelectedBook+1:]...) // desempaquetar

			dataBase.Libraries[nlib].books = books // reasignar la nueva/misma referencia del slice

			fmt.Println(books)
			for ib, book := range books {
				fmt.Printf("--------------------%v---------------------\n", ib)
				fmt.Printf("Nombre del libro %s\n", book.Title)
				fmt.Printf("Nombre del Autor %s\n", book.Autor)
				fmt.Printf("Creado en %v\n", book.CreatedAt.Format("2006-01-02"))
			}

		case "7":
			//Eliminar libreria
			// Guardar dataBase.Libraries en una variable antes de iterarla
			libraries := dataBase.Libraries

			fmt.Println("Selecciona la libreria que deseas eliminar:")
			for i, library := range libraries {
				fmt.Printf("%v - %s\n", i, library.name)
			}

			libIndex, _ := reader.ReadString('\n')
			libIndex = DeleteNewLine(libIndex)
			nlib, err := strconv.Atoi(libIndex)

			if err != nil {
				fmt.Println("Error al convertir el índice:", err)
				break
			}

			if nlib < 0 || nlib >= len(libraries) {
				fmt.Println("libreria invalido")
				break
			}

			libraries = append(libraries[:nlib], libraries[nlib+1:]...)

			dataBase.Libraries = libraries

		case "8":
			if len(dataBase.Libraries) == 0 {
				fmt.Println("No existen librerias")
				break
			}
			for _, library := range dataBase.Libraries {
				fmt.Println("----------------------------------------------------")
				fmt.Printf("Nombre de la libreria %v\n", library.name)
				fmt.Printf("Direccion de la libreria %v\n", library.address)
				// for _, book := range library.books {
				// 	fmt.Printf("Nombre del libro %v\n", book.Title)
				// 	fmt.Printf("Autor del libro %v\n", book.Title)
				// }
			}

		}

	}
}

// func ModificarVariable(v *int) {
// 	*v = 2222 // Modifica el valor al que apunta el puntero v
// }

// func NoModificarVariable(v int) int {
// 	v += 1
// 	return v
// }

func DeleteNewLine(s string) string {
	return strings.TrimRight(s, "\r\n")
}

func NewLibrary(name string, address string) (*Library, error) {
	if name == "" || address == "" {
		return nil, errors.New("falta el nombre o direccion")
	}

	newLibrary := &Library{
		name:    name,
		address: address,
		books:   []Book{},
	}

	return newLibrary, nil

}

func AddBook(Title string, Autor string, CreatedAt time.Time, library *Library) error {

	if Title == "" || Autor == "" || CreatedAt.IsZero() {
		return errors.New("falta el nombre o direccion")
	}

	book := Book{
		Title:     Title,
		Autor:     Autor,
		CreatedAt: CreatedAt,
	}

	library.books = append(library.books, book)

	return nil
}

func UpdateBook(book *Book) error {
	if book == nil {
		return errors.New("el libro no existe")
	}

	book.Title = "UPDATEADO"

	return nil
}

func UpdateLibrary(library *Library, datos LibraryUpdate) error {
	if library == nil {
		return errors.New("la libreria no existe")
	}

	// fmt.Println(*datos.name)

	library.name = *datos.name

	return nil
}
