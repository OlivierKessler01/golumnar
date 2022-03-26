# golumnar
A lightweight column-oriented data storage system trying to follow the SQL language standard for the features
it offers.

## Development environments
```bash
docker-composer up -d #Run the docker environment
./tools/run #Compile and run
```

## Configuration
The golumnar main configuration file should be located in `/etc/golumnar/golumnar.cnf` on any Linux systems.

Here are the available configuration values :

- `log_file=<ABSOLUTE_PATH_TO_THE_LOG_FILE>`

## Available commands
- `CREATE [TABLE] <NAME> VALUES (<VALUES>);` Where VALUES are separated by commas.
Create a new table on disk.

- `SELECT <FIELD> FROM <TABLE> WHERE <CONDITION>;`
Queries a file to retrieve data.

## Storage format
Golumnar creates on file per field and supports 2 forms of storage :
- Sequential array-like storage
- B-TREE index-like storage




