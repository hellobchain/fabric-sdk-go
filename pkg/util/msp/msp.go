package msp

import (
	"encoding/pem"
	"github.com/pkg/errors"
	"github.com/wsw365904/fabric-sdk-go/internal/github.com/hyperledger/fabric/bccsp"
	internalmsp "github.com/wsw365904/fabric-sdk-go/internal/github.com/hyperledger/fabric/msp"
	"github.com/wsw365904/fabric-sdk-go/pkg/util/algo"
	"github.com/wsw365904/newcryptosm/x509"
	"github.com/wsw365904/newcryptosm/x509/pkix"
	"github.com/wsw365904/third_party/hyperledger/fabric-config/configtx"
	"github.com/wsw365904/third_party/hyperledger/fabric-config/configtx/membership"
	"github.com/wsw365904/wswlog/wlogging"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var mspLogger = wlogging.MustGetLoggerWithoutName()

const (
	cacerts              = "cacerts"
	admincerts           = "admincerts"
	signcerts            = "signcerts"
	keystore             = "keystore"
	intermediatecerts    = "intermediatecerts"
	crlsfolder           = "crls"
	configfilename       = "config.yaml"
	tlscacerts           = "tlscacerts"
	tlsintermediatecerts = "tlsintermediatecerts"
)

func getPemMaterialCrlFromDir(dir string) ([]*pkix.CertificateList, error) {
	mspLogger.Debugf("Reading directory %s", dir)

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return nil, err
	}

	content := make([]*pkix.CertificateList, 0)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read directory %s", dir)
	}

	for _, f := range files {
		fullName := filepath.Join(dir, f.Name())

		f, err := os.Stat(fullName)
		if err != nil {
			mspLogger.Warningf("Failed to stat %s: %s", fullName, err)
			continue
		}
		if f.IsDir() {
			continue
		}

		mspLogger.Debugf("Inspecting file %s", fullName)

		item, err := readPemFileToCrl(fullName)
		if err != nil {
			mspLogger.Warningf("Failed reading file %s: %s", fullName, err)
			continue
		}

		content = append(content, item)
	}

	return content, nil
}

func getPemMaterialCertificatesFromDir(dir string) ([]*x509.Certificate, error) {
	mspLogger.Debugf("Reading directory %s", dir)

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return nil, err
	}

	content := make([]*x509.Certificate, 0)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read directory %s", dir)
	}

	for _, f := range files {
		fullName := filepath.Join(dir, f.Name())

		f, err := os.Stat(fullName)
		if err != nil {
			mspLogger.Warningf("Failed to stat %s: %s", fullName, err)
			continue
		}
		if f.IsDir() {
			continue
		}

		mspLogger.Debugf("Inspecting file %s", fullName)

		item, err := readPemFileToCertificate(fullName)
		if err != nil {
			mspLogger.Warningf("Failed reading file %s: %s", fullName, err)
			continue
		}

		content = append(content, item)
	}

	return content, nil
}

func readPemFileToCertificate(file string) (*x509.Certificate, error) {
	bytes, err := readFile(file)
	if err != nil {
		return nil, errors.Wrapf(err, "reading from file %s failed", file)
	}

	b, _ := pem.Decode(bytes)
	if b == nil { // TODO: also check that the type is what we expect (cert vs key..)
		return nil, errors.Errorf("no pem content for file %s", file)
	}
	return x509.ParseCertificate(b.Bytes)
}

func readPemFileToCrl(file string) (*pkix.CertificateList, error) {
	bytes, err := readFile(file)
	if err != nil {
		return nil, errors.Wrapf(err, "reading from file %s failed", file)
	}

	b, _ := pem.Decode(bytes)
	if b == nil { // TODO: also check that the type is what we expect (cert vs key..)
		return nil, errors.Errorf("no pem content for file %s", file)
	}
	return x509.ParseCRL(b.Bytes)
}

func readFile(file string) ([]byte, error) {
	fileCont, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read file %s", file)
	}

	return fileCont, nil
}

func loadCertificateAt(dir, certificatePath string, ouType string) *x509.Certificate {

	if certificatePath == "" {
		mspLogger.Debugf("Specific certificate for %s is not configured", ouType)
		return nil
	}

	f := filepath.Join(dir, certificatePath)
	raw, err := readPemFileToCertificate(f)
	if err != nil {
		mspLogger.Warnf("Failed loading %s certificate at [%s]: [%s]", ouType, f, err)
	} else {
		return raw
	}
	return nil
}

