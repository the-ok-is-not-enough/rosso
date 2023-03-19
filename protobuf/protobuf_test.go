package protobuf

import (
   "fmt"
   "testing"
)

func Test_Add(t *testing.T) {
   checkin := Message{
      4: Message{ // checkin
         1: Message{ // build
            10: Varint(29), // sdkVersion
         },
      },
      14: Varint(3), // version
      18: Message{ // deviceConfiguration
         1: Varint(999), // touchScreen
         2: Varint(999),
         3: Varint(999),
         4: Varint(999),
         5: Varint(999),
         6: Varint(999),
         7: Varint(999),
         8: Varint(999),
         9: Slice[String]{
            "org.apache.http.legacy",
            "android.test.runner",
            "global-miui11-empty.jar",
         },
         11: String("nativePlatform"),
         15: Slice[String]{
            "GL_OES_compressed_ETC1_RGB8_texture",
            "GL_KHR_texture_compression_astc_ldr",
         },
      },
   }
   androids := []string{
      "android.hardware.bluetooth",
      "android.hardware.bluetooth_le",
      "android.hardware.camera",
      "android.hardware.camera.autofocus",
      "android.hardware.camera.front",
      "android.hardware.location",
      "android.hardware.location.gps",
      "android.hardware.location.network",
      "android.hardware.microphone",
      "android.hardware.opengles.aep",
      "android.hardware.screen.landscape",
      "android.hardware.screen.portrait",
      "android.hardware.sensor.accelerometer",
      "android.hardware.sensor.compass",
      "android.hardware.sensor.gyroscope",
      "android.hardware.telephony",
      "android.hardware.touchscreen",
      "android.hardware.touchscreen.multitouch",
      "android.hardware.usb.host",
      "android.hardware.wifi",
      "android.software.device_admin",
      "android.software.midi",
   }
   for _, android := range androids {
      checkin.Get(18).Add(26, Message{
         1: String(android),
      })
   }
   fmt.Println(checkin)
}

func Test_Literal(t *testing.T) {
   checkin := Message{
      4: Message{ // checkin
         1: Message{ // build
            10: Varint(29), // sdkVersion
         },
      },
      14: Varint(3), // version
      18: Message{ // deviceConfiguration
         1: Varint(999), // touchScreen
         2: Varint(999),
         3: Varint(999),
         4: Varint(999),
         5: Varint(999),
         6: Varint(999),
         7: Varint(999),
         8: Varint(999),
         9: Slice[String]{
            "org.apache.http.legacy",
            "android.test.runner",
            "global-miui11-empty.jar",
         },
         11: String("nativePlatform"),
         15: Slice[String]{
            "GL_OES_compressed_ETC1_RGB8_texture",
            "GL_KHR_texture_compression_astc_ldr",
         },
         26: Slice[Message]{
            {1: String("android.hardware.bluetooth")},
            {1: String("android.hardware.bluetooth_le")},
            {1: String("android.hardware.camera")},
            {1: String("android.hardware.camera.autofocus")},
            {1: String("android.hardware.camera.front")},
            {1: String("android.hardware.location")},
            {1: String("android.hardware.location.gps")},
            {1: String("android.hardware.location.network")},
            {1: String("android.hardware.microphone")},
            {1: String("android.hardware.opengles.aep")},
            {1: String("android.hardware.screen.landscape")},
            {1: String("android.hardware.screen.portrait")},
            {1: String("android.hardware.sensor.accelerometer")},
            {1: String("android.hardware.sensor.compass")},
            {1: String("android.hardware.sensor.gyroscope")},
            {1: String("android.hardware.telephony")},
            {1: String("android.hardware.touchscreen")},
            {1: String("android.hardware.touchscreen.multitouch")},
            {1: String("android.hardware.usb.host")},
            {1: String("android.hardware.wifi")},
            {1: String("android.software.device_admin")},
            {1: String("android.software.midi")},
         },
      },
   }
   fmt.Println(checkin)
}
