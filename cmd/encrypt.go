/*
Copyright © 2024 Sean Campbell <sean.campbell13@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/natac13/go-enigma-machine/pkg/enigma"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a message using the Enigma machine.",
	Long:  `Encrypt a message using the Enigma machine.`,
	// ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string)	([]string, cobra.ShellCompDirective) {
	// 	var comps []string
	// 	if len(args) == 0 {
	// 		comps = cobra.AppendActiveHelp(comps, "Please provide a message to encrypt.")
	// 	}
	// 	if len(args) == 1 {
	// 		comps = cobra.AppendActiveHelp(comps, "Hello, World!")
	// 	}
	//
	// 	return comps, cobra.ShellCompDirectiveNoFileComp
	// },
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cobra.CheckErr(fmt.Errorf("you must provide a message to encrypt"))
		}
		message := strings.Trim(args[0], " ")
		if message == "" {
			cobra.CheckErr(fmt.Errorf("you must provide a message to encrypt"))
		}

		plugboard := enigma.NewPlugboard()
		reflector, err := enigma.NewReflector([]rune(enigma.REFLECTOR_B_WIRING))
		cobra.CheckErr(err)

		rotor1, err := enigma.NewRotor(
			[]rune(enigma.ROTOR_III_WIRING),
			enigma.ROTOR_III_NOTCH,
		)
		cobra.CheckErr(err)
		rotor2, err := enigma.NewRotor(
			[]rune(enigma.ROTOR_II_WIRING),
			enigma.ROTOR_II_NOTCH,
		)
		cobra.CheckErr(err)
		rotor3, err := enigma.NewRotor(
			[]rune(enigma.ROTOR_I_WIRING),
			enigma.ROTOR_I_NOTCH,
		)
		cobra.CheckErr(err)
		rotors := []*enigma.Rotor{rotor1, rotor2, rotor3}
		em := enigma.NewEnigmaMachine(plugboard, rotors, reflector)
		encrypted, err := em.EncryptString(message)
		cobra.CheckErr(err)

		fmt.Printf("Original message: %s\n", message)
		fmt.Printf("Encrypted message: %s\n", encrypted)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
