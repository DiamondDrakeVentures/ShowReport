# ShowReport Migrator

Data migration tool for [ShowReport].

## Usage

Enter the `migrator` directory and build the tool.

``` shell
cd migrator
go build .
```

Run the tool.

``` shell
./migrator -i <INPUT> -s <FORMAT> -o <OUTPUT>
```

## Parameters

| Parameter               | Type     | Default                    | Description                           |
|-------------------------|----------|----------------------------|---------------------------------------|
| `--input`, `-i`         | `string` |                            | **Required.** Path to source file     |
| `--input-format`, `-s`  | `string` |                            | **Required.** Format of source file   |
| `--output`, `-o`        | `string` | `migrated.csv`             | Path to target file                   |
| `--output-format`, `-f` | `string` | `5x`                       | Format of target file                 |
| `--user`, `-u`          | `string` | `migrator@diamonddrake.co` | User attributed for the migrated data |
| `--help`, `-h`          | Flag     |                            | Display the help text                 |
| `--version`             | Flag     |                            | Display the version                   |

[ShowReport]: https://www.appsheet.com/start/636a018f-4165-436b-ba1f-251e72d7205e
