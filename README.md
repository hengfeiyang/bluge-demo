# bluge-demo
Demo for bluge

```
go run main.go
2022/12/26 18:35:14 indexed 0 documents
2022/12/26 18:35:28 indexed 1000 documents
2022/12/26 18:35:42 indexed 2000 documents
2022/12/26 18:35:57 indexed 3000 documents
2022/12/26 18:36:13 indexed 4000 documents
2022/12/26 18:36:28 indexed 5000 documents
2022/12/26 18:36:42 indexed 6000 documents
2022/12/26 18:36:56 indexed 7000 documents
2022/12/26 18:37:11 indexed 8000 documents
2022/12/26 18:37:26 indexed 9000 documents
2022/12/26 18:37:43 took 5 ms, found 10000 documents
2022/12/26 18:37:43 map[@timestamp:2022-12-26 10:35:14.029792 +0000 UTC _id:0 email:jone@mail.com gender:male name:John Doe phone:1234567890]
2022/12/26 18:37:43 map[@timestamp:2022-12-26 10:35:14.142486 +0000 UTC _id:1 email:jone@mail.com gender:male name:John Doe phone:1234567890]
2022/12/26 18:37:43 map[@timestamp:2022-12-26 10:35:14.154815 +0000 UTC _id:2 email:jone@mail.com gender:male name:John Doe phone:1234567890]
2022/12/26 18:37:43 map[@timestamp:2022-12-26 10:35:14.173795 +0000 UTC _id:3 email:jone@mail.com gender:male name:John Doe phone:1234567890]
2022/12/26 18:37:43 map[@timestamp:2022-12-26 10:35:14.195195 +0000 UTC _id:4 email:jone@mail.com gender:male name:John Doe phone:1234567890]
2022/12/26 18:37:43 map[@timestamp:2022-12-26 10:35:14.220206 +0000 UTC _id:5 email:jone@mail.com gender:male name:John Doe phone:1234567890]
2022/12/26 18:37:43 map[@timestamp:2022-12-26 10:35:14.242228 +0000 UTC _id:6 email:jone@mail.com gender:male name:John Doe phone:1234567890]
2022/12/26 18:37:43 map[@timestamp:2022-12-26 10:35:14.263687 +0000 UTC _id:7 email:jone@mail.com gender:male name:John Doe phone:1234567890]
2022/12/26 18:37:43 map[@timestamp:2022-12-26 10:35:14.276628 +0000 UTC _id:8 email:jone@mail.com gender:male name:John Doe phone:1234567890]
2022/12/26 18:37:43 map[@timestamp:2022-12-26 10:35:14.297706 +0000 UTC _id:9 email:jone@mail.com gender:male name:John Doe phone:1234567890]
```

```
ll data/index_dir
total 9536
-rw-------  1 yanghengfei  staff   292K Dec 26 18:35 000000000487.seg
-rw-------  1 yanghengfei  staff    15K Dec 26 18:35 000000000a3a.seg
-rw-------  1 yanghengfei  staff   646K Dec 26 18:35 000000000a3c.seg
-rw-------  1 yanghengfei  staff    46K Dec 26 18:36 00000000131b.seg
-rw-------  1 yanghengfei  staff    31K Dec 26 18:36 00000000155a.seg
-rw-------  1 yanghengfei  staff   2.1M Dec 26 18:36 000000001fdd.seg
-rw-------  1 yanghengfei  staff   1.2M Dec 26 18:37 0000000030f6.seg
-rw-------  1 yanghengfei  staff   330K Dec 26 18:37 000000003562.seg
-rw-------  1 yanghengfei  staff   1.7K Dec 26 18:37 000000003563.seg
-rw-------  1 yanghengfei  staff   1.8K Dec 26 18:37 000000003564.seg
-rw-------  1 yanghengfei  staff   1.7K Dec 26 18:37 000000003565.seg
-rw-------  1 yanghengfei  staff   1.7K Dec 26 18:37 000000003566.seg
-rw-------  1 yanghengfei  staff   307B Dec 26 18:37 000000005c75.snp
```

Tips

> 1. if you not sleep for merge, quickly quit after write data, it maybe many seg files. because it has no time to do merge.
> 2. it will has hundreds seg files when it writing.
> 3. if you want reduce files during writing, just use batch. like: 1000 documents one batch, it will just create 10 files for 10,000 documents.
