package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/iahta/blog_aggregator/internal/config"
	"github.com/iahta/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

// give access to handlers to the application state, state struct holds pointer to a config
type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	//storing the config in a new instance of the state struct
	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}
	//new instance of the commands struct with an initialized map of handler functions
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	//register the handler with its function
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)

	//os.Args used to get the cl arguments passed in
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}
	//split the arguments to create a command instance
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
