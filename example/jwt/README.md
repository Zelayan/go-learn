# jwt go

JWT 由什么组成？

JWT 由三部分组成：

 - 标头：token 的类型和使用的签名算法。

 - token 的类型可以是“JWT”，而签名算法可以是 HMAC 或 SHA256。

 - 有效负载：token 中包含声明的第二部分。这些声明包括特定于应用程序的数据（例如：用户 ID、用户名）、token 到期时间 (exp)、颁发者 (iss)、主题 (sub) 等。
签名：编码的标头、编码的有效负载和您提供的密码用于创建签名。




#参考
 https://www.nexmo.com/legacy-blog/2020/09/07/using-jwt-for-authentication-in-a-golang-application-dr-2