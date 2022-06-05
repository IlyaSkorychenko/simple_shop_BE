**Server**
-
build **(from project path)**:

`go build -o simple_shop_build cmd/simpleshop/main.go`

Console args:
 - ``-env_path=[path to .env]``***(Default ``.env`` )***

**CLI**

Build CLI **(from project path)**:

`go build -o cli_build cmd/cli/main.go`

Console args:
- ``-migrate=[up, down]``***(No default)***
