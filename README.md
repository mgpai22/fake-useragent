# fake-useragent
Up-to-date simple useragent faker with real world database in Golang

## Features

- Data is pre-downloaded & post-processed from [Intoli LLC](https://github.com/intoli/user-agents/tree/main/src) and the data is part of the package itself
- The most up-to-date database if user-agents
- The data consists of a wide range of browser agents and various browsers
- Retrieves user-agent strings (both of type: `desktop`, `tablet` and/or `mobile` UAs)
- Retrieve user-agent fron JSON file in Go struct, with fields like `Useragent`, `Percent`, `Type`, `DeviceBrand`, `Browser`, `BrowserVersion`, `Os`, `OsVersion` and `Platform`
- This Golang package has the same functionality as popular Python libray "Fake-useragent" and based on same Json user-agents database

 ```sh
go get github.com/lib4u/fake-useragent
```
