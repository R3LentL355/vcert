/*
 * Copyright 2018 Venafi, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"github.com/Venafi/vcert/pkg/certificate"
)

const (
	commandGenCSRName  = "gencsr"
	commandEnrollName  = "enroll"
	commandPickupName  = "pickup"
	commandRevokeName  = "revoke"
	commandRenewName   = "renew"
	commandGetcredName = "getcred"
)

var (
	genCsrFlags  = flag.NewFlagSet(commandGenCSRName, flag.PanicOnError)
	genCsrParams commandFlags

	enrollFlags  = flag.NewFlagSet(commandEnrollName, flag.PanicOnError)
	enrollParams commandFlags

	pickupFlags = flag.NewFlagSet(commandPickupName, flag.PanicOnError)
	pickParams  commandFlags

	revokeFlags  = flag.NewFlagSet(commandRevokeName, flag.PanicOnError)
	revokeParams commandFlags

	renewFlags  = flag.NewFlagSet(commandRenewName, flag.PanicOnError)
	renewParams commandFlags

	getcredFlags  = flag.NewFlagSet(commandGetcredName, flag.PanicOnError)
	getcredParams commandFlags

	flags commandFlags
)

type commandFlags struct {
	verbose            bool
	url                string
	tppURL             string //deprecated
	tppUser            string
	tppPassword        string
	tppToken           string
	apiKey             string
	cloudURL           string //deprecated
	zone               string
	caDN               string
	csrOption          string
	keyType            certificate.KeyType
	keyTypeString      string
	keySize            int
	keyCurve           certificate.EllipticCurve
	keyCurveString     string
	keyPassword        string
	friendlyName       string
	commonName         string
	distinguishedName  string
	thumbprint         string
	org                string
	country            string
	state              string
	locality           string
	orgUnits           stringSlice
	dnsSans            stringSlice
	ipSans             ipSlice
	emailSans          emailSlice
	format             string
	file               string
	keyFile            string
	csrFile            string
	certFile           string
	chainFile          string
	chainOption        string
	noPrompt           bool
	pickupID           string
	trustBundle        string
	noPickup           bool
	testMode           bool
	testModeDelay      int
	revocationReason   string
	revocationNoRetire bool
	pickupIDFile       string
	timeout            int
	insecure           bool
	config             string
	profile            string
	clientP12          string
	clientP12PW        string
	clientId           string
	scope              string
}
