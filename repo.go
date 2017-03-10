package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

type Repo struct {
	Path   string `json:"path"`
	Branch string `json:"branch"`
	URL    string `json:"url"`
	queue  []exec.Cmd
}

func (r Repo) Checkout(branch string) Repo {
	rep := r.new()
	cmd := Git("checkout", branch)
	cmd.Dir = rep.Path
	rep.queue = append(rep.queue, cmd)
	return rep
}

func (r Repo) Clone() Repo {
	rep := r.new()
	rep.queue = append(rep.queue, Git("clone", rep.URL, rep.Path))
	return rep
}

func (r Repo) Reset(to string) Repo {
	rep := r.new()
	cmd := Git("reset", "--hard", to)
	cmd.Dir = rep.Path
	rep.queue = append(rep.queue, cmd)
	return rep
}

func (r Repo) RemoteUpdate(args ...string) Repo {
	rep := r.new()
	args = append([]string{"remote", "update"}, args...)
	cmd := Git(args...)
	cmd.Dir = rep.Path
	rep.queue = append(rep.queue, cmd)
	return rep
}

func (r Repo) Status() Repo {
	rep := r.new()
	cmd := Git("status", "-sb")
	cmd.Dir = rep.Path
	rep.queue = append(rep.queue, cmd)
	return rep
}

func (r Repo) Run() {
	if err := r.Commit(os.Stdin, os.Stdout, os.Stderr); err != nil {
		log.Fatal(err)
	}
}

func (rep Repo) Commit(r io.Reader, w io.Writer, e io.Writer) error {
	for _, c := range rep.queue {
		c.Stdin, c.Stdout, c.Stderr = r, w, e
		// log.Printf("%s %s", c.Path, c.Args)
		if err := c.Run(); err != nil {
			return err
		}
	}
	return nil
}

func (r Repo) new() Repo {
	rep := Repo{
		Path:   r.Path,
		URL:    r.URL,
		Branch: r.Branch,
		queue:  make([]exec.Cmd, len(r.queue)),
	}
	copy(rep.queue, r.queue)
	return rep
}

func RepoExists(r Repo) bool {
	f, err := os.Open(r.Path)
	if err != nil {
		return false
	}
	defer f.Close()

	_, err = f.Stat()
	if err != nil {
		return false
	}

	return true
}

func Git(params ...string) exec.Cmd {
	cmd := exec.Command("git", params...)
	return *cmd
}