// GetVerifyingMspConfig returns an MSP config given directory, ID and type
func GetVerifyingMspConfig(dir, ID, mspType string) (*configtx.MSP, error) {
	switch mspType {
	case internalmsp.ProviderTypeToString(internalmsp.FABRIC):
		return getMspConfig(dir, ID)
	default:
		return nil, errors.Errorf("unknown MSP type '%s'", mspType)
	}
}

func getMspConfig(dir string, ID string) (*configtx.MSP, error) {
	cacertDir := filepath.Join(dir, cacerts)
	admincertDir := filepath.Join(dir, admincerts)
	intermediatecertsDir := filepath.Join(dir, intermediatecerts)
	crlsDir := filepath.Join(dir, crlsfolder)
	configFile := filepath.Join(dir, configfilename)
	tlscacertDir := filepath.Join(dir, tlscacerts)
	tlsintermediatecertsDir := filepath.Join(dir, tlsintermediatecerts)

	cacerts, err := getPemMaterialCertificatesFromDir(cacertDir)
	if err != nil || len(cacerts) == 0 {
		return nil, errors.WithMessagef(err, "could not load a valid ca certificate from directory %s", cacertDir)
	}

	admincert, err := getPemMaterialCertificatesFromDir(admincertDir)
	if err != nil && !os.IsNotExist(err) {
		return nil, errors.WithMessagef(err, "could not load a valid admin certificate from directory %s", admincertDir)
	}

	intermediatecerts, err := getPemMaterialCertificatesFromDir(intermediatecertsDir)
	if os.IsNotExist(err) {
		mspLogger.Debugf("Intermediate certs folder not found at [%s]. Skipping. [%s]", intermediatecertsDir, err)
	} else if err != nil {
		return nil, errors.WithMessagef(err, "failed loading intermediate ca certs at [%s]", intermediatecertsDir)
	}

	tlsCACerts, err := getPemMaterialCertificatesFromDir(tlscacertDir)
	tlsIntermediateCerts := []*x509.Certificate{}
	if os.IsNotExist(err) {
		mspLogger.Debugf("TLS CA certs folder not found at [%s]. Skipping and ignoring TLS intermediate CA folder. [%s]", tlsintermediatecertsDir, err)
	} else if err != nil {
		return nil, errors.WithMessagef(err, "failed loading TLS ca certs at [%s]", tlsintermediatecertsDir)
	} else if len(tlsCACerts) != 0 {
		tlsIntermediateCerts, err = getPemMaterialCertificatesFromDir(tlsintermediatecertsDir)
		if os.IsNotExist(err) {
			mspLogger.Debugf("TLS intermediate certs folder not found at [%s]. Skipping. [%s]", tlsintermediatecertsDir, err)
		} else if err != nil {
			return nil, errors.WithMessagef(err, "failed loading TLS intermediate ca certs at [%s]", tlsintermediatecertsDir)
		}
	} else {
		mspLogger.Debugf("TLS CA certs folder at [%s] is empty. Skipping.", tlsintermediatecertsDir)
	}

	crls, err := getPemMaterialCrlFromDir(crlsDir)
	if os.IsNotExist(err) {
		mspLogger.Debugf("crls folder not found at [%s]. Skipping. [%s]", crlsDir, err)
	} else if err != nil {
		return nil, errors.WithMessagef(err, "failed loading crls at [%s]", crlsDir)
	}

	// Load configuration file
	// if the configuration file is there then load it
	// otherwise skip it
	var ouis []membership.OUIdentifier
	var nodeOUs membership.NodeOUs
	_, err = os.Stat(configFile)
	if err == nil {
		// load the file, if there is a failure in loading it then
		// return an error
		raw, err := ioutil.ReadFile(configFile)
		if err != nil {
			return nil, errors.Wrapf(err, "failed loading configuration file at [%s]", configFile)
		}

		configuration := internalmsp.Configuration{}
		err = yaml.Unmarshal(raw, &configuration)
		if err != nil {
			return nil, errors.Wrapf(err, "failed unmarshalling configuration file at [%s]", configFile)
		}

		// Prepare OrganizationalUnitIdentifiers
		if len(configuration.OrganizationalUnitIdentifiers) > 0 {
			for _, ouID := range configuration.OrganizationalUnitIdentifiers {
				f := filepath.Join(dir, ouID.Certificate)
				raw, err := readPemFileToCertificate(f)
				if err != nil {
					return nil, errors.Wrapf(err, "failed loading OrganizationalUnit certificate at [%s]", f)
				}

				oui := membership.OUIdentifier{
					Certificate:                  raw,
					OrganizationalUnitIdentifier: ouID.OrganizationalUnitIdentifier,
				}
				ouis = append(ouis, oui)
			}
		}

		// Prepare NodeOUs
		if configuration.NodeOUs != nil && configuration.NodeOUs.Enable {
			mspLogger.Debug("Loading NodeOUs")
			nodeOUs = membership.NodeOUs{
				Enable: true,
			}
			if configuration.NodeOUs.ClientOUIdentifier != nil && len(configuration.NodeOUs.ClientOUIdentifier.OrganizationalUnitIdentifier) != 0 {
				nodeOUs.ClientOUIdentifier = membership.OUIdentifier{OrganizationalUnitIdentifier: configuration.NodeOUs.ClientOUIdentifier.OrganizationalUnitIdentifier}
			}
			if configuration.NodeOUs.PeerOUIdentifier != nil && len(configuration.NodeOUs.PeerOUIdentifier.OrganizationalUnitIdentifier) != 0 {
				nodeOUs.PeerOUIdentifier = membership.OUIdentifier{OrganizationalUnitIdentifier: configuration.NodeOUs.PeerOUIdentifier.OrganizationalUnitIdentifier}
			}
			if configuration.NodeOUs.AdminOUIdentifier != nil && len(configuration.NodeOUs.AdminOUIdentifier.OrganizationalUnitIdentifier) != 0 {
				nodeOUs.AdminOUIdentifier = membership.OUIdentifier{OrganizationalUnitIdentifier: configuration.NodeOUs.AdminOUIdentifier.OrganizationalUnitIdentifier}
			}
			if configuration.NodeOUs.OrdererOUIdentifier != nil && len(configuration.NodeOUs.OrdererOUIdentifier.OrganizationalUnitIdentifier) != 0 {
				nodeOUs.OrdererOUIdentifier = membership.OUIdentifier{OrganizationalUnitIdentifier: configuration.NodeOUs.OrdererOUIdentifier.OrganizationalUnitIdentifier}
			}

			// Read certificates, if defined

			// ClientOU
			if nodeOUs.ClientOUIdentifier.Certificate == nil {
				nodeOUs.ClientOUIdentifier.Certificate = loadCertificateAt(dir, configuration.NodeOUs.ClientOUIdentifier.Certificate, "ClientOU")
			}
			// PeerOU
			if nodeOUs.PeerOUIdentifier.Certificate == nil {
				nodeOUs.PeerOUIdentifier.Certificate = loadCertificateAt(dir, configuration.NodeOUs.PeerOUIdentifier.Certificate, "PeerOU")
			}
			// AdminOU
			if nodeOUs.AdminOUIdentifier.Certificate == nil {
				nodeOUs.AdminOUIdentifier.Certificate = loadCertificateAt(dir, configuration.NodeOUs.AdminOUIdentifier.Certificate, "AdminOU")
			}
			// OrdererOU
			if nodeOUs.OrdererOUIdentifier.Certificate == nil {
				nodeOUs.OrdererOUIdentifier.Certificate = loadCertificateAt(dir, configuration.NodeOUs.OrdererOUIdentifier.Certificate, "OrdererOU")
			}
		}
	} else {
		mspLogger.Debugf("MSP configuration file not found at [%s]: [%s]", configFile, err)
	}

	// Set FabricCryptoConfig
	cryptoConfig := membership.CryptoConfig{
		SignatureHashFamily:            bccsp.SHA2,
		IdentityIdentifierHashFunction: bccsp.SHA256,
	}

	// Compose FabricMSPConfig
	fmspconf := &configtx.MSP{
		Name:                          ID,
		RootCerts:                     cacerts,
		IntermediateCerts:             intermediatecerts,
		Admins:                        admincert,
		RevocationList:                crls,
		OrganizationalUnitIdentifiers: ouis,
		CryptoConfig:                  cryptoConfig,
		TLSRootCerts:                  tlsCACerts,
		TLSIntermediateCerts:          tlsIntermediateCerts,
		NodeOUs:                       nodeOUs,
	}

	// wsw add
	if algo.GetGMFlag() {
		mspLogger.Debug("gm alg")
		cryptoConfig = membership.CryptoConfig{
			SignatureHashFamily:            bccsp.SM,
			IdentityIdentifierHashFunction: bccsp.SM3,
		}
		fmspconf.CryptoConfig = cryptoConfig
	}

	return fmspconf, nil
}
