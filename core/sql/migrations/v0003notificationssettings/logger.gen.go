// Code generated by berty.tech/core/.scripts/generate-logger.sh

package v0003notificationssettings

import "go.uber.org/zap"

func logger() *zap.Logger {
	return zap.L().Named("core.sql.migrations.v0003notificationssettings")
}
