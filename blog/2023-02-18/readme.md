# February 18 2023

How do we know which slice to intercept?

~~~
record type
16 handshake
handshake type
01 client hello

record type
16 handshake
handshake type
10 Client Key Exchange
~~~

https://github.com/refraction-networking/utls/blob/v1.2.2/conn.go#L27

Its also the first one.
