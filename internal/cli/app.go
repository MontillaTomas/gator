package cli

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/MontillaTomas/blog-aggregator/internal/config"
	"github.com/MontillaTomas/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

func Run() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
		return
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		os.Exit(1)
		return
	}
	defer db.Close()
	dbQueries := database.New(db)

	s := &state{cfg: cfg, db: dbQueries}
	cmds := &commands{handlers: make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogin)
	cmds.register("register", registerHandler)
	cmds.register("reset", resetHandler)
	cmds.register("users", usersHandler)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(addfeedHandler))
	cmds.register("feeds", feedsHandler)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(browseHandler))

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	if err := cmds.run(s, cmd); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
		return
	}
}
