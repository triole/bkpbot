# Bkpbot

<!--- mdtoc: toc begin -->

1.	[How to use?](#how-to-use-)
2.	[Crontab](#crontab)<!--- mdtoc: toc end -->

My personal backup solution. Archives a list of files into a single file. Supported formats are `zip` and `tar`.

## How to use?

Usually we backup into `~/.backup`. Look into `howto` for a model of how backup scripts could be organized. Remember to put the bkpbot binary to `~/.backup/bin/bkpbot`.

Configuration file examples are in `testdata` or `howto`.

## Crontab

Backups may be launched by crontab. A few examples:

```crontab
5 5 * * * /home/ole/.backup/run/run.sh daily 9
3 3 * * 1 /home/ole/.backup/run/run.sh weekly 11
1 1 1 * * /home/ole/.backup/run/run.sh monthly 13

# only run on odd days
0 0 1-31/2 * * /home/ole/.backup/run.sh daily 9

# only run on even days
0 0 2-30/2 * * /home/ole/.backup/run.sh daily 9
```
