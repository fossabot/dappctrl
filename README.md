# Privatix Controller.

## Building instructions:

1. Install Golang if it's not installed. Make sure that `$HOME/go/bin` is added
to `$PATH`.

2. Build `dappctrl` package:

    ```bash
    DAPPCTRL=github.com/privatix/dappctrl
    DAPPCTRL_DIR=$HOME/go/src/$DAPPCTRL
    mkdir -p $DAPPCTRL_DIR
    git clone git@github.com:Privatix/dappctrl.git $DAPPCTRL_DIR
    go get -d $DAPPCTRL/...
    go get -u gopkg.in/reform.v1/reform
    go generate $DAPPCTRL
    go install $DAPPCTRL
    ```

3. Install PostgreSQL if it's not installed.

4. Prepare a `dappctrl` database instance:

    ```bash
    psql -U postgres -d postgres -c 'CREATE DATABASE dappctrl'
    psql -U postgres -d dappctrl -f $DAPPCTRL_DIR/data/schema.sql
    psql -U postgres -d dappctrl -f $DAPPCTRL_DIR/data/test_data.sql
    ```

5. Modify `dappctrl.config.json` if you need non-default configuration and run:

    ```bash
    dappctrl -config=$DAPPCTRL_DIR/dappctrl.config.json
    ```

6. Build OpenVPN session trigger:

    ```bash
    go install $DAPPCTRL/tool/dapptrig
    ```
