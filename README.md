# CLI dev

#### Function

* [X]  time
* [X]  word
* [X]  translate

#### Install

~~~bash
git init 
git clone https://github.com/lance547/long.git
cd long

go build
go install 

~~~

#### Command:

##### translate

If you want to use the translate command ,you must to get the baidu translate's appid and the secret at first ,and replace the value in the /internal/translate file.

~~~bash
long translate -w hello -t zh
~~~
