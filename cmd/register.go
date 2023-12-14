/*
Copyright © 2023 Mateusz Jenek <mateusz.jenek@gmail.com>

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

	"github.com/mateuszjenek/ddd-cli/internal/application"
	"github.com/mateuszjenek/ddd-cli/internal/infrastructure"
	"github.com/mateuszjenek/ddd-cli/internal/infrastructure/port"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var registerCustomerFirstName string
var registerCustomerLastName string
var registerCustomerEmail string

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register new customer",
	Long: `This command register new customer to the database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := port.NewSQLiteDatabasePort(viper.GetString("db_file"))
		if err != nil {
			return fmt.Errorf("failed to create new sqlite database port: %v", err)
		}
		customerService := infrastructure.NewSqlCustomerRepository(db)
		registerNewCustomer := application.NewRegisterNewCustomer(customerService)
		customer, err := registerNewCustomer.RegisterNewCustomer(registerCustomerFirstName, registerCustomerLastName, registerCustomerEmail)
		if err != nil {
			return fmt.Errorf("failed to register new customer: %v", err)
		}
		fmt.Printf("Customer created: %v", customer)
		return nil
	},
}

func init() {
	customerCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVar(&registerCustomerFirstName, "first-name", "FIRST_NAME", "First name of the customer")
	registerCmd.MarkFlagRequired("first-name")
	registerCmd.Flags().StringVar(&registerCustomerLastName, "last-name", "LAST_NAME", "Last name of the customer")
	registerCmd.MarkFlagRequired("last-name")
	registerCmd.Flags().StringVar(&registerCustomerEmail, "email", "EMAIL", "Email of the customer")
	registerCmd.MarkFlagRequired("email")
}
