### New Endpoints: None
-----------------------

### Deleted Endpoints: 22
-------------------------
GET /api/v1/certificate-credentials  
POST /api/v1/certificate-credentials  
PUT /api/v1/certificate-credentials  
GET /api/v1/certificate-credentials/list  
POST /api/v1/certificate-credentials/lockmanager  
POST /api/v1/certificate-credentials/makedefault  
POST /api/v1/certificate-credentials/validate  
DELETE /api/v1/certificate-credentials/{id}  
GET /api/v1/cronjob/ccf-crypto/status  
POST /api/v1/cronjob/reconcile-partner-scopes  
POST /api/v1/cronjob/rotate-tenant-keks  
POST /api/v1/cronjob/scan-certificate-credential-expiry  
POST /api/v1/dns-cert/cert-manager/disable  
POST /api/v1/dns-cert/cert-manager/enable  
POST /api/v1/dns-cert/external-dns/disable  
POST /api/v1/dns-cert/external-dns/enable  
POST /api/v1/dns-cert/trust-manager/disable  
POST /api/v1/dns-cert/trust-manager/enable  
POST /api/v1/gateway-api/disable  
POST /api/v1/gateway-api/enable  
GET /api/v1/users/overload-protection  
POST /api/v1/users/toggleoverloadprotection  

### Modified Endpoints: 47
--------------------------
GET /api/v1/accessprofiles
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

GET /api/v1/alertingprofiles
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

POST /api/v1/aws/efs-file-system-by-vpc-list
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: externalId
          - Deleted property: roleArn

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
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

POST /api/v1/aws/regions
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: externalId
          - Deleted property: roleArn

POST /api/v1/aws/subnet-list
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: externalId
          - Deleted property: roleArn
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Items changed
            - Properties changed
              - Deleted property: ipv6Cidr
              - Modified property: availableIpCount
                - Nullable changed from true to false

POST /api/v1/aws/vpc-list
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: externalId
          - Deleted property: roleArn

POST /api/v1/aws/zones
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: externalId
          - Deleted property: roleArn

POST /api/v1/azure/create
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: importedFileShareResourceGroupName
          - Modified property: subnets
            - Items changed
              - Properties changed
                - New property: subnetType

GET /api/v1/azure/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Required changed
                  - Deleted required property: importedFileShareResourceGroupName
                - Properties changed
                  - Deleted property: importedFileShareResourceGroupName
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

POST /api/v1/azure/vnet-subnet-list
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: ipMode

POST /api/v1/checker/aws
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: externalId
          - Deleted property: roleArn

POST /api/v1/custom-cas
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: commonName
          - Deleted property: generateSelfSigned
          - Deleted property: validityDays
          - Modified property: caCert
            - Nullable changed from true to false
          - Modified property: caKey
            - Nullable changed from true to false

GET /api/v1/custom-cas/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Required changed
                  - Deleted required property: caCert
                - Properties changed
                  - Deleted property: caCert

POST /api/v1/dns-cert/enable
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: certificateCredentialId

POST /api/v1/dns-cert/validate
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: authMethod
          - Deleted property: authVariant
          - Deleted property: awsRoleArn
          - Deleted property: certificateCredentialId
          - Deleted property: customCertificateAuthorityId
          - Deleted property: dnsCredentialId
          - Deleted property: gcpServiceAccountEmail
          - Deleted property: pdnsApiKey
          - Deleted property: pdnsApiUrl
          - Deleted property: pdnsServerId
          - Deleted property: tsigAlgorithm

GET /api/v1/dns-cert/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - New property: enabled
            - Deleted property: appliedCertManagerVersion
            - Deleted property: appliedExternalDnsVersion
            - Deleted property: appliedIssuerVersion
            - Deleted property: appliedTrustManagerVersion
            - Deleted property: certManagerEnabled
            - Deleted property: desiredCertManagerVersion
            - Deleted property: desiredExternalDnsVersion
            - Deleted property: desiredIssuerVersion
            - Deleted property: desiredTrustManagerVersion
            - Deleted property: externalDnsEnabled
            - Deleted property: failureReason
            - Deleted property: projectStatus
            - Deleted property: trustManagerEnabled

GET /api/v1/dnscredentials
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Items changed
            - Required changed
              - Deleted required property: providerType
            - Properties changed
              - Deleted property: providerType

POST /api/v1/dnscredentials
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: authMethod
          - Deleted property: authVariant
          - Deleted property: awsRoleArn
          - Deleted property: gcpServiceAccountEmail
          - Deleted property: pdnsApiKey
          - Deleted property: pdnsApiUrl
          - Deleted property: pdnsServerId

PUT /api/v1/dnscredentials
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: authMethod
          - Deleted property: authVariant
          - Deleted property: awsRoleArn
          - Deleted property: gcpServiceAccountEmail
          - Deleted property: pdnsApiKey
          - Deleted property: pdnsApiUrl
          - Deleted property: pdnsServerId
          - Modified property: domainFilter
            - Nullable changed from false to true
          - Modified property: name
            - Nullable changed from false to true
          - Modified property: providerType
            - Nullable changed from false to true

GET /api/v1/dnscredentials/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Required changed
                  - Deleted required property: authMethod
                  - Deleted required property: authVariant
                - Properties changed
                  - Deleted property: authMethod
                  - Deleted property: authVariant
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

