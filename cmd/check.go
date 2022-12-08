package check

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jessevdk/go-flags"
)

type options struct {
	PROC_NAME string `short:"p" description:"process name" required:"true"`
}

var opts options

func getProcessNameToPID(processName string) (int, error) {
	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		return 0, fmt.Errorf("can't open procfs: %s", err)
	}

	for _, dir := range files {
		pid, err := strconv.Atoi(dir.Name())
		if err != nil {
			continue
		}
		cmdline, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
		if err != nil {
			log.Fatal(err)
		}
		if strings.Index(string(cmdline), processName) > -1 {
			return pid, nil
		}
	}

	return 0, fmt.Errorf("no such process: %s", processName)

}

func symLinkCheck(link string) error {
	fmt.Println(link)
	if _, err := os.Stat(link); os.IsNotExist(err) {
		return fmt.Errorf("no such directory: %s", link)
	}

	return nil
}

func checkProcessCWD(pid int) error {
	fmt.Println(pid)
	link, err := os.Readlink(fmt.Sprintf("/proc/%d/cwd", pid))
	if err != nil {
		return fmt.Errorf("error")
	}

	return symLinkCheck(link)
}

func parseArgs(args []string) error {
	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		return err
	}

	return nil
}

func Do() {
	err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pid, err := getProcessNameToPID(opts.PROC_NAME)
	if err != nil {
		fmt.Println(err)
	}

	err = checkProcessCWD(pid)
	if err != nil {
		fmt.Println(err)
	}
}
