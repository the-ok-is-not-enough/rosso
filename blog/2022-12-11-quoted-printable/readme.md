# December 11 2022

- https://wikipedia.org/wiki/Percent-encoding
- https://wikipedia.org/wiki/Quoted-printable
<https://wikipedia.org/wiki/Binary-to-text_encoding>

The closest of these are `mime/quotedprintable` and `net/url`. Neither handles
two byte code points. I think Quoted-Printable is the closer match. For single
byte code points, it only encodes the escape character:

~~~
=3D
~~~

while percent-encoding does many:

~~~
%20 %21 %22 %23 %25 %27 %28 %29 %2A %2C %2F %3B %3C %3E %3F %5B %5C %5D %5E %60
%7B %7C %7D
~~~

Quoted-Printable does insert extra newlines, but since we are encoding one rune
at a time, that is not an issue.

## mime/quotedprintable

~~~go
package main

import (
   "io"
   "mime/quotedprintable"
   "os"
)

func main() {
   w := quotedprintable.NewWriter(os.Stdout)
   w.Binary = true
   for r := '\x00'; r <= '~'; r++ {
      if r >= 1 {
         os.Stdout.WriteString(" ")
      }
      io.WriteString(w, string(r))
      w.Close()
   }
}
~~~

out:

~~~
=00 =01 =02 =03 =04 =05 =06 =07 =08 =09 =0A =0B =0C =0D =0E =0F =10 =11 =12 =13
=14 =15 =16 =17 =18 =19 =1A =1B =1C =1D =1E =1F =20 ! " # $ % & ' ( ) * + , - .
/ 0 1 2 3 4 5 6 7 8 9 : ; < =3D > ? @ A B C D E F G H I J K L M N O P Q R S T U
V W X Y Z [ \ ] ^ _ ` a b c d e f g h i j k l m n o p q r s t u v w x y z { | }
~
~~~

https://godocs.io/mime/quotedprintable

## net/url

~~~go
package main

import (
   "fmt"
   "net/url"
)

func main() {
   fmt.Println("PathEscape")
   for r := '\x00'; r <= '~'; r++ {
      fmt.Print(url.PathEscape(string(r)))
      if r == '~' {
         fmt.Println()
      } else {
         fmt.Print(" ")
      }
   }
   fmt.Println("QueryEscape")
   for r := '\x00'; r <= '~'; r++ {
      fmt.Print(url.QueryEscape(string(r)))
      if r == '~' {
         fmt.Println()
      } else {
         fmt.Print(" ")
      }
   }
}
~~~

out:

~~~
PathEscape
%00 %01 %02 %03 %04 %05 %06 %07 %08 %09 %0A %0B %0C %0D %0E %0F %10 %11 %12 %13
%14 %15 %16 %17 %18 %19 %1A %1B %1C %1D %1E %1F %20 %21 %22 %23 $ %25 & %27 %28
%29 %2A + %2C - . %2F 0 1 2 3 4 5 6 7 8 9 : %3B %3C = %3E %3F @ A B C D E F G H
I J K L M N O P Q R S T U V W X Y Z %5B %5C %5D %5E _ %60 a b c d e f g h i j k
l m n o p q r s t u v w x y z %7B %7C %7D ~

QueryEscape
%00 %01 %02 %03 %04 %05 %06 %07 %08 %09 %0A %0B %0C %0D %0E %0F %10 %11 %12 %13
%14 %15 %16 %17 %18 %19 %1A %1B %1C %1D %1E %1F + %21 %22 %23 %24 %25 %26 %27
%28 %29 %2A %2B %2C - . %2F 0 1 2 3 4 5 6 7 8 9 %3A %3B %3C %3D %3E %3F %40 A B
C D E F G H I J K L M N O P Q R S T U V W X Y Z %5B %5C %5D %5E _ %60 a b c d e
f g h i j k l m n o p q r s t u v w x y z %7B %7C %7D ~
~~~

https://godocs.io/net/url

## mime

~~~go
package main

import (
   "fmt"
   "mime"
   "strings"
)

func main() {
   var b strings.Builder
   for r := ' '; r <= '~'; r++ {
      b.WriteRune(r)
   }
   b.WriteRune('😀')
   fmt.Println(mime.QEncoding.Encode("utf-8", b.String()))
}
~~~

out:

~~~
=?utf-8?q?_!"#$%&'()*+,-./0123456789:;<=3D>=3F@ABCDEFGHIJKLMNOPQRSTUVWXYZ?=
=?utf-8?q?[\]^=5F`abcdefghijklmnopqrstuvwxyz{|}~=F0=9F=98=80?=
~~~

https://godocs.io/mime

## encoding/ascii85

~~~go
package main

import (
   "encoding/ascii85"
   "io"
   "os"
)

func main() {
   w := ascii85.NewEncoder(os.Stdout)
   io.WriteString(w, "hello world") // BOu!rD]j7B
}
~~~

https://godocs.io/encoding/ascii85

## encoding/base32

~~~go
package main

import (
   "encoding/base32"
   "io"
   "os"
)

