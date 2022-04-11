# Caddy Logs CLI

Caddy Logs CLI is a tool that filters the access logs of [Caddy Server](https://caddyserver.com).
It uses https://github.com/AliAlhajji/caddylogs to filter the logs.

# Install

```
go install github.com/AliAlhajji/caddylogs-cli@latest
```

Or download the source code and build it yourself:
```
go get github.com/AliAlhajji/caddylogs-cli
```

# Usage

```
caddylogs-cli [options] <path_to_log_file>
```

Here I am assuming there is access logs files called access.log

To show all the logs without filtering:
```
caddylogs-cli access.log
```

## Available filters:

- `--first <n>` or `-f <n>`: get the last `n` records of the filtered logs.
```
caddylogs-cli --first 10 access.log
```

- `--last <n>` or `-l <n>`: get the last `n` records of the filtered logs.
```
caddylogs-cli --last 10 access.log
```

- `--info` or `-i`: get only the logs of level `info`
```
caddylogs-cli --info access.log
```

- `--error` or `-e`: get only the logs of level `error`
```
caddylogs-cli --error access.log
```

- `--url-contains <value>` or `-u <value>`: get the logs where requested URL contains `value`:
```
caddylogs-cli -url-contains ".jpg" access.log
```

- `--referer-contains <value>` or `-r <value>`: get the logs where request referer contains `value`:
```
caddylogs-cli --referer-contains "twitter.com" access.log
```

- `--logger-is <value>` or `-g <value>`: get the logs where the used logger is exactly `value`:
```
caddylogs-cli --logger-is "http.log.access.log0" access.log
```

- `--status-code <value>` or `-s <value>`: get the logs where the the repsonse status code is `value`:
```
caddylogs-cli --status-code 404 access.log
```

- `--header <header=value>` or `-x <header=value>`: get the logs where the the request has a `header` equal to `value`:
```
caddylogs-cli --header key1=value1 --header key2=value2  access.log
```
This option can be used multiple times

- `--count` or `-c`: print only the number of logs that match the filters:
```
caddylogs-cli --count access.log
```

- `--reverse`: print the results in reverse (most recent first):
```
caddylogs-cli --reverse access.log
```

You can use a combination of these options to filter the logs more flexible (except for `--info` and `--error`, they cannot be used together):
```
caddylogs-cli --info --url-contains "about.html" --referer-contains "home.html" --first 10 access.log
```
This will return the first 10 logs that:
- are info logs
- have "about.html" in the request url
- have "home.html" in the request's referer.

To get only the number of these logs:
```
caddylogs-cli --info --url-contains "about.html" --referer-contains "home.html" --count access.log
```

# Note

This tool is was created for fun to filter the access logs of my website. It might look simple and primitive, but I find it very useful. I am sharing the code hoping that it will help someone else, or that someone else will find it interesting to extend it with more functionalities.

This tool uses https://github.com/AliAlhajji/caddylogs module to filter the logs. I decided to separate the code into its own module to make it easier for me to re-use in other ideas, such as a Telegram bot that send the filtered lgos on-demmand without having to access the server myself.
