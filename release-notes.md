### New Endpoints: None
-----------------------

### Deleted Endpoints: None
---------------------------

### Modified Endpoints: 4
-------------------------
POST /api/v1/custom-cas
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Modified property: caCert
            - Nullable changed from false to true
          - Modified property: caKey
            - Nullable changed from false to true
          - Modified property: name
            - Nullable changed from false to true

POST /api/v1/dns-cert/validate
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Modified property: certIssuerType
            - Nullable changed from false to true
          - Modified property: dnsProvider
            - Nullable changed from false to true
          - Modified property: domain
            - Nullable changed from false to true

POST /api/v1/dnscredentials
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Modified property: domainFilter
            - Nullable changed from false to true
          - Modified property: name
            - Nullable changed from false to true
          - Modified property: providerType
            - Nullable changed from false to true

POST /api/v1/dnscredentials/validate
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Modified property: certIssuerType
            - Nullable changed from false to true
          - Modified property: dnsProvider
            - Nullable changed from false to true
          - Modified property: domain
            - Nullable changed from false to true


