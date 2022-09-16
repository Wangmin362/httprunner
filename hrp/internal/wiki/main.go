package wiki

import (
	"os/exec"

	"github.com/rs/zerolog/log"

	"gitcdteam.skyguardmis.com/gatorcloud/httprunner/hrp/internal/sdk"
)

func OpenWiki() error {
	sdk.SendEvent(sdk.EventTracking{
		Category: "OpenWiki",
		Action:   "hrp wiki",
	})
	log.Info().Msgf("%s https://httprunner.com", openCmd)
	return exec.Command(openCmd, "https://httprunner.com").Run()
}
