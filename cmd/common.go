package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/aws/aws-sdk-go/service/sts"
	log "github.com/sirupsen/logrus"

	"github.com/mulesoft-labs/aws-keycloak/provider"
)

func makeProvider() (*provider.Provider, error) {
	k, err := provider.NewKeycloakProvider(kr, kcprofile, kcConf)
	if err != nil {
		return nil, err
	}
	if region == "" {
		if v, e := kcConf["default_region"]; e {
			region = v
		} else {
			region = provider.DefaultRegion
		}
	}
	a := &provider.AwsProvider{
		Keyring:  kr,
		Region:   region,
		Duration: int64(duration),
	}
	return &provider.Provider{
		A:          a,
		K:          k,
		AlwaysAuth: alwaysAuth,
	}, nil
}

func getAwsStsCreds() (sts.Credentials, error) {
	p, err := makeProvider()
	if err != nil {
		return sts.Credentials{}, err
	}
	stscreds, _, err := p.Retrieve(awsrole)
	return stscreds, err
}

func listRoles() ([]string, error) {
	p, err := makeProvider()
	if err != nil {
		return []string{}, err
	}
	return p.List()
}

/**
 * Appends AWS env vars to existing env
 */
func runWithAwsEnv(includeFullEnv bool, name string, arg ...string) error {
	stscreds, err := getAwsStsCreds()
	if err != nil {
		return err
	}

	env := []string{
		fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", *stscreds.AccessKeyId),
		fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", *stscreds.SecretAccessKey),
		fmt.Sprintf("AWS_SESSION_TOKEN=%s", *stscreds.SessionToken),
		fmt.Sprintf("AWS_KEYCLOAK_PROFILE=%s", awsrole),
	}
	if region != "" {
		env = append(env, fmt.Sprintf("AWS_DEFAULT_REGION=%s", region))
		env = append(env, fmt.Sprintf("AWS_REGION=%s", region))
	}

	log.Debugf("Running command `%s %s` with AWS env vars set", name, strings.Join(arg, " "))
	if includeFullEnv {
		env = append(os.Environ(), env...)
	}
	return runWithEnv(name, env, arg...)
}

/**
 * This method will only return if there is an erorr running the subcommand.
 * Otherwise it will Exit with the appropriate exit code.
 */
func runWithEnv(name string, env []string, arg ...string) error {
	binary, err := exec.LookPath(name)
	if err != nil {
		return fmt.Errorf("Error finding `%s`. Is it installed and in your PATH? %s", name, err)
	}

	cmd := exec.Command(binary, arg...)
	cmd.Env = env

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// capture all signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan)

	go func() {
		// forward any signals onto the child process
		for {
			sig := <-sigChan
			log.Debugf("Sending signal to child process `%s`", sig.String())
			_ = cmd.Process.Signal(sig)
		}
	}()

	// This is subtly different from simply `cmd.Run()`, though I don't understand why.
	if err = cmd.Start(); err != nil {
		return err
	}

	if err = cmd.Wait(); err != nil {
		_ = cmd.Process.Signal(os.Kill)
		var exit *exec.ExitError
		if errors.As(err, &exit) {
			os.Exit(exit.ProcessState.ExitCode())
		}
	}
	return err //which must be nil
}
