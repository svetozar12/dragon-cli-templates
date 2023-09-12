package main

import (
	"@dragon-cli-template/apps/api/internal/bootstrap"

	_ "github.com/lib/pq"
)

func main() {
	bootstrap.Bootstrap()
}
