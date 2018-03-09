package courier

import (
	"fmt"
	"os"
	"strconv"
)

var (
	//CourierPort is the port which the application should run on
	CourierPort = envMust("COURIER_PORT")

	//SMTPServerPort is the port of your SMTP server of choice
	SMTPServerPort = mustInt(envMust("SMTP_PORT"))

	// SMTPHost is the host of your SMTP Host
	SMTPHost = envMust("SMTP_HOST")

	// SMTPPassword is your authentication method
	SMTPPassword = envMust("SMTP_PASSWORD")

	// SMTPUser is your authentication method
	SMTPUser = envMust("SMTP_USER")
)

func envMust(key string) (val string) {
	val = os.Getenv(key)
	if val == "" {
		panic(fmt.Errorf("No environment variable defined for %s", key))
	}
	return val
}

func mustInt(variable string) (i int64) {
	var err error
	if i, err = strconv.ParseInt(variable, 10, 64); err != nil {
		panic(err.Error())
	}
	return i
}
