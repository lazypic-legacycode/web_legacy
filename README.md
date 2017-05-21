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
퍼블리쉬 리포지터리와 개발 리포지터리를 구분하기 위해서 일반적으로 upstream 알리아스를 설정합니다.

upstream 을 설정하는 방법입니다.

- github.com/lazypic/web 리포지터리를 자신의 계정으로 fork 합니다.
- fork한 리포지터리를 `clone` 합니다.
- `clone`된 폴더로 이동후 터미널에서 아래처럼 타이핑 합니다.
```
$ git remote add upstream https://github.com/lazypic/web
```

개발중인 코드를 업로드 할 때는 `git push origin master`를 사용하고 리뷰를 합니다.

퍼블리쉬된 코드를 pull 할 때는 `git pull upstream master`를 사용합니다.

#### 디자인 & 정보
- 하단의 아이콘의 가로 사이즈는 40x40 이다.
- 이미지는 유지보수의 편리함을 위해서 SVG를 사용한다.
- 사용했던 Bitcoin 거래소 : https://bitwhere.com
- 주제색을 5%만 사용할 것. : http://post.naver.com/viewer/postView.nhn?volumeNo=3899006&memberNo=956644
