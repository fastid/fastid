package command

import (
	"bufio"
	"context"
	"fmt"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/migrations"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/fastid/fastid/internal/services"
	"github.com/ggwhite/go-masker"
	"golang.org/x/term"
	internalLog "log"
	"net/mail"
	"os"
	"strings"
	"syscall"
)

func CreateSuperUser() {
	// Configs
	cfg, err := config.New()
	if err != nil {
		internalLog.Fatalln(err.Error())
	}

	// Context
	ctx := context.Background()

	// Logger
	cfg.LOGGER.Level = "fatal"
	log := logger.New(cfg)

	// DB
	database, err := db.New(cfg, ctx)
	if err != nil {
		log.Fatal(ctx, err.Error())
	}

	// Migrations
	migration, err := migrations.New(cfg, database)
	if err != nil {
		log.Fatal(ctx, err.Error())
	}

	if err = migration.Upgrade(); err != nil {
		log.Infof(ctx, "migration %s", err.Error())
	}

	// Repository
	repos := repositories.New(cfg, log, database)

	// Service
	srv := services.New(cfg, log, repos)

	reader := bufio.NewReader(os.Stdin)
	var username string
	var password string
	var email string

	if cfg.APP.MasterID == "username" {
		for {
			fmt.Print("Enter your username:")
			data, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			data = strings.TrimSuffix(data, "\n")
			data = strings.TrimRight(data, " ")
			data = strings.TrimLeft(data, " ")

			if data != "" {
				username = data

				userData, err := srv.Users().GetByUsername(ctx, username)
				if err != nil {
					log.Fatal(ctx, err.Error())
					return
				}

				if userData.UserId.String() != "00000000-0000-0000-0000-000000000000" {
					fmt.Println("The username is already in use!")
					continue
				}

				break
			}
		}
	}

	for {
		fmt.Print("Enter your email address:")
		data, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		data = strings.TrimSuffix(data, "\n")
		data = strings.TrimRight(data, " ")
		data = strings.TrimLeft(data, " ")

		if data != "" {
			addr, err := mail.ParseAddress(data)
			if err != nil {
				continue
			}
			email = addr.Address

			userData, err := srv.Users().GetByEmail(ctx, email)
			if err != nil {
				log.Fatal(ctx, err.Error())
				return
			}

			if userData.UserId.String() != "00000000-0000-0000-0000-000000000000" {
				fmt.Println("The email address is already in use!")
				continue
			}

			break
		}
	}

	for {
		fmt.Print("Enter your password:")
		byteData, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			panic(err)
		}
		data := string(byteData)
		data = strings.TrimSuffix(data, "\n")
		data = strings.TrimRight(data, " ")
		data = strings.TrimLeft(data, " ")

		if data != "" {
			password = data

			if len(password) < cfg.PasswordMinLength {
				fmt.Printf("\nThe minimum password length must be %d\n", cfg.PasswordMinLength)
				continue
			}

			if len(password) > cfg.PasswordMaxLength {
				fmt.Printf("\nThe maximum password length must be %d\n", cfg.PasswordMaxLength)
				continue
			}

			break
		} else {
			fmt.Println("")
		}
	}

	fmt.Println("\n\n======= You have provided the following information =======")
	fmt.Printf("Username: %s \nEmail: %s \nPassword: %s", username, email, masker.Password(password))
	fmt.Println("\n===========================================================")

	for {
		fmt.Print("Create a user? (y/N):")
		data, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		data = strings.TrimSuffix(data, "\n")
		data = strings.TrimRight(data, " ")
		data = strings.TrimLeft(data, " ")
		data = strings.ToLower(data)

		if data == "n" {
			break
		} else if data == "y" {
			userData := services.UserData{
				Email:     email,
				Username:  username,
				Password:  password,
				Active:    true,
				SuperUser: true,
			}

			if username == "" {
				userData.Username = nil
			}

			err := srv.Users().Create(ctx, &userData)
			if err != nil {
				log.Fatal(ctx, err.Error())
				return
			}

			fmt.Printf("Create user %s\n", userData.UserId)
			break
		}
	}
}
