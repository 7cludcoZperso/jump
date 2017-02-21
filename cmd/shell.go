package cmd

import (
	"os"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/shell"
)

func shellCmd(args cli.Args, _ *config.Config) {
	hint := args.CommandName()
	if len(hint) == 0 {
		hint = os.Getenv("SHELL")
	}

	sh := shell.Guess(hint)
	shortcut := args.Get("--shortcut", "j")

	cli.Outf("%s", sh.MustCompile(shortcut))
}

func init() {
	cli.RegisterCommand("shell", "Display a shell integration script.", shellCmd)
}
