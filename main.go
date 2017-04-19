package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

const (
	ComposerRepo  = ".compose-repo"
	Manifest      = "manifest.json"
	ComposeFile   = "docker-compose.yml"
	defaultBranch = "master"
	defaultRemote = "origin/master"
)

func unmarshalRepos() []Repo {
	f, err := os.Open(filepath.Join(ComposerRepo, Manifest))

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, _ := ioutil.ReadAll(f)

	var repos []Repo

	if err := json.Unmarshal(data, &repos); err != nil {
		log.Fatal(err)
	}
	return repos
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("===> ")
	app := cli.NewApp()
	app.Name = "composer-repo"
	app.Usage = "sync multiple git repos"
	app.Version = "0.0.6"
	app.Commands = []cli.Command{
		{
			Name:        "init",
			Usage:       "initialize repo with manifest.json and docker-compose.yml",
			Description: "",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "repo,r",
					Usage: "path to manifest repo",
				},
			},
			Action: initCmd,
		},
		{
			Name:        "sync",
			Usage:       "syncronize repos from manifest.json",
			Description: "",
			Action:      syncCmd,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "hard,r",
					Usage: "uses git reset --hard to update",
				},
				cli.BoolFlag{
					Name:  "prune,p",
					Usage: "prune all the remotes that are updated",
				},
			},
		},
		{
			Name:        "fsync",
			Usage:       "alias to sync -r -p",
			Description: "",
			Action:      syncCmd,
			Flags: []cli.Flag{
				cli.BoolTFlag{
					Name:  "hard,r",
					Usage: "uses git reset --hard to update",
				},
				cli.BoolTFlag{
					Name:  "prune,p",
					Usage: "prune all the remotes that are updated",
				},
			},
		},
		{
			Name:        "status",
			Aliases:     []string{"st"},
			Usage:       "show status across all repos",
			Description: "",
			Action:      statusCmd,
		},
	}

	app.Run(os.Args)
}
