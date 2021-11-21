# HTTP Reusing

读取完 http.Response 的 Body并关闭它，否则不会重用底层的TCP连接

通过Reusing()和NotResuing()方法分析关闭和不关闭Body，然后通过wireshark抓取tcp包，来验证该功能

