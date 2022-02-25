# STM

Super Test Mail generator

You give it an email, it will date and uuid that email for you, in case you need to sign up for lots of user accounts when testing your stuff.

## usage

```bash
go get
go run . jonathan@example.com

# output
# jonathan+2022-02-25+ed4001f3-7536-49bf-ad13-266ac7ad1ca5@example.com
```

on Mac OS you can pipe it to your clipboard like:

```bash
go run . jonathan@example.com | pbcopy
```

## test

```bash
go test
```

You can also benchmark the function for an amount of times:

```bash
go test -bench=. --benchtime=1000000
```
