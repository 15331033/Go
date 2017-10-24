// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"Go/agenda/entity"
	"github.com/spf13/cobra"
	"os"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "To register an account",
	Long:  `To regist your account,the following parameters should be avaliable,u/user is your username,p/password is your password,e/email is your email,P/Phone is your phone`,
	Run: func(cmd *cobra.Command, args []string) {
		username, err1 := cmd.Flags().GetString("user")
		password, err2 := cmd.Flags().GetString("password")
		email, err3 := cmd.Flags().GetString("email")
		phone, err4 := cmd.Flags().GetString("phone")
		tempErr := []error{err1, err2, err3, err4}
		isempty := outPutErr(tempErr)
		if !isempty {
			fmt.Println("Register Fail!")
			os.Exit(1)
		}
		ifNull(username, password, email, phone)
		ifExist(username)
		var tempUser entity.User
		tempUser.SetUser(username, password, email, phone)
		s := &entity.Storage
		s.CreatUser(tempUser)

		fmt.Println("Register Success!")
		fmt.Println("Your information is\nusername: " + username + "\npassword: " + password + "\nemail: " + email + "\ntelephone: " + phone)
	},
}

func getAll(u entity.User) bool { return true }

func switcher(u *entity.User) { u.SetPassword("laji") }

func ifNull(u string, p string, e string, t string) {
	i := 0
	err := ""
	if u == "" {
		err += "\nYou must input your name."
		i += 1
	}
	if p == "" {
		err += "\nYou must input your phone."
		i += 1
	}
	if e == "" {
		err += "\nYou must input your email."
		i += 1
	}
	if t == "" {
		err += "\nYou must input your phone."
		i += 1
	}
	if i != 0 {
		fmt.Println("Register Fail!" + err)
		os.Exit(2)
	}
}

func ifExist(u string) {
	s := &entity.Storage
	userList := s.QueryUser(getAll)
	for _, v := range userList {
		if (&v).GetUsername() == u {
			fmt.Println("Register Fail!")
			fmt.Println("Your name " + u + " is already exist.")
			os.Exit(3)
		}
	}
}

func outPutErr(errs []error) bool {
	for _, value := range errs {
		if value != nil {
			fmt.Println(value)
			return false
		}
	}
	return true
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("user", "u", "", "Help message for username")
	registerCmd.Flags().StringP("password", "p", "", "the password,it should not be empty")
	registerCmd.Flags().StringP("email", "e", "", "the email,,it should be empty")
	registerCmd.Flags().StringP("phone", "t", "", "your phone,it should not be empyt")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}