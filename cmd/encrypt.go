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
	"github.com/spf13/viper"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a message using the Enigma machine.",
	Long:  `Encrypt a message using the Enigma machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cobra.CheckErr(fmt.Errorf("you must provide a message to encrypt"))
		}
		message := strings.Trim(args[0], " ")
		if message == "" {
			cobra.CheckErr(fmt.Errorf("you must provide a message to encrypt"))
		}

		reflectorSelection := viper.GetString("reflector")
		if reflectorSelection == "" {
			cobra.CheckErr(fmt.Errorf("reflector default not set this should not happen"))
		}

		rotorSelection := viper.GetStringSlice("rotors")

		if len(rotorSelection) == 0 {
			cobra.CheckErr(fmt.Errorf("rotor default not set this should not happen"))
		}

		rotorPositions := viper.GetString("rotor-positions")
		if len(rotorPositions) == 0 {
			cobra.CheckErr(fmt.Errorf("rotor positions default not set this should not happen"))
		}
		if len(rotorSelection) != len(rotorPositions) {
			cobra.CheckErr(fmt.Errorf("rotor selection and rotor positions must have the same length"))
		}

		plugboardPairsSelection := viper.GetStringSlice("plugboard.pairs")
		if len(plugboardPairsSelection) > 10 {
			cobra.CheckErr(fmt.Errorf("plugboard pairs must be 10 or fewer"))
		}

		// create enigma machine
		plugboard := enigma.NewPlugboard()

		reflector, err := enigma.CreateReflectorFromSelection(reflectorSelection)
		cobra.CheckErr(err)

		rotors := make([]*enigma.Rotor, len(rotorSelection))
		for i, rotorName := range rotorSelection {
			rotor, err := enigma.CreateRotorFromSelection(rotorName)
			cobra.CheckErr(err)
			rotors[i] = rotor
		}

		em := enigma.NewEnigmaMachine(plugboard, rotors, reflector)

		err = em.SetRotorPositions(strings.Split(rotorPositions, ""))
		cobra.CheckErr(err)

		if len(plugboardPairsSelection) > 0 {
			for _, pair := range plugboardPairsSelection {
				if len(pair) != 2 {
					cobra.CheckErr(fmt.Errorf("plugboard pairs must be two characters long"))
				}
				a := rune(pair[0])
				b := rune(pair[1])
				em.AddPlugboardConnection(a, b)
			}
		}

		encrypted, err := em.EncryptString(message)
		cobra.CheckErr(err)

		fmt.Printf(`
Enigma machine settings used:
- Reflector: %s
- Rotors: %s
- Rotor positions: %s
- Plugboard pairs: %s

`,
			reflectorSelection,
			rotorSelection,
			rotorPositions,
			plugboardPairsSelection,
		)

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
	encryptCmd.Flags().StringP("reflector", "u", "", "Reflector to use")
	encryptCmd.Flags().StringSliceP("rotors", "r", []string{}, "Rotors to use")
	encryptCmd.Flags().StringP("rotor-positions", "d", "", "Rotor positions to use")
	encryptCmd.Flags().StringSliceP("plugboard-pairs", "p", []string{}, "Plugboard pairs to use")

	viper.BindPFlag("reflector", encryptCmd.Flags().Lookup("reflector"))
	viper.BindPFlag("rotors", encryptCmd.Flags().Lookup("rotors"))
	viper.BindPFlag("rotor-positions", encryptCmd.Flags().Lookup("rotor-positions"))
	viper.BindPFlag("plugboard.pairs", encryptCmd.Flags().Lookup("plugboard-pairs"))

	viper.SetDefault("reflector", "B")
	viper.SetDefault("rotors", []string{"III", "II", "I"})
	viper.SetDefault("rotor-positions", "AAA")
	viper.SetDefault("plugboard.pairs", []string{})
}
