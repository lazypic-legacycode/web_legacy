# LAZYPIC 웹서비스
* [lazypic.org](http://lazypic.org)

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

#### 개발자 규칙 : 같이 개발하기
퍼블리쉬 리포지터리와 개발 리포지터리를 구분하기 위해서 일반적으로 upstream 알리아스를 설정합니다.

upstream 을 설정하는 방법입니다.

- github.com/lazypic/web 리포지터리를 자신의 계정으로 Fork 합니다.
- Fork 한 리포지터리를 `clone` 합니다.
- `clone`된 폴더로 이동후 터미널에서 아래처럼 타이핑 합니다.
```
$ git remote add upstream https://github.com/lazypic/web
```

개발중인 코드를 업로드 할 때는 `git push origin master`를 사용하고 Pull Request 를 합니다.

퍼블리쉬된 코드를 pull 할 때는 `git pull upstream master`를 사용합니다.

#### 디자이너가 이미지 업로드하는 방법
- 이 방식은 디자이너에게 조금 불편하지만, 서버가 바뀌거나 백업시 가장 안전합니다.
- 최초 이미지를 업로드한다면, 홈 디렉토리에 현 코드저장소를 다운로드합니다.
```
cd ~
git clone http://github.com/lazypic/wb
```

- 이미지 올리는 방법을 설명합니다.
- 커피캣 32화를 예시로 작성합니다.
- 0032.png 파일을 ~/web/toon/coffeecat 경로에 저장합니다.
```
$ cd ~/web/toon/coffeecat
$ git pull origin master // 개발자가 업로드한 내용이 있을 수 있다 업데이트한다.
$ git add 0032.png
$ git commit -m "커피캣 32화 업로드"
$ git push origin master
```

- 정상적으로 업로드가 되었으면 서버에 접속합니다.
```
$ ssh root@lazypic.org
# cd ~/go/src/github.com/lazypic/web
# git pull
```

#### 디자인 & 정보
- 하단의 아이콘의 가로 사이즈는 40x40 이다.
- 이미지는 유지보수의 편리함을 위해서 SVG를 사용한다.
- 사용했던 Bitcoin 거래소 : https://bitwhere.com
- 주제색을 5%만 사용할 것. : http://post.naver.com/viewer/postView.nhn?volumeNo=3899006&memberNo=956644
