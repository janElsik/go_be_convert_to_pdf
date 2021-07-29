package convert

import (
	"fmt"
	"github.com/gofiber/fiber"
	"os"
	"os/exec"
	"strings"
)

func Convert(c *fiber.Ctx) {
	file, err := c.FormFile("data")
	if err != nil {
		fmt.Println("err formfile:", err)
		c.SendStatus(500)
		return
	}

	err = c.SaveFile(file, file.Filename)
	if err != nil {
		fmt.Println("Err downloadingFile: ", err)

	}

	cmd := exec.Command("unoconv", "-f", "pdf", file.Filename)
	if err := cmd.Run(); err != nil {
		fmt.Printf("error with converting to pdf: %v \n", err)
		fmt.Println(file.Filename)
	}

	tempSlice := strings.Split(file.Filename, ".")
	tempSlice[len(tempSlice)-1] = ".pdf"
	pdfFileName := strings.Join(tempSlice, "")

	err = c.SendFile(pdfFileName)
	if err != nil {
		fmt.Println("Err sending file, ", err)
	}

	fmt.Println(pdfFileName)

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			fmt.Println("Err removing file:", err)
		}
	}(file.Filename)

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			fmt.Println("Err removing file:", err)
		}
	}(pdfFileName)
}
