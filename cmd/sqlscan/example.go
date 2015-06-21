package main

type Transaction struct {
	Hello string `db:"HELLO"`
	World string `db:"WORLD"`
}
