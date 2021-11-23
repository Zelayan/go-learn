
# HTTP Reusing

读取完 http.Response 的 Body并关闭它，否则不会重用底层的TCP连接

通过Reusing()和NotResuing()方法分析关闭和不关闭Body，然后通过wireshark抓取tcp包，来验证该功能

resp.Body不关闭会导致每次都会建立新的连接，通过两个该例子，wireshark抓包可以看到，未关闭body需要经常去三次握手建立连接，而关闭body后，可以复用
# 参考

