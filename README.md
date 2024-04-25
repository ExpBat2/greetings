#Saludos en Go 

Este paquete es un ejercicio de prueba para obtener saludos personales. 

##Instalacion 

```bash 
go get -u github.com/ExpBat2/greetings

```

#ejemplo de uso

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	names := []string{"Paco", "Javier", "Moni"}
	mes, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(mes)

	mess, err := greetings.Hello("Paco")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mess)
}
