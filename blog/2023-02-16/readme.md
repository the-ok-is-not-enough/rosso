# TLS

- https://github.com/refraction-networking/utls/blob/master/examples
- https://godocs.io/crypto/tls#example-Dial
- https://stackoverflow.com/questions/73596694/how-to-wrap-net-conn-read

With the code in this folder, we can now inject custom client hello. So all we
need now is code to Marshal a client hello. The standard code is here:

<https://github.com/golang/go/blob/2da8a555/src/crypto/tls/handshake_messages.go#L97-L301>

the modified code is here:

<https://github.com/refraction-networking/utls/blob/d139a4a6/handshake_messages.go#L101-L305>

they are exactly the same. OK I have an Unmarshal function now, so how do I get
some bytes to test it out?
