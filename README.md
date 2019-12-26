# SimplePKI Client

This library contains a simple client interface for:
  -  Generating auth tokens
  -  Creating Certificate Authorities and Intermediate Certificate Authorities
  -  Creating In Memory Key/Cert pairs
  -  Generating a *tls.Config* used for all mTLS HTTPS client/server configs


## Configuring the Client

The config for the client library is driven by eith ENV variable or a configuration file stored at `/etc/simplepki.yml`, `/etc/simplepki.yaml`, or `$HOME/.simplepki`.

### ENV Variable Configuration

The global variables are needed for all client operations.

#### Global Variables

  ##### SIMPLEPKI_ACCOUNT

    Account to generate certificate authorities, intermediate authorities, and certificates under.

  ##### SIMPLEPKI_ENDPOINT

    Endpoint serving the SimplePKI API/Lambdas

  ##### SIMPLEPKI_TOKEN

    JWT Token that providers authorization for a specified account to generate certificate authorities, intermediate authorities, and certificates.

#### Certificate Authority Creation Variables

  ##### SIMPLEPKI_CERTIFICATE_AUTHORITY

    Name of the certificate authority to create.

#### Intermediate Certificate Authority Creation Variables
	
  ##### SIMPLEPKI_CERTIFICATE_AUTHORITY

    Name of the certificate authority to sign the intermediate authority.

  ##### SIMPLEPKI_INTERMEDIATE_CERTIFICATE_AUTHORITY

    Name of the intermediate authority to create.

#### Certificate Signing/Creation Variables

  ##### SIMPLEPKI_CHAIN

    Name of the desired chain to sign the certificate with. For example, say a certificate authority *ImACA* exists as does an intermediate authority *ImAInter* and we wish to have the certificate *ImACert* signed by the chain formed by the two mentioned authorities. This value would then be `ImaCA/ImACert`.

  ##### SIMPLEPKI_ID

    Name/ID to include in the certificates identifying URI. From the chain example above; this value would be `ImACert`.

  ##### SIMPLEPKI_SUBJ_ALT_NAMES

    Subject Alternative Names to be added to the certificate.


#### Token Generation Varibales

  ##### SIMPLEPKI_TOKEN_GENERATOR
    
    Resource to use to generation auth tokens. Currently only supports calling AWS Lambda by ARN.

  ##### SIMPLEPKI_TOKEN_PREFIX

    Glob matching string that delegates which chains/certificate can be generated. For instance:
      - A prefix _*_ will allow any certificate authority, intermediate authority, and certificate to be generated for a given account.
      - A prefix _thisca/*_ will allow for any intermediate authority and certificate to be sign by/created from the CA chain with the root certificate authority _thisca_.
      - A prefic _thisca/thisinter/*_ will allow for any futher intermediate authority or certificate to be created from the chain 

  ##### SIMPLEPKI_TOKEN_TTL

    Time-to-live in hours for the generated token.
