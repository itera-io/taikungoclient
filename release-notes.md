### New Endpoints: 16
---------------------
GET /api/v1/certificate-credentials  
POST /api/v1/certificate-credentials  
PUT /api/v1/certificate-credentials  
GET /api/v1/certificate-credentials/list  
POST /api/v1/certificate-credentials/lockmanager  
POST /api/v1/certificate-credentials/makedefault  
POST /api/v1/certificate-credentials/validate  
DELETE /api/v1/certificate-credentials/{id}  
GET /api/v1/custom-cas  
POST /api/v1/custom-cas  
PUT /api/v1/custom-cas  
GET /api/v1/custom-cas/list  
POST /api/v1/custom-cas/lockmanager  
POST /api/v1/custom-cas/makedefault  
POST /api/v1/custom-cas/validate  
DELETE /api/v1/custom-cas/{id}  

### Deleted Endpoints: 12
-------------------------
GET /api/v1/certificate-profiles  
POST /api/v1/certificate-profiles  
PUT /api/v1/certificate-profiles  
GET /api/v1/certificate-profiles/list  
POST /api/v1/certificate-profiles/lockmanager  
POST /api/v1/certificate-profiles/makedefault  
POST /api/v1/certificate-profiles/parse-certificate  
POST /api/v1/certificate-profiles/validate  
POST /api/v1/certificate-profiles/validate-trust-bundle  
DELETE /api/v1/certificate-profiles/{id}  
GET /api/v1/certificate-profiles/{id}/trust-bundle  
POST /api/v1/cronjob/scan-certificate-profile-expiry  

### Modified Endpoints: 7
-------------------------
GET /api/v1/aws/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Deleted property: importedVpcCidr
                  - Deleted property: importedVpcId
                  - Deleted property: ipMode
                  - Deleted property: vpcMode

POST /api/v1/cronjob/scan-certificate-credential-expiry
- Summary changed from 'Legacy alias for scan-certificate-profile-expiry (pre-rename cron path)' to 'Scan certificate credentials approaching or past expiry and emit alert logs'
- OperationID changed from 'CronJobScanCertificateCredentialExpiryLegacy' to 'cronjob-scan-certificate-credential-expiry'

POST /api/v1/dns-cert/cert-manager/enable
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - New property: certificateCredentialId
          - New property: customCertificateAuthorityId
          - Deleted property: certificateProfileId

POST /api/v1/dns-cert/enable
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - New property: certificateCredentialId
          - New property: customCertificateAuthorityId
          - Deleted property: certificateProfileId

POST /api/v1/dns-cert/validate
- Summary changed from 'Validate DNS/Cert enablement inputs (provider credentials, signing CA reference, issuer rules)' to 'Validate DNS provider credentials'
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - New property: caCert
          - New property: caKey
          - New property: certificateCredentialId
          - New property: customCertificateAuthorityId
          - Deleted property: certificateProfileId

GET /api/v1/dns-cert/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Deleted property: certificateProfileId
            - Deleted property: certificateProfileName
            - Deleted property: enabled

POST /api/v1/dnscredentials/validate
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - New property: caCert
          - New property: caKey
          - New property: certificateCredentialId
          - New property: customCertificateAuthorityId
          - Deleted property: certificateProfileId


