package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("nombre vacio")
	}
	message := fmt.Sprintf(randForm(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	mes := make(map[string]string)

	for _, name := range names {
		mess, err := Hello(name)
		if err != nil {
			return nil, err
		}
		mes[name] = mess
	}
	return mes, nil
}

func randForm() string {
	forms := []string{
		"Hola, %v. Bienvenido :D",
		"Que bueno verte %v :3",
		"Pasala bien %v, disfruta tu dia ;)",
	}

	return forms[rand.Intn(len(forms))]
}
