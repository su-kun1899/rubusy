# rubusy

rubusy is cron job schedule searcher which tell you when it scheduled.

## Usage

```
$ rubusy crontab.txt
from: 2017/10/17 18:35 - to: 2018/10/17 18:35
==============================================
2017/10/18 09:15 : 15 9 * * * /tmp/hoge1.sh
2017/10/17 20:10 : 10 20 * * * /tmp/fuga1.sh
2017/12/01 10:10 : 10 10 * 12 * /tmp/fuga2.sh
2017/10/17 18:35 : * * * * * /tmp/fuga3.sh
2018/10/03 10:00 : * 10 3 10 * /tmp/fuga4.sh
2017/10/24 06:15 : 15 6 24 * * /tmp/piyo.sh
```
