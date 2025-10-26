package cli

import (
	"github.com/MontillaTomas/blog-aggregator/internal/config"
	"github.com/MontillaTomas/blog-aggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
