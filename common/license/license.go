//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package license helps manage commercial licenses and check if they
// are valid for the version of UniOffice used.
package license ;import _b "github.com/ifanfairuz/unioffice/internal/license";

// SetMeteredKey sets the metered License API key required for SaaS operation.
// Document usage is reported periodically for the product to function correctly.
func SetMeteredKey (apiKey string )error {return _b .SetMeteredKey (apiKey )};

// SetMeteredKeyUsageLogVerboseMode sets the metered License API Key usage log verbose mode.
// Default value `false`, set to `true` will log the credit usages and print out to console with log level INFO.
func SetMeteredKeyUsageLogVerboseMode (val bool ){_b .SetMeteredKeyUsageLogVerboseMode (val )};

// LegacyLicenseType is the type of license
type LegacyLicenseType =_b .LegacyLicenseType ;

// SetMeteredKeyPersistentCache sets the metered License API Key persistent cache.
// Default value `true`, set to `false` will report the usage immediately to license server,
// this can be used when there's no access to persistent data storage.
func SetMeteredKeyPersistentCache (val bool ){_b .SetMeteredKeyPersistentCache (val )};

// LicenseKey represents a loaded license key.
type LicenseKey =_b .LicenseKey ;

// GetMeteredState checks the currently used metered document usage status,
// documents used and credits available.
func GetMeteredState ()(_b .MeteredStatus ,error ){return _b .GetMeteredState ()};

// LegacyLicense holds the old-style unioffice license information.
type LegacyLicense =_b .LegacyLicense ;

// GetLicenseKey returns the currently loaded license key.
func GetLicenseKey ()*LicenseKey {return _b .GetLicenseKey ()};const (LicenseTierUnlicensed =_b .LicenseTierUnlicensed ;LicenseTierCommunity =_b .LicenseTierCommunity ;LicenseTierIndividual =_b .LicenseTierIndividual ;LicenseTierBusiness =_b .LicenseTierBusiness ;
);

// SetLegacyLicenseKey installs a legacy license code. License codes issued prior to June 2019.
// Will be removed at some point in a future major version.
func SetLegacyLicenseKey (s string )error {return _b .SetLegacyLicenseKey (s )};

// SetLicenseKey sets and validates the license key.
func SetLicenseKey (content string ,customerName string )error {return _b .SetLicenseKey (content ,customerName );};

// MakeUnlicensedKey returns a default key.
func MakeUnlicensedKey ()*LicenseKey {return _b .MakeUnlicensedKey ()};