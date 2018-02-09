# Privatix Controller.

## Building instructions:

1. Install Golang if it's not installed. Make sure that `$HOME/go/bin` is added
to `$PATH`.

2. Build `pxctrl` package:

    ```bash
    git clone git@github.com:Privatix/pxctrl.git $HOME/go/src/pxctrl
    go get -d pxctrl/...
    go get -u gopkg.in/reform.v1/reform
    go generate pxctrl
    go install pxctrl
    ```

3. Install PostgreSQL if it's not installed.

4. Prepare a `pxctrl` database instance:

    ```bash
    psql -U postgres -d postgres -c 'CREATE DATABASE pxctrl'
    psql -U postgres -d pxctrl -f $HOME/go/src/pxctrl/data/schema.sql
    psql -U postgres -d pxctrl -f $HOME/go/src/pxctrl/data/test_data.sql
    ```

5. Modify `config.json` if you need non-default configuration and run:

    ```bash
    $HOME/go/bin/pxctrl -config=$HOME/go/src/pxctrl/pxctrl.config.json
    ```

6. Build OpenVPN session trigger:

    ```bash
    go install pxctrl/tool/pxtrig
    ```
