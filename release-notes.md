### New Endpoints: None
-----------------------

### Deleted Endpoints: 2
------------------------
POST /api/v1/dns-cert/assign-certificate-profile  
POST /api/v1/dns-cert/unassign-certificate-profile  

### Modified Endpoints: 5
-------------------------
POST /api/v1/certificate-profiles
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Modified property: issuerMode
            - Deleted enum values: [letsencrypt venafi custom-ca self-signed none]
            - Nullable changed from false to true
          - Modified property: venafiType
            - Deleted enum values: [TPP Cloud]
            - Nullable changed from false to true

PUT /api/v1/certificate-profiles
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Modified property: issuerMode
            - Deleted enum values: [letsencrypt venafi custom-ca self-signed none]
            - Nullable changed from false to true
          - Modified property: venafiType
            - Deleted enum values: [TPP Cloud]
            - Nullable changed from false to true

GET /api/v1/certificate-profiles/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Deleted property: acmeEmail
                  - Deleted property: acmeServer
                  - Deleted property: customCaCert
                  - Deleted property: customCaPrivateKey
                  - Deleted property: venafiAccessToken
                  - Deleted property: venafiApiKey
                  - Deleted property: venafiBaseUrl
                  - Deleted property: venafiType
                  - Deleted property: venafiZone
                  - Deleted property: venafiZoneUuid
                  - Modified property: issuerMode
                    - Deleted enum values: [letsencrypt venafi custom-ca self-signed none]
                    - Nullable changed from false to true

POST /api/v1/googlecloud/create
- Request body changed
  - Content changed
    - Modified media type: multipart/form-data
      - Schema changed
        - Properties changed
          - Deleted property: ipMode

GET /api/v1/projects/visibility/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Required changed
            - Deleted required property: disableCTProfiles
            - Deleted required property: enableCTProfiles
          - Properties changed
            - Deleted property: disableCTProfiles
            - Deleted property: enableCTProfiles


