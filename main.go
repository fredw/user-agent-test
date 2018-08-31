package main

import (
	"fmt"
	"github.com/avct/uasurfer"
	"github.com/mssola/user_agent"
	"strings"
)

type Device struct {
	name      string
	userAgent string
}

func main() {

	devices := []Device{
		{
			name:      "Android 8",
			userAgent: "Mozilla/5.0 (Linux; Android 8.0; Android SDK built for x86 Build/PSR1.180720.012; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 hellofresh",
		},
		{
			name:      "Android 9",
			userAgent: "Mozilla/5.0 (Linux; Android 9; Android SDK built for x86 Build/PSR1.180720.012; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 hellofresh",
		},
		{
			name:      "Samsung Galaxy S8",
			userAgent: "Mozilla/5.0 (Linux; Android 7.0; SM-G892A Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/60.0.3112.107 Mobile Safari/537.36",
		},
		{
			name:      "Samsung Galaxy S6",
			userAgent: "Mozilla/5.0 (Linux; Android 6.0.1; SM-G920V Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.98 Mobile Safari/537.36",
		},
		{
			name:      "Apple Iphone X",
			userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
		},
		{
			name:      "Apple Iphone 7",
			userAgent: "Mozilla/5.0 (iPhone9,3; U; CPU iPhone OS 10_0_1 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/14A403 Safari/602.1",
		},
		{
			name:      "Microsoft Lumia 650",
			userAgent: "Mozilla/5.0 (Windows Phone 10.0; Android 6.0.1; Microsoft; RM-1152) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Mobile Safari/537.36 Edge/15.15254",
		},
		{
			name:      "Windows 10 - Edge",
			userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246",
		},
		{
			name:      "Mac OS X - Safari",
			userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9",
		},
		{
			name:      "Mac OS X - Chrome",
			userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36",
		},
		{
			name:      "Linux - Firefox",
			userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1",
		},
	}

	for _, device := range devices {
		fmt.Printf("\n\n%v\n", device.name)
		fmt.Printf("%v\n", strings.Repeat("=", 50))

		fmt.Printf("\nLibrary 'MSSOLA'\n")
		fmt.Printf("%v\n", strings.Repeat("-", 30))
		userAgentMssola(device.userAgent)

		fmt.Printf("\nLibrary 'AVCT'\n")
		fmt.Printf("%v\n", strings.Repeat("-", 30))
		userAgentAvct(device.userAgent)
	}
}

func userAgentMssola(userAgent string) {
	ua := user_agent.New(userAgent)

	fmt.Printf("OS Plataform: %v\n", ua.Platform())
	fmt.Printf("OS Version: %v\n", ua.OS())
	fmt.Printf("Mobile: %v\n", ua.Mobile())
}

func userAgentAvct(userAgent string) {
	parsed := uasurfer.Parse(userAgent)

	fmt.Printf("OS Name: %v\n", parsed.OS.Name)
	fmt.Printf("OS Plataform: %v\n", parsed.OS.Platform)
	fmt.Printf("OS Version: %v\n", parsed.OS.Version)
	fmt.Printf("Device type: %v\n", parsed.DeviceType.String())
}
