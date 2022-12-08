package check

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
		// pid以外のprocファイルは不要
		if err != nil {
			continue
		}
		cmdline, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
		if err != nil {
			log.Fatal(err)
		}
		if strings.Index(string(cmdline), processName) > -1 {
			// 自身がヒットするので除外する
			if strings.Index(string(cmdline), "check-unicorn-cwd-latest") > -1 {
				continue
			}
			return pid, nil
		}
	}

	return 0, fmt.Errorf("no such process: %s", processName)

}

// シンボリックリンク先のディレクトリが存在するかを確認する
func symLinkCheckExists(link string) error {
	if _, err := os.Stat(strings.Split(link, " ")[0]); os.IsNotExist(err) {
		return fmt.Errorf("no such directory: %s", link)
	}

	return nil
}

func symLinkCheckLatest(link, dir string) error {
	files, _ := ioutil.ReadDir(dir)
	var newestFile string
	var newestTime int64 = 0
	for _, f := range files {
		fi, err := os.Stat(filepath.Join(dir, "/", f.Name()))
		if err != nil {
			continue
		}
		currTime := fi.ModTime().Unix()

		// 最新ディレクトリ名を取得する
		if currTime > newestTime {
			newestTime = currTime
			newestFile = f.Name()
		}
	}

	// 現在のcwdが最新のディレクトリでない場合はエラーを返す
	if strings.Split(link, " ")[0] != filepath.Join(dir, "/", newestFile) {
		return fmt.Errorf("Current reference is not up-to-date")
	}

	return nil
}

func checkProcessCWD(pid int) error {
	link, err := os.Readlink(fmt.Sprintf("/proc/%d/cwd", pid))
	if err != nil {
		return fmt.Errorf("error")
	}

	err = symLinkCheckExists(link)
	if err != nil {
		return err
	}

	err = symLinkCheckLatest(link, opts.RELEASE_DIR)
	if err != nil {
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
	err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 指定されたディレクトリが存在しない場合はエラー終了
	if _, err := os.Stat(opts.RELEASE_DIR); os.IsNotExist(err) {
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
