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
	PROC_NAME   string `short:"p" description:"process name" required:"true"`
	RELEASE_DIR string `short:"d" description:"release directory" required:"true"`
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
			if strings.Index(string(cmdline), "check-unicorn-cwd-latest") > -1 {
				continue
			}
			return pid, nil
		}
	}

	return 0, fmt.Errorf("no such process: %s", processName)

}

func symLinkCheckExists(link string) error {
	fmt.Println(link)
	if _, err := os.Stat(strings.Split(link, " ")[0]); os.IsNotExist(err) {
		fmt.Println(link)
		return fmt.Errorf("no such directory: %s", link)
	}

	return nil
}

func symLinkCheckLatest(link, dir string) error {
	files, _ := ioutil.ReadDir(dir)
	var newestFile string
	var newestTime int64 = 0
	for _, f := range files {
		fi, err := os.Stat(dir + "/" + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		currTime := fi.ModTime().Unix()
		if currTime > newestTime {
			newestTime = currTime
			newestFile = f.Name()
		}
	}
	fmt.Println(newestFile)

	if strings.Split(link, "")[0] != dir+"/"+newestFile {
		return fmt.Errorf("Current reference is not up-to-date")
	}

	return nil
}

func checkProcessCWD(pid int) error {
	fmt.Println(pid)
	link, err := os.Readlink(fmt.Sprintf("/proc/%d/cwd", pid))
	if err != nil {
		return fmt.Errorf("error")
	}

	err = symLinkCheckExists(link)
	if err != nil {
		fmt.Println("not exists")
		return err
	}

	err = symLinkCheckLatest(link, opts.RELEASE_DIR)
	if err != nil {
		fmt.Println("not latest")
		return err
	}

	return nil
}

func parseArgs(args []string) error {
	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		return err
	}

	return nil
}

func Do() {
	fmt.Println(os.Args[0])
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
