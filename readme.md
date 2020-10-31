# Bkpbot

<!--- mdtoc: toc begin -->

1.	[Synopsis](#synopsis)
2.	[Configuration files](#configuration-files)
3.	[How to use?](#how-to-use-)
4.	[Build and Test](#build-and-test)
5.	[Disclaimer](#disclaimer)<!--- mdtoc: toc end -->

## Synopsis

My personal backup solution. Archives a list of files into a single file. Supported formats are `zip` and `tar`. Can be built as static binary not requiring any dependencies to make sure backups always run no matter what packets are installed.

## Configuration files

Configuration files are in toml format. They basically consist of two blocks: variables and. Have a look into `testdata` to find out what they can do.

## How to use?

I usually run backups on servers using crontab. Here are a few examples how entries look like for different backups.

```crontab
# daily, weekly, monthly saved into different folders
5 5 * * * bkpbot conf.toml -l bkpbot.log -k 9 -s daily
3 3 * * 1 bkpbot conf.toml -l bkpbot.log -k 11 -s weekly
1 1 1 * * bkpbot conf.toml -l bkpbot.log -k 13 -s monthly

# only on odd or even days
0 0 1-31/2 * * bkpbot conf.toml -l bkpbot.log -k 9 -s odd_days
0 0 2-30/2 * * bkpbot conf.toml -l bkpbot.log -k 9 -s even_days
```

## Build and Test

I use [task](https://github.com/go-task/task) as a runner for builds and tests.

```shell

task

# just test
task test

# remember what you can do if you have task
task -l
```

## Disclaimer

Warning. Use this software at your own risk. I may not be hold responsible for data loss, starving your kittens or losing the bling bling powerpoint presentations you made to impress human resources with the employment's performance.
