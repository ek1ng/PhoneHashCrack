# PhoneHashCrack

用于AirDrop取证，提取电话号码Hash后爆破脚本

## 如何构建

```
go mod tidy

go build -o main
```


## 如何使用

```
./main -h                       
 ____  _                      _   _           _      ____                _    
|  _ \| |__   ___  _ __   ___| | | | __ _ ___| |__  / ___|_ __ __ _  ___| | __
| |_) | '_ \ / _ \| '_ \ / _ \ |_| |/ _` / __| '_ \| |   | '__/ _` |/ __| |/ /
|  __/| | | | (_) | | | |  __/  _  | (_| \__ \ | | | |___| | | (_| | (__|   < 
|_|   |_| |_|\___/|_| |_|\___|_| |_|\__,_|___/_| |_|\____|_|  \__,_|\___|_|\_\
                                                                              

Usage of /tmp/go-build3070724103/b001/exe/main:
  -end string
        input end phone hash
  -start string
        input start phone hash

```

example : `./main --start 11cbd --end 893d4`

![image.png](https://ek1ng-typora.oss-cn-hangzhou.aliyuncs.com/img/20230802180347.png)
