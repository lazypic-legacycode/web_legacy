# lazypicutres web code.
* [lazypictures web service.](http://lazyd.org)

#### install
```
go get -u github.com/lazypic/web
```
#### run
```
cd $GOPATH/src/github.com/lazypic/web
web -http :80
```

#### 서버에서 실행할때.
- 서버에서는 리버스 프록시를 사용하기 때문에 아래처럼 서비스를 실행한다.
```
web -http :8080
```

#### 같이 개발하기
- github.com/lazypic/web을 `origin` 대신 `upstream` 으로 설정합니다.
- 설정방법은 아래와 같습니다.
```
$ git remote add upstream https://github.com/lazypic/web
```

#### 디자인 & 정보
- 하단의 아이콘의 가로 사이즈는 40x40 이다.
- 이미지는 유지보수의 편리함을 위해서 SVG를 사용한다.
- 사용했던 Bitcoin 거래소 : https://bitwhere.com
- 주제색을 5%만 사용할 것. : http://post.naver.com/viewer/postView.nhn?volumeNo=3899006&memberNo=956644
