/*******************************************************************************
* Copyright 2019 Stefan Majewsky <majewsky@gmx.net>
* SPDX-License-Identifier: GPL-3.0-only
* Refer to the file "LICENSE" for details.
*******************************************************************************/

package main

import (
	"os"
	"regexp"

	"github.com/sapcc/go-bits/logg"
	osutil_user "github.com/tredoe/osutil/user"
)

var (
	envDefaults = map[string]string{
		//empty value = not optional
		"PORTUNUS_DEBUG":              "false",
		"PORTUNUS_LDAP_SUFFIX":        "",
		"PORTUNUS_SERVER_BINARY":      "portunus-server",
		"PORTUNUS_SERVER_GROUP":       "portunus",
		"PORTUNUS_SERVER_HTTP_LISTEN": "127.0.0.1:8080",
		"PORTUNUS_SERVER_HTTP_SECURE": "true",
		"PORTUNUS_SERVER_STATE_DIR":   "/var/lib/portunus",
		"PORTUNUS_SERVER_USER":        "portunus",
		"PORTUNUS_SLAPD_BINARY":       "slapd",
		"PORTUNUS_SLAPD_GROUP":        "ldap",
		"PORTUNUS_SLAPD_SCHEMA_DIR":   "/etc/openldap/schema",
		"PORTUNUS_SLAPD_STATE_DIR":    "/var/run/portunus-slapd",
		"PORTUNUS_SLAPD_USER":         "ldap",
	}

	boolRx        = regexp.MustCompile(`^(?:true|false)$`)
	ldapSuffixRx  = regexp.MustCompile(`^dc=[a-z0-9_-]+(?:,dc=[a-z0-9_-]+)*$`)
	userOrGroupRx = regexp.MustCompile(`^[a-z_][a-z0-9_-]*\$?$`)
	envFormats    = map[string]*regexp.Regexp{
		"PORTUNUS_DEBUG":              boolRx,
		"PORTUNUS_LDAP_SUFFIX":        ldapSuffixRx,
		"PORTUNUS_SERVER_GROUP":       userOrGroupRx,
		"PORTUNUS_SERVER_HTTP_LISTEN": regexp.MustCompile(`^(?:[0-9.]+|\[[0-9a-f:]+\]):[0-9]+$`),
		"PORTUNUS_SERVER_HTTP_SECURE": boolRx,
		"PORTUNUS_SERVER_USER":        userOrGroupRx,
		"PORTUNUS_SLAPD_GROUP":        userOrGroupRx,
		"PORTUNUS_SLAPD_USER":         userOrGroupRx,
	}
)

func readConfig() (environment map[string]string, ids map[string]int) {
	//last-minute initializations in envDefaults
	if os.Getenv("PORTUNUS_SLAPD_TLS_CERTIFICATE") != "" {
		envDefaults["PORTUNUS_SLAPD_TLS_CERTIFICATE"] = ""
		envDefaults["PORTUNUS_SLAPD_TLS_DOMAIN_NAME"] = ""
		envDefaults["PORTUNUS_SLAPD_TLS_PRIVATE_KEY"] = ""
		envDefaults["PORTUNUS_SLAPD_TLS_CA_CERTIFICATE"] = ""
	}

	//read and validate all relevant environment variables
	environment = make(map[string]string)
	for key, defaultValue := range envDefaults {
		value := os.Getenv(key)
		if value == "" {
			value = defaultValue
		}
		if value == "" {
			logg.Fatal("missing required environment variable: " + key)
		}
		if rx := envFormats[key]; rx != nil {
			if !rx.MatchString(value) {
				logg.Fatal("malformed environment variable: %s must look like /%s/", value, rx.String())
			}
		}
		environment[key] = value
		os.Unsetenv(key) //avoid unintentional leakage of env vars to child processes
	}

	//resolve user/group names into IDs
	ids = map[string]int{
		"PORTUNUS_SERVER_UID": getUIDForName(environment["PORTUNUS_SERVER_USER"]),
		"PORTUNUS_SERVER_GID": getGIDForName(environment["PORTUNUS_SERVER_GROUP"]),
		"PORTUNUS_SLAPD_UID":  getUIDForName(environment["PORTUNUS_SLAPD_USER"]),
		"PORTUNUS_SLAPD_GID":  getGIDForName(environment["PORTUNUS_SLAPD_GROUP"]),
	}

	return
}

func getGIDForName(name string) int {
	group, err := osutil_user.LookupGroup(name)
	if err != nil {
		logg.Fatal("cannot find group %s: %s", name, err.Error())
	}
	if group == nil {
		logg.Fatal("cannot find group %s", name)
	}
	return group.GID
}

func getUIDForName(name string) int {
	user, err := osutil_user.LookupUser(name)
	if err != nil {
		logg.Fatal("cannot find user %s: %s", name, err.Error())
	}
	if user == nil {
		logg.Fatal("cannot find user %s", name)
	}
	return user.UID
}
