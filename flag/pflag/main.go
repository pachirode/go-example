package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
)

type host struct {
	value string
}

func (h *host) String() string {
	return h.value
}

func (h *host) Set(v string) error {
	h.value = v
	return nil
}

func (h *host) Type() string {
	return "host"
}

func base() {
	var ip *int = pflag.Int("ip", 1234, "help message for ip")

	var port int
	pflag.IntVar(&port, "port", 8080, "help message for port")

	var h host
	pflag.Var(&h, "host", "help message for host")

	// 解析命令行参数
	pflag.Parse()

	fmt.Printf("ip: %d\n", *ip)
	fmt.Printf("port: %d\n", port)
	fmt.Printf("host: %+v\n", h)
}

func normalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "old-flag-name":
		name = "new-flag-name"
		break
	}

	return pflag.NormalizedName(name)
}

func aliasFlag() {
	flagSet := pflag.NewFlagSet("test", pflag.ExitOnError)
	var ip = flagSet.IntP("new-flag-name", "i", 1234, "help message for new-flag-name")

	flagSet.SetNormalizeFunc(normalizeFunc)
	flagSet.Parse(os.Args[1:])

	fmt.Printf("ip: %d\n", *ip)
}

func main() {
	//base()
	aliasFlag()
}
