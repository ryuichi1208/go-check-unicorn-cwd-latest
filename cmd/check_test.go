package check

import "testing"

func Test_getProcessNameToPID(t *testing.T) {
	type args struct {
		processName string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getProcessNameToPID(tt.args.processName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getProcessNameToPID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getProcessNameToPID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_symLinkCheckExists(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := symLinkCheckExists(tt.args.link); (err != nil) != tt.wantErr {
				t.Errorf("symLinkCheckExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_symLinkCheckLatest(t *testing.T) {
	type args struct {
		link string
		dir  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := symLinkCheckLatest(tt.args.link, tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("symLinkCheckLatest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_checkProcessCWD(t *testing.T) {
	type args struct {
		pid int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkProcessCWD(tt.args.pid); (err != nil) != tt.wantErr {
				t.Errorf("checkProcessCWD() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_parseArgs(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseArgs(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("parseArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Do()
		})
	}
}
