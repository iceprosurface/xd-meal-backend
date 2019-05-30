# xd-meal-backend

# how to develop

The backend system need env variable to connent DB. 

You should export following code to your 'zshrc' or 'bashrc' file.

```
export DATABASE_USERNAME = "<your database user name>"
export DATABASE_PASSWORD = "<your database user psw>"
export DATABASE_IP       = "<your database ip>"
export DATABASE_PORT     = "<your database port>"
export DATASE_NAME       = "<your database name>"
```


# how to run 

We would like to use `make` as our build tool or the pipeline.


```
# run go server
make run

# build for current os
make build

# build for linux
make build-linux

# clean all build
make clean
```