func main() {
   {
      w := base32.NewEncoder(base32.HexEncoding, os.Stdout)
      io.WriteString(w, "hello world") // D1IMOR3F41RMUSJC
   }
   os.Stdout.WriteString("\n")
   {
      w := base32.NewEncoder(base32.StdEncoding, os.Stdout)
      io.WriteString(w, "hello world") // NBSWY3DPEB3W64TM
   }
}
~~~

https://godocs.io/encoding/base32

## encoding/base64

~~~go
package main

import (
   "encoding/base64"
   "io"
   "os"
)

func main() {
   {
      w := base64.NewEncoder(base64.StdEncoding, os.Stdout)
      io.WriteString(w, "hello world") // aGVsbG8gd29y
   }
   os.Stdout.WriteString("\n")
   {
      w := base64.NewEncoder(base64.URLEncoding, os.Stdout)
      io.WriteString(w, "hello world") // aGVsbG8gd29y
   }
}
~~~

https://godocs.io/encoding/base64

## BaseXML

in:

~~~
hello wor
~~~

out:

~~~
V͆qXovćޠd?4?
~~~

https://github.com/kriswebdev/BaseXML/tree/master/BaseXML%20BS%20for%20XML1.0%20for%20C

## github.com/adiabat/bech32

~~~go
package main

import (
   "fmt"
   "github.com/adiabat/bech32"
)

func main() {
   data := []byte("hello world")
   s := bech32.Encode("", data)
   fmt.Println(s) // 1dpjkcmr0ypmk7unvvsunn0xe
}
~~~

https://godocs.io/github.com/adiabat/bech32

## BinHex

in:

~~~
hello world
~~~

out:

~~~
(This file must be converted with BinHex 4.0)
:#hGPBR9dD@acAh"X!$mr2cmr2cmr!!!!!!!,!!!!!+m5D'9XE'mJGfpbE'3lj!!
!:
~~~

https://webutils.pl/index.php?idx=binhex

## encoding/hex

~~~go
package main

import (
   "encoding/hex"
   "io"
   "os"
)

func main() {
   w := hex.NewEncoder(os.Stdout)
   io.WriteString(w, "hello world") // 68656c6c6f20776f726c64
}
~~~

https://godocs.io/encoding/hex

## github.com/mash/go-intelhex

~~~go
package main

import (
   "encoding/hex"
   "github.com/mash/go-intelhex"
   "os"
)

func main() {
   src := []byte("hello world")
   var record intelhex.Record
   record.ByteCount = int64(len(src))
   record.Data = hex.EncodeToString(src)
   os.Stdout.Write(record.Format(99)) // :0B00000068656c6c6f20776f726c6499
}
~~~

https://godocs.io/github.com/mash/go-intelhex

## RFC 1751 (S/KEY)

in:

~~~
EB33 F77E E73D 4053
~~~

out:

~~~
TIDE ITCH SLOW REIN RULE MOT
~~~

https://datatracker.ietf.org/doc/html/rfc1751

## github.com/kisom/srec

~~~go
package main

import (
   "fmt"
   "github.com/kisom/srec"
)

func main() {
   data := []byte("hello world")
   s := srec.Dump16(nil, data, 0)
   fmt.Println(s)
}
~~~

out:

~~~
S0030000FC
S10E000068656C6C6F20776F726C6495
S5030001FB
S9030000FC
~~~

https://godocs.io/github.com/kisom/srec

## Tektronix hex

~~~
srec_cat in.txt -binary -o out.txt -Tektronix
~~~

in:

~~~
hello world
~~~

out:

~~~
/00000C0C68656C6C6F20776F726C640AA6
~~~

- <https://wikipedia.org/wiki/Tektronix_hex_format>
- https://sourceforge.net/projects/srecord/files/srecord-win32

## github.com/tejainece/uu

~~~go
package main

import (
   "github.com/tejainece/uu"
   "os"
)

func main() {
   bytes := []byte("hello world")
   bytes = uu.EncodeLine(bytes)
   os.Stdout.Write(bytes) // +:&5L;&\@=V]R;&0
}
~~~

- https://godocs.io/github.com/tejainece/uu
- https://wikipedia.org/wiki/Uuencoding

## Xxencoding

in:

~~~
http://www.wikipedia.org
~~~

out:

~~~
OO5FoQ1cj9rRrRmtrOKhdQ4JYOK2iPr7b1Ec+
~~~

https://wikipedia.org/wiki/Xxencoding

## github.com/tilinna/z85

~~~go
package main

import (
   "github.com/tilinna/z85"
   "os"
)

func main() {
   src := []byte("hello world!")
   dst := make([]byte, z85.EncodedLen(len(src)))
   _, err := z85.Encode(dst, src)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(dst) // xK#0@zY<mxA+]nf
}
~~~

https://godocs.io/github.com/tilinna/z85

## fmt

~~~go
package main

import "fmt"

func main() {
   b := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
   fmt.Printf("%q\n", b) // "\x00\x01\x02\x03\x04\x05\x06\a\b\t"
}
~~~

https://godocs.io/fmt

## strconv

~~~go
package main

import "strconv"

func main() {
   b := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
   s := strconv.Quote(string(b))
   println(s) // "\x00\x01\x02\x03\x04\x05\x06\a\b\t"
}
~~~

https://godocs.io/strconv
