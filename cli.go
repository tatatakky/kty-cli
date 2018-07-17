package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

/** kty サブコマンド用の実装 **/
type Kty struct{}

func (f *Kty) Help() string {
	return "app foo"
}

func (f *Kty) Run(args []string) int {


	info := receive_args(args)

	for _, v := range info {
		fmt.Println(v)
	}

	return 0

}

func (f *Kty) Synopsis() string {
	return "Print \"Foo!\""
}

func receive_args(arguments []string) []string {

	info := []string{}

	info = append(info, OptionJudge(Options{"-u", arguments}))
	info = append(info, OptionJudge(Options{"-m", arguments}))
	info = append(info, OptionJudge(Options{"-s", arguments}))

	return info

}

func remove(args []string, search string) []string {
	result := []string{}
	for _, v := range args {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

type Options struct {
	option1  string
	argments []string
}

func OptionJudge(opt Options) string {

	var info string
	for i := range opt.argments {
		if opt.option1 == opt.argments[i] {

			idx := if_equal(Ifequal{opt.option1, opt.argments})+1
			if opt.option1 == "-s" {
				return opt.argments[idx]
			} else {
				for _, v := range opt.argments[idx:len(opt.argments)] {
						if v[0] == '-' {
							break
						}
						info += v + " "
				}
			}
			return info[:len(info)-1]
		}
	}
	return ""
}

type Ifequal struct {
	option string
	args   []string
}

func if_equal(ie Ifequal) int {
	for i, v := range ie.args {
		if v == ie.option {
			return i
		}
	}
	return -1
}

func main() {
	// コマンドの名前とバージョンを指定
	c := cli.NewCLI("app", "1.0.0")

	// サブコマンドの引数を指定
	c.Args = os.Args[1:]

	/*receive_args(os.Args)*/

	// サブコマンド文字列 と コマンド実装の対応付け
	c.Commands = map[string]cli.CommandFactory{
		"send": func() (cli.Command, error) {
			return &Kty{}, nil
		},
	}

	// コマンド実行
	exitStatus, err := c.Run()

	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)

}
