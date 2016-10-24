Aliyunddns, Update Aliyun DDNS record from OpenWRT or CLI
======================================================

[Homepage](https://github.com/gutenye/aliyun-ddns) |
[Documentation](https://github.com/gutenye/aliyun-ddns/wiki) |
[Issue Tracker](https://github.com/gutenye/aliyun-ddns/issues) |
[MIT License](http://choosealicense.com/licenses/mit) |
[by Guten](http://guten.me) |
[Gratipay](https://gratipay.com/gutenye) |
[Bountysource](https://www.bountysource.com/teams/gutenye)

|                |                                                            |
|----------------|------------------------------------------------------------|
|                | **Install**                                   |
| Binaries       | https://github.com/gutenye/aliyun-ddns/releases |
| Go             | `go get github.com/gutenye/aliyun-ddns`         |

Getting started
---------------

**OpenWRT Usage**

Start a http server in your personal computer, then use builtin OpenWRT dddns service.

```
$ aliyun-ddns server 3000
$ edit dddns custom url in OpenWRT: <server-ip>:3000/?id=12345679&rr=www&value=[IP]&access_key_id=[USERNAME]&access_key_secret=[PASSWORD]
```

**Cmdline Usage**

```
$ edit ~/.aliyun-ddnsrc

  ACCESS_KEY_ID = "x"
  ACCESS_KEY_SECRET = "x"

$ aliyun-ddns list example.com

12345678 www 1.1.1.1 example.com A

$ aliyun-ddns update 12345678 www 2.2.2.2
```

Development
===========

Contributing
-------------

* Submit any bugs/features/ideas to github issue tracker.
* Thanks to [all contributors](https://github.com/gutenye/aliyun-ddns/contributors?type=a).

Resource
--------

* [Mickyxing/aliyun-ddns](https://github.com/Mickyxing/aliyun-ddns), similar project, but written in Python.

Copyright
---------

The MIT License

Copyright (c) 2015 Guten Ye

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
