package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// crea un grupo de espera
	wg := &sync.WaitGroup{}

	// agrega dos goroutines al grupo
	wg.Add(2)
	go dibujarRetrato(wg, "LA GIOCONDA")
	go dibujarRetrato(wg, "LA JOVEN DE LA PERLA")

	// espera a que terminen todas las goroutines
	wg.Wait()

	fmt.Println("Fin")
}

func dibujarRetrato(wg *sync.WaitGroup, name string) {
	// indica al grupo que la goroutine ha terminado
	defer wg.Done()

	fmt.Printf("Dibujando retrato %s...\n", name)

	// simula que tarda una cantidad aleatoria de segundos entre 0 y 10
	seconds := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10)
	time.Sleep(time.Duration(seconds) * time.Second)

	fmt.Printf("Retrato %s terminado en %d segundos.\n", name, seconds)
}
