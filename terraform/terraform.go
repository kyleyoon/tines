package terraform

import (
	"bytes"
	"io"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// as util
func commandExists(cmd string) bool {
	path, err := exec.LookPath(cmd)
	if err != nil {
		log.Fatalf("terraform check: %s", err)
	}

	log.Println(path)
	return err == nil
}

// SetRemoteState is
func SetRemoteState(d map[string]gjson.Result) {
	log.Debugf("Access Key", d["access_key"])
	log.Debugf("Secret Key", d["secret_key"])
	log.Debugf("Security Token", d["security_token"])

	commandExists("terraform")

	cmd := exec.Command(
		"terraform", "init",
		"--reconfigure",
		`-backend-config=access_key=`+d["access_key"].String(),
		`-backend-config=secret_key=`+d["secret_key"].String(),
		`-backend-config=token=`+d["security_token"].String(),
	)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	// outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	// fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr)
}
