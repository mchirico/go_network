package ssh

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type SSH struct {
	Host        string
	User        string
	Password    string
	Server      string
	UseSSHkey   bool
	SSHPubKey   string
	CMD         string
	File        string
	client      *ssh.Client
	config      *ssh.ClientConfig
	Repeats     int
	StatusAbort bool
}

func (s *SSH) GetHostKey() (ssh.PublicKey, error) {
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], s.Host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, errors.New(fmt.Sprintf("error parsing %q: %v", fields[2], err))
			}
			break
		}
	}

	if hostKey == nil {
		return nil, errors.New(fmt.Sprintf("no hostkey for %s", s.Host))
	}
	return hostKey, nil
}

func GetConfigForKey(user string, keyfile string) *ssh.ClientConfig {

	// If you want to validate key ...
	// hostKey, err := GetHostKey("smtp.aipiggybot.io")
	//if err != nil {
	//	log.Fatal(err)
	//}

	var hostKey ssh.PublicKey

	key, err := ioutil.ReadFile(filepath.Join(os.Getenv("HOME"), ".ssh", keyfile))
	if err != nil {
		log.Fatalf("Unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("Unable to parse private key: %v", err)

	}
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// Ignore key validation
	config.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	return config
}

func GetConfigForPassword(user string, password string) *ssh.ClientConfig {
	var hostKey ssh.PublicKey

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// Ignore key validation
	config.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	return config
}

func (s *SSH) Config() {

	var config *ssh.ClientConfig
	if s.UseSSHkey {
		config = GetConfigForKey(s.User, s.SSHPubKey)
	} else {
		config = GetConfigForPassword(s.User, s.Password)
	}

	client, err := ssh.Dial("tcp", s.Server, config)
	if err != nil {
		log.Println("Failed to dial: ", err)
		s.StatusAbort = true
		return
	}
	s.StatusAbort = false
	s.client = client
	s.config = config

}

func (s *SSH) Exec(results chan string) {

	if s.StatusAbort == true {
		results <- "Abort"
		return
	}

	client, err := ssh.Dial("tcp", s.Server, s.config)
	if err != nil {
		log.Println("Failed to dial2: ", err)
		results <- "Abort Dial2"
		return
	}

	defer client.Close()

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Println("Failed to create session: ", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(s.CMD); err != nil {
		log.Printf("Failed to run: %v" + err.Error())
	}

	results <- b.String()
}

// Simple Example

func (s *SSH) CmdServers() {

	s.Config()

	for i := 0; i < s.Repeats; i++ {
		results := make(chan string, 0)
		go s.Exec(results)
		data := <-results
		if len(data) > 10 {
			fmt.Println(s.Server)
			Append(s.File, data)
		}
		close(results)
	}
}

func Append(file string, data string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("Can't open file: %v\n", err)
	}

	defer f.Close()

	if _, err = f.WriteString(data); err != nil {
		log.Printf("can't write: %v\n", err)
	}
}
