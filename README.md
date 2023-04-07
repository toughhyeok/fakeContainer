# fakeContainer
golang으로 Docker Container 인 척?
(only linux)

---

## 루트 파일시스템 다운로드
```bash
$ wget https://github.com/tianon/docker-brew-ubuntu-core/raw/88ba31584652db8b96a29849ea2809d99ce3cc31/focal/ubuntu-focal-oci-amd64-root.tar.gz
$ mkdir /tmp/ubuntu
$ tar zxf ubuntu-focal-oci-amd64-root.tar.gz -C /tmp/ubuntu
```

## RUN
```bash
[toughhyeok@origin-hostname fakeContainer]$ sudo su
[root@origin-hostname fakeContainer]$ go run . run /bin/bash
>>> Running: [/bin/bash] as 30941
>>> Running child: [/bin/bash] as 1
root@fake-container:/# ps
>>>    PID TTY          TIME CMD
>>>      1 ?        00:00:00 exe
>>>      5 ?        00:00:00 bash
>>>      8 ?        00:00:00 ps
root@fake-container:/# ls
>>> bin  boot  dev  etc  home  lib  lib32  lib64  libx32  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
```
