# 다운로드받을 디렉토리로 이동
- 원하는 곳에 다운로드 받으면 됩니다.
```shell
# cd ~
# mkdir glibc && cd glibc
```

# 압축 파일 다운로드 및 압축 해제
## 포맷
```shell
# wget https://ftp.gnu.org/gnu/glibc/glibc-{X.Y}.tar.xz
# tar -xf glibc-{X.Y}.tar.xz && cd glibc-{X.Y}
```

- `{X.Y}`은 다운받고자 하는 버전으로
  - 본 예제에선 2.34을 받고자 함.
- [버전 리스트](https://ftp.gnu.org/gnu/glibc)

## 예시
```shell
# wget https://ftp.gnu.org/gnu/glibc/glibc-2.34.tar.xz
# tar -xf glibc-2.34.tar.xz && cd glibc-2.34
```

# 컴파일
## 포맷
```shell
# mkdir build && cd build
# ../configure --prefix=/opt/glibc-{X.Y}
```
- `{X.Y}`은 다운 받은 버전

## 예시
```shell
# mkdir build && cd build
# ../configure --prefix=/opt/glibc-2.34
# make -j4
# make install
```