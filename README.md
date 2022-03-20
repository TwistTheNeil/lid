# lid

Associate files with live disks, storage, and backup devices

### How to use

```console
n@xps13:~$ ./lid init db test.db --db test.db
n@xps13:~$ ./lid device add --name 8tbtest --uuid 73eda09a-2348-48d2-b8d4-b17972d6e98d --db test.db
n@xps13:~$ ./lid analyze /mnt/videos/ --db test.db -r 4 -d -1 --save --device 8tbtest
n@xps13:~$ ./lid --db test.db web
```
