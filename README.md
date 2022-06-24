# redis-qsg
- [redis-qsg](#redis-qsg)
- [소개](#소개)
- [컴파일 방법](#컴파일-방법)

# 소개
redis quick start guide go lang example code

# 컴파일 방법
- 개발환경
  - CentOS Linux release 7.9.2009
  - go1.18.3
  - go 환경변수
```
GOPATH="$HOME/proj"
GOROOT="$HOME/go"
```

- 초기화
```shell
go mod init ntels.com/redis-qsg
go mod tidy
```

- 빌드
```shell
go build
```

- 실행
```
$ ./redis-qsg -h
Usage of ./redis-qsg:
  -c    Clear redis 'Count'
  -i int
        Report Interval (default 1000)
  -m int
        number of redis INCR('Count') run (default 100)
  -p int
        number of go routines to run (default 1)
$
```
