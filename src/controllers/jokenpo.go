package controllers

import (
	"github.com/LulianoM/GamesSimulatorAPI/src/structs"
	"github.com/gofiber/fiber/v2"
)

func Jokenpo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	round := structs.Jokenpo{
		UserChoice:     data["userChoice"],
		ComputerChoice: data["computerChoice"],
	}
	round.YouWin = winJokenpo(round)
	return c.JSON(round)
}

func winJokenpo(round structs.Jokenpo) bool {
	if round.UserChoice == "rock" {
		if round.ComputerChoice == "paper" {
			return false
		} else {
			return true
		}
	}
	if round.UserChoice == "paper" {
		if round.ComputerChoice == "scissors" {
			return false
		} else {
			return true
		}
	}
	if round.UserChoice == "scissors" {
		if round.ComputerChoice == "rock" {
			return false
		} else {
			return true
		}
	}
	return false
}
