package pytest

import (
	"gitcdteam.skyguardmis.com/gatorcloud/httprunner/hrp/internal/builtin"
	"gitcdteam.skyguardmis.com/gatorcloud/httprunner/hrp/internal/sdk"
)

func RunPytest(args []string) error {
	sdk.SendEvent(sdk.EventTracking{
		Category: "RunAPITests",
		Action:   "hrp pytest",
	})

	args = append([]string{"run"}, args...)
	return builtin.ExecPython3Command("httprunner", args...)
}
