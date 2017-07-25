package main

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/skycoin/skycoin-exchange/src/coin/skycoin"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/thoj/go-ircevent"
)

const channel = "#testing"
const serverssl = "81.7.17.69:6667"

func main() {
	ircnick1 := "bbot2"
	irccon := irc.IRC(ircnick1, "bbot3")
	irccon.VerboseCallbackHandler = true
	irccon.Debug = true
	irccon.UseTLS = false
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(channel) })
	irccon.AddCallback("366", func(e *irc.Event) {})

	userAddrBind := make(map[string]string)

	irccon.AddCallback("PRIVMSG", func(event *irc.Event) {

		command := strings.Split(event.Message(), " ")

		switch switchCmd := command[0]; switchCmd {
		case "HELP":
			irccon.Privmsgf(event.Nick, "Available Commands: HELP - SET - GET - SKYBALANCE - CHECK")
		case "SET":
			skyAddr := command[1]
			userAddrBind[event.Nick] = skyAddr
			irccon.Privmsgf(event.Nick, "Your username/skycoin address pair is: ", userAddrBind)
		case "GET":
			if userAddrBind[event.Nick] == "" {
				irccon.Privmsgf(event.Nick, "First run SET with a skycoin wallet address to access this command")
			} else {
				skyAddr := userAddrBind[event.Nick]
				fmt.Println(skyAddr)
				pubkey, _ := cipher.GenerateKeyPair()
				pubkeyAddr := cipher.AddressFromPubKey(pubkey)
				skyBtcAddrBind := make(map[string]cipher.Address)
				skyBtcAddrBind[skyAddr] = pubkeyAddr
				irccon.Privmsgf(event.Nick, "Your skycoin/bitcoin address pair is: ", skyBtcAddrBind)
			}
		case "SKYBALANCE":
			skyAddr := command[1]
			balance := skycoin.GetBalance(skyAddr)
			irccon.Privmsgf("The address of that skycoin address is: ", balance)
		case "CHECK":
			irccon.Privmsgf(event.Nick, "UNDER CONSTRUCTION")
		default:
			irccon.Privmsgf(event.Nick, "Did not recognize that command. Type 'HELP' for a list of possible commands")
		}

		fmt.Print("received message\n")
		fmt.Printf("nick=\"%s\", message=\"%s\"\n", event.Nick, event.Message())

	})
	err := irccon.Connect(serverssl)
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	irccon.Loop()
}
