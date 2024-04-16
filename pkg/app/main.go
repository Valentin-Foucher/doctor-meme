package app

import "github.com/Valentin-Foucher/doctor-meme/pkg/utils"

func main() {
	config := &utils.Configuration{}
	utils.ReadConfiguration(config)

}
