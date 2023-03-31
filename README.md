# crongap

A command-line tool that parses a crontab and reports all cronjobs that are due to run between the two specified timestamps.

My two common use cases are:

1. To determine what cron jobs were likely missed during system down time.
2. To identify "quiet" times, when no cron jobs are running (see "--blanks").

## Installation

`crongap` will run on most Linux, MacOS and Windows systems.

To install it, just `cd` into the directory in which you wish to install it and then copy-paste the appropriate one-liner from below (based on the destination O/S and architecture).

### Linux (32-bit)

```
curl -s -L -o crongap https://github.com/alasdairmorris/crongap/releases/latest/download/crongap-linux-386 && chmod +x crongap
```

### Linux (64-bit)

```
curl -s -L -o crongap https://github.com/alasdairmorris/crongap/releases/latest/download/crongap-linux-amd64 && chmod +x crongap
```

### Mac OS X (Intel)

```
curl -s -L -o crongap https://github.com/alasdairmorris/crongap/releases/latest/download/crongap-darwin-amd64 && chmod +x crongap
```

### Mac OS X (Apple Silicon)

```
curl -s -L -o crongap https://github.com/alasdairmorris/crongap/releases/latest/download/crongap-darwin-arm64 && chmod +x crongap
```

### Windows (32-bit)

```
curl -s -L -o crongap.exe https://github.com/alasdairmorris/crongap/releases/latest/download/crongap-windows-386.exe
```

### Windows (64-bit)

```
curl -s -L -o crongap.exe https://github.com/alasdairmorris/crongap/releases/latest/download/crongap-windows-amd64.exe
```


### Build From Source

If you have Go installed and would prefer to build the app yourself, you can do:

```
go install github.com/alasdairmorris/crongap@latest
```


## Usage

```
A command-line tool that parses a crontab and reports all cronjobs that are due to run between the two specified timestamps.

Usage:
  crongap [-f <crontab>] [-b] <starttime> <endtime>
  crongap -h | --help
  crongap --version

Options:
  -f, --crontab <f>  The crontab file to be parsed [default: -]
  -b, --blanks       Output blank lines for times when no jobs are due
  <starttime>        The start of the time window (format YYYY-MM-DDHH:mm)
  <endtime>          The end of the time window (format YYYY-MM-DDHH:mm)
```

## Examples

```
$ crontab -l | crongap 2022-01-0114:00 2022-01-0114:15
2022-01-01 14:00:00 /usr/local/bin/myscript.sh
2022-01-01 14:05:00 /usr/local/bin/myscript.sh
2022-01-01 14:10:00 /usr/local/bin/myscript.sh
2022-01-01 14:15:00 /usr/local/bin/myscript.sh
```

```
$ crontab -l | crongap --blanks 2022-01-0114:00 2022-01-0114:15
2022-01-01 14:00:00 /usr/local/bin/myscript.sh
2022-01-01 14:01:00
2022-01-01 14:02:00
2022-01-01 14:03:00
2022-01-01 14:04:00
2022-01-01 14:05:00 /usr/local/bin/myscript.sh
2022-01-01 14:06:00
2022-01-01 14:07:00
2022-01-01 14:08:00
2022-01-01 14:09:00
2022-01-01 14:10:00 /usr/local/bin/myscript.sh
2022-01-01 14:11:00
2022-01-01 14:12:00
2022-01-01 14:13:00
2022-01-01 14:14:00
2022-01-01 14:15:00 /usr/local/bin/myscript.sh
```

```
$ cat /etc/cron.d/* | crongap --blanks 2022-01-0114:00 2022-01-0114:15
...
...
```

## License

[MIT](LICENSE)
