# Privatix Controller.

## Building instructions:

1. Install Golang if it's not installed.

2. Build `pxctrl` package:

    ```bash
    git clone git@github.com:Privatix/pxctrl.git $HOME/go/src/pxctrl
    cd $HOME/go/src/pxctrl
    go get -d ./...
    go generate
    go install
    ```

3. Install PostgreSQL if it's not installed.

4. Prepare a `pxctrl` database instance:

    ```bash
    psql -U postgres -d postgres -c 'CREATE DATABASE pxctrl'
    psql -U postgres -d pxctrl -f $HOME/go/src/pxctrl/data/schema.sql
    ```

5. Modify `config.json` if you need non-default configuration and run:

    ```bash
    $HOME/go/bin/pxctrl -config=$HOME/go/src/pxctrl/config.json
    ```
