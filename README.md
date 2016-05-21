# PGit

[![Build Status](https://travis-ci.org/mrrooijen/pgit.svg)](https://travis-ci.org/mrrooijen/pgit)

Lightweight command-line utility for managing private repositories on a server.


### Platforms

- Darwin (MacOSX)
- Linux


### Requirements

- Linux Server
  - i.e. $5/mo [Digital Ocean] node (I'm not affiliated with DO)


### Installation

Grab the [latest release] for your platform, and place the `pgit` executable
in a folder available in your `$PATH`. Next, add the `PGIT_URL` environment
variable to your `$HOME/.bashrc` or similar file.

    export PGIT_URL=<user>@<host>[:port]


### Setup Server

SSH into your server as root (or as a user with sudo privileges).

Install [git](http://git-scm.com/book/en/v2/Getting-Started-Installing-Git) on the server.

Create a `git` user:

```sh
sudo adduser \
  --home /home/git \
  --shell /bin/bash \
  --disabled-password \
  git
```

Prepare the `authorized_keys` file:

```sh
sudo mkdir /home/git/.ssh
sudo touch /home/git/.ssh/authorized_keys
sudo chmod 700 /home/git/.ssh
sudo chmod 600 /home/git/.ssh/authorized_keys
sudo chown -R git:git /home/git/.ssh
```

Now store your `$HOME/.ssh/id_rsa.pub` in `/home/git/.ssh/authorized_keys`.


### Usage

Run `pgit help` to view the help screen:

```
Lightweight command-line utility for managing private repositories on a server.

Usage:

  pgit command [arguments]

Commands:

  list                         list all repositories
  clip                         clip the url of a repository to your clipboard
  create <name>                create a new repository
  rename <name> <new_name>     rename an existing repository
  destroy <name>               destroy an existing repository
  version                      display the current pgit version
  help                         display this help screen

https://github.com/mrrooijen/pgit
```


### Backups

Checkout this [guide] that uses [Backup] to perform periodic backup operations
against your repositories.



### Author / License

Released under the [MIT License] by [Michael van Rooijen].

[Michael van Rooijen]: https://twitter.com/mrrooijen
[MIT License]: https://github.com/mrrooijen/pgit/blob/master/LICENSE
[Backup]: https://github.com/mrrooijen/backup
[Digital Ocean]: https://www.digitalocean.com/
[guide]: https://github.com/mrrooijen/pgit/wiki/Backups
[latest release]: https://github.com/mrrooijen/pgit/releases/latest
