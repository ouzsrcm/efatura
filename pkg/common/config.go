package invoice

import (
	"github.com/google/uuid"
)

var PRODUCTION = true

const PROD_URL, TEST_URL = "https://earsivportal.efatura.gov.tr", "https://earsivportaltest.efatura.gov.tr"

func getUrl() string {
	if PRODUCTION {
		return PROD_URL
	}
	return TEST_URL
}

func getUUID() string {
	return uuid.New().String()
}

func enableTestMode() {
	PRODUCTION = false
}
