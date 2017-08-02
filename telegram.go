package main

import (
	"fmt"
	//"os"
	"net/http"
	"log"
	"io/ioutil"
	//"crypto/tls"
	"strings"
	"encoding/json"
	tb "github.com/TelegramApi/bot"
	//"github.com/skycoin/skycoin-exchange/src/coin/skycoin"
	//"github.com/skycoin/skycoin/src/cipher"
)

//help - DONE
//define - DONE
//set - DONE
//get - DONE
//check  - in prog

func main(){
bot := tb.Create("425873758:AAFE7cUnYaAY1uCeIUdNL_PN26DufG1spSM")

//file, fileErr := os.Create("file")
//if fileErr != nil {
//	fmt.Println(fileErr)
//	return
//}
	
bot.Listen()

userAddrBind := make(map[string]string)

for update := range bot.Updates {

		var outputMessage string
		command := strings.Split(update.Message.Text, " ")

		switch command[0] {
		case "/help":
			outputMessage = "Available Commands: \n /help \n /define [command (without '/' prefix)] \n /set [skycoinAddress] \n /get \n /check \n NOTE: /get and /check will only work if you have already run /set successfully \n run /define to learn more about each command"
		case "/define":
			if len(command) < 2 {
				outputMessage = "Please specifiy a command."
			} else if len(command) > 2 { 
				outputMessage = "Too many inputs, please try again"
			} else {
				cmd := command[1]
				switch cmd {
					case "help":
						outputMessage = "/help: lists all available commands"
					case "define":
						outputMessage = "/define [command]: defines the function of all commands"
					case "set":
						outputMessage = "/set [skycoinAddress]: pairs a given skycoin address with the firstname of chat partner"
					case "get":
						outputMessage = "/get: checks that the current chat partner has already set a skycoin address, then generates a bitcoin deposit address to be paired with that skycoin address"
					case "check":	
						outputMessage = "/check: checks status of transactions involving chat partner's paired bitcoin deposit address"
					default:
						outputMessage = "Didn't recognize that command, please try again"
					}
			}
		case "/set":
			if len(command) < 2 {
				outputMessage = "Plese specifiy a skycoin address to pair with this telegram user account."
			} else if len(command) > 2 { 
				outputMessage = "Too many inputs, please try again"
			} else {
			skyAddr := command[1]
			user := update.Message.From.FirstName
			userAddrBind[string(user)] = skyAddr
			outputMessage = "Your name/skycoin address pair is: " + user + " : " + string(skyAddr) 
		}
		case "/get":
			user := update.Message.From.FirstName
			if userAddrBind[user] == "" {
				outputMessage = "First run /set with a skycoin wallet address to access this command"
			} else {
				user := update.Message.From.FirstName
				skyAddr := userAddrBind[user]
				fmt.Println(skyAddr)
				res, err := http.Get("http://localhost:7071/bind?skyaddr=" + skyAddr)
				if err != nil {
					outputMessage = "Something went wrong, could not get status"
					log.Println(err)
				}
				status, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					outputMessage = "Something went wrong, could not get status"
					log.Println(err)
				}

				statusString := string(status)
				//fmtString := strings.Split(statusString, "\n")

				//outputMessage = fmtString[5]
				//skyBtcAddrBind := make(map[string]cipher.Address)
				//skyBtcAddrBind[skyAddr] = pubkeyAddr
				//fmt.Fprintf(file, "%v\n", skyBtcAddrBind)
				outputMessage = "Your skycoin/bitcoin address pair is: " + string(skyAddr) + " : " + statusString 
			}
		case "/check":
			//user := update.Message.From.FirstName
			//if userAddrBind[user] == "" {
			//	outputMessage = "First run /set with a skycoin wallet address to access this command"
			//} else {
			//	skyAddr := userAddrBind[user]
				res, err := http.Get("http://localhost:7071/status?skyaddr=QDaJoA1Gcvt9oRVr1rfwUKCRYxM1A75KML")
				if err != nil {
					outputMessage = "Something went wrong, could not get status"
					log.Println(err)
				}
				status, err := ioutil.ReadAll(res.Body)
				res.Body.Close()
				if err != nil {
					outputMessage = "Something went wrong, could not get status"
					log.Println(err)
				}

			statusString := string(status)
			fmtString := strings.Split(statusString, "\n")

			outputMessage = fmtString[5]

			//}
		default:
			outputMessage = "No command issued, type /help for list of available commands"
		}

		var chat tb.User
		json.Unmarshal(update.Message.Chat, &chat)

		var keyboard = tb.ReplyKeyboardMarkup{Keyboard: [][]string{[]string{"/help"}}}
		bot.SendMessage(chat.Id, outputMessage, false, 0, keyboard)
	}

}






