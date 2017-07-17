
# pword

Generate secure passwords. 

# Overview

`pword` is a utility to generate strong passwords that are memorable. 

```
$ pword online
wipeout doorstep copier announcer
pantyhose drainpipe wikipedia glycerin
remote receipt aidless ladder
makeover tattoo generator yogurt
```

It is based on the concept behind [XKCD 936](https://xkcd.com/936/) and [XKCD-password-generator](https://github.com/redacted/XKCD-password-generator). It makes use of [EFF's typo-tolerant wordlist](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases) to provide live autocomplete.

# CLI Usage

```
$ pword help
NAME:
   pword - generate secure passwords

USAGE:
   pword [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     online   Generates passwords for use on websites
     offline  Generates passwords for use offline (laptops, encrypted drives)
     crypto   Generates extremely secure passwords
     recall   Utility with autocomplete to help you recall passwords
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --count NUM, -c NUM  Generates NUM passwords for you to choose from (default: "auto")
   -1                   Equivalent to --count 1
   --stronger           Chooses from a list of 7,776 words instead
   --help, -h           show help
   --version, -v        print the version
```

# Library Usage

WIP

```
import "github.com/serverwentdown/pword/pw"
```

# Modes

## Online

This mode generates a 4-word password from the 1,296-word list. The number of possible combinations would be 

```
1,296 ^ 4 = 2,821,109,907,456
```

If an attacker could brute-force passwords on a website at 1,000 requests/second (where he would definitely hit rate limits), it would take about

```
2,821,109,907,456 / 1000 / 60 / 60 / 24 / 365 = 89.46
```

years to crack it. This is secure enough for online websites.  

## Offline

This mode generates a 6-word password. This equates to

```
1,296 ^ 6 = 4,738,381,338,321,616,896
```

possible combinations. With the MD5 hash and [8 Nvidia GTX 1080 GPUs](https://gist.github.com/epixoip/a83d38f412b4737e99bbef804a270c40) cracking the password would take about

```
4,738,381,338,321,616,896 / 25,000,000,000 / 60 / 60 / 24 / 365 = 6.010
```

years to crack. That would cost the attacker a lot of energy. If the password was hashed using SHA512 it would take about

```
4,738,381,338,321,616,896 / 1,100,000,000 / 60 / 60 / 24 / 365 = 136.6
```

years to crack. This is definitely secure enough for offline use (your laptop's password, encrypted drive)

## Crypto

This mode generates a 8-word password. There would be

```
1,296 ^ 8 = 7,958,661,109,946,400,884,391,936
```

possible combinations, and would take about

```
7,958,661,109,946,400,884,391,936 / 1,100,000,000 / 60 / 60 / 24 / 365 = 229400000
```

years to crack. If someone were to sponsor 8000 Nvidia GTX 1080 GPUs, it will still take a huge number of years to crack. 
