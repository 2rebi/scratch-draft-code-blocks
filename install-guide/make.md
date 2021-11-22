# 다운로드받을 디렉토리로 이동
- 원하는 곳에 다운로드 받으면 됩니다.
```shell
# cd /usr/local/bin/
# mkdir make && cd make
```

# 압축 파일 다운로드 및 압축 해제
## 포맷
```shell
# wget http://ftp.gnu.org/gnu/make/make-{X.Y}.tar.gz
# tar xvf make-{X.Y}.tar.gz && cd make-{X.Y}
```

- `{X.Y}`은 다운받고자 하는 버전으로
  - 본 예제에선 4.1을 받고자 함.
- [버전 리스트](http://ftp.gnu.org/gnu/make)

## 예시
```shell
# wget http://ftp.gnu.org/gnu/make/make-4.1.tar.gz
# tar xvf make-4.1.tar.gz && cd make-4.1
```

# 컴파일
```shell
# ./configure && make
```

# 덮어쓰기
```shell
# whereis make
make: /usr/bin/make /usr/local/bin/make /usr/share/man/man1/make.1.gz /usr/share/man/man1p/make.1p.gz /usr/share/info/make.info-1.gz /usr/share/info/make.info-2.gz /usr/share/info/make.info.gz
#  mv ./make /usr/bin/make
mv: overwrite ‘/usr/bin/make’? y
```