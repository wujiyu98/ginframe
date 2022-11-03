package boot

import (
	_ "github.com/wujiyu98/ginframe/config"
	_ "github.com/wujiyu98/ginframe/database"
	"github.com/wujiyu98/ginframe/router"
)

func Start() {
	router.Run()
}
