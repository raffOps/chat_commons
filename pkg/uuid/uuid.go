package uuid

import (
	"context"
	"fmt"
	"github.com/edwingeng/wuid/redis/v8/wuid"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

func sanityCheck() {
	variables := []string{"REDIS_HOST", "REDIS_PORT"}
	for _, variable := range variables {
		if _, ok := os.LookupEnv(variable); !ok {
			log.Fatalf("Undefined variable: %s", variable)
		}
	}
}

var newClient = func() (redis.UniversalClient, bool, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{fmt.Sprintf("%s:%s", host, port)}, // Adicione o host corretamente
	})

	// Testa a conexão para verificar se está tudo correto
	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, false, err // Retorna o erro caso haja um problema
	}

	return client, true, nil
}

func NewUUIDGenerator(id string, key string) *wuid.WUID {
	sanityCheck()
	w := wuid.NewWUID(id, nil)
	err := w.LoadH28FromRedis(newClient, key)
	if err != nil {
		panic(err)
	}
	return w
}