POST /api/v1/dnscredentials/validate
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: authMethod
          - Deleted property: authVariant
          - Deleted property: awsRoleArn
          - Deleted property: certificateCredentialId
          - Deleted property: customCertificateAuthorityId
          - Deleted property: dnsCredentialId
          - Deleted property: gcpServiceAccountEmail
          - Deleted property: pdnsApiKey
          - Deleted property: pdnsApiUrl
          - Deleted property: pdnsServerId
          - Deleted property: tsigAlgorithm

GET /api/v1/flavors/projects/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Deleted property: gpuDetails

GET /api/v1/generic-kubernetes/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

POST /api/v1/googlecloud/create
- Request body changed
  - Content changed
    - Modified media type: multipart/form-data
      - Schema changed
        - Properties changed
          - New property: importedSubnetworkName
          - Deleted property: importedFilestoreProjectId
          - Deleted property: subnets
          - Modified property: billingAccountId
            - Nullable changed from false to true
          - Modified property: config
            - Nullable changed from false to true
          - Modified property: folderId
            - Nullable changed from false to true
          - Modified property: name
            - Nullable changed from false to true
          - Modified property: region
            - Nullable changed from false to true

POST /api/v1/googlecloud/filestore-list
- Request body changed
  - Content changed
    - Modified media type: multipart/form-data
      - Schema changed
        - Properties changed
          - Deleted property: importedFilestoreProjectId
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Items changed
            - Properties changed
              - Deleted property: ipAddress

GET /api/v1/googlecloud/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Required changed
                  - Deleted required property: createdBy
                  - Deleted required property: lastModified
                  - Deleted required property: lastModifiedBy
                - Properties changed
                  - New property: importedSubnetworkName
                  - Deleted property: createdBy
                  - Deleted property: lastModified
                  - Deleted property: lastModifiedBy

POST /api/v1/googlecloud/subnetwork-list
- Request body changed
  - Content changed
    - Modified media type: multipart/form-data
      - Schema changed
        - Properties changed
          - Deleted property: ipMode

POST /api/v1/internal/get-user-info
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Deleted property: userName

GET /api/v1/kubernetesprofiles/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

GET /api/v1/notifications/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: category
                    - Deleted enum values: [EnableGatewayApiModern DisableGatewayApiModern]
                  - Modified property: username
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

GET /api/v1/openshift/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

GET /api/v1/openstack/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

GET /api/v1/projects
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Required changed
                  - Deleted required property: gatewayApiModern
                - Properties changed
                  - Deleted property: gatewayApiModern
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

GET /api/v1/projects/visibility/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Required changed
            - Deleted required property: disableCertManager
            - Deleted required property: disableExternalDns
            - Deleted required property: disableGatewayApi
            - Deleted required property: disableTrustManager
            - Deleted required property: enableCertManager
            - Deleted required property: enableExternalDns
            - Deleted required property: enableGatewayApi
            - Deleted required property: enableTrustManager
          - Properties changed
            - Deleted property: disableCertManager
            - Deleted property: disableExternalDns
            - Deleted property: disableGatewayApi
            - Deleted property: disableTrustManager
            - Deleted property: enableCertManager
            - Deleted property: enableExternalDns
            - Deleted property: enableGatewayApi
            - Deleted property: enableTrustManager

GET /api/v1/prometheusbillings
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

GET /api/v1/proxmox/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

GET /api/v1/robot/details
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: createdBy
              - Type changed from 'object' to 'string'
              - Properties changed
                - Deleted property: accountId
                - Deleted property: displayName
                - Deleted property: isRobot
                - Deleted property: userId

GET /api/v1/robot/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

GET /api/v1/s3credentials/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

GET /api/v1/servers
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

GET /api/v1/servers/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
            - Modified property: project
              - Required changed
                - Deleted required property: gatewayApiModern
              - Properties changed
                - Deleted property: gatewayApiModern
                - Modified property: cloudSubnets
                  - Items changed
                    - Properties changed
                      - Deleted property: ipv6Cidr
                      - Modified property: availableIpCount
                        - Nullable changed from true to false

GET /api/v1/sse/recent-events/{organizationId}
- Modified query param: Category
  - Schema changed
    - Deleted enum values: [EnableGatewayApiModern DisableGatewayApiModern]

GET /api/v1/standalone/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
            - Modified property: project
              - Properties changed
                - Modified property: cloudSubnets
                  - Items changed
                    - Properties changed
                      - Deleted property: ipv6Cidr
                      - Modified property: availableIpCount
                        - Nullable changed from true to false

GET /api/v1/virtual-cluster/{parentProjectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
            - Modified property: project
              - Properties changed
                - Modified property: cloudSubnets
                  - Items changed
                    - Properties changed
                      - Deleted property: ipv6Cidr
                      - Modified property: availableIpCount
                        - Nullable changed from true to false

GET /api/v1/vsphere/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId

GET /api/v1/zadara/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: createdBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId
                  - Modified property: lastModifiedBy
                    - Type changed from 'object' to 'string'
                    - Nullable changed from false to true
                    - Properties changed
                      - Deleted property: accountId
                      - Deleted property: displayName
                      - Deleted property: isRobot
                      - Deleted property: userId


