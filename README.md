# composer-repo

## Install

    $ brew install https://raw.githubusercontent.com/kron4eg/compose-repo/master/compose-repo.rb

Or simly:

    $ go get -u github.com/kron4eg/compose-repo

## manifest file format

    [
        {
            "path": "localpath",
            "branch": "custom_branch_other_then_master",
            "url": "git@github.com:project/repo.git"
        }
    ]

## Usage

    $ compose-repo init -r git@github.com:project/manifest-repo.git
    $ compose sync -r

This will symlink `docker-compose.yml` file from manifest-repo.
