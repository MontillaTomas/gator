package cli

import (
	"fmt"
	"os"

	"github.com/MontillaTomas/blog-aggregator/internal/config"
)

func Run() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
		return
	}

	s := &state{cfg: cfg}
	cmds := &commands{handlers: make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(1)
		return
	}

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
