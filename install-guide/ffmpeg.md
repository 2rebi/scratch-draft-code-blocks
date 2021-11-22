# 다운로드받을 디렉토리로 이동
- 원하는 곳에 다운로드 받으면 됩니다.
```shell
# cd /usr/local/bin/
# mkdir ffmpeg && cd ffmpeg
```

# 압축 파일 다운로드 및 압축 해제

- 본 예제는 최신 버전 설치를 다룸
- [버전 리스트](https://www.johnvansickle.com/ffmpeg/old-releases/)

## 예시
```shell
# wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
# tar -xf ffmpeg-release-amd64-static.tar.xz
# ls
ffmpeg-4.4.1-amd64-static  ffmpeg-release-amd64-static.tar.xz
# cd ffmpeg-4.4.1-amd64-static
# cp -a . ..
```

# 링크 걸기
```shell
# ln -s /usr/local/bin/ffmpeg/ffmpeg /usr/bin/ffmpeg
# ln -s /usr/local/bin/ffmpeg/ffprobe /usr/bin/ffprobe
# ffmpeg -version
ffmpeg version 4.4.1-static https://johnvansickle.com/ffmpeg/  Copyright (c) 2000-2021 the FFmpeg developers
built with gcc 8 (Debian 8.3.0-6)
...
```