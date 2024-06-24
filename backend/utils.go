package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	dbName     string
	dbUser     string
	dbPass     string
	dbHost     string
	dbPort     string
	ListenAddr string
}

func NewEnv() *Env {
	godotenv.Load()
	return &Env{
		dbName:     getEnv("DB_NAME"),
		dbUser:     getEnv("DB_USER"),
		dbPass:     getEnv("DB_PASS"),
		dbHost:     getEnv("DB_HOST"),
		dbPort:     getEnv("DB_PORT"),
		ListenAddr: getEnv("LISTEN_ADDR"),
	}
}

func GetDbConnString(env *Env) string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", env.dbName, env.dbUser, env.dbPass, env.dbHost, env.dbPort)
}

func getEnv(key string) string {
	value, found := os.LookupEnv(key)
	if !found {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}
