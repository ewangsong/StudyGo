package main

import (
	"ewangsong/otp"
	"fmt"
)

func main() {
	hotp := &otp.HOTP{
		BasicOtp: otp.BasicOtp{
			Secret:  "AAAABBBBCCCCDDDD",
			Length:  6,
			Counter: 0,
		},
	}

	fmt.Println(hotp.Get())

	totp := &otp.TOTP{
		BasicOtp: otp.BasicOtp{
			Secret: "RB5C3SX7JNMOLKXB",
			Length: 6,
		},
	}
	fmt.Println(totp.Get())

	fmt.Println(otp.GenerateRandomSecret())
}
