### New Endpoints: None
-----------------------

### Deleted Endpoints: 4
------------------------
GET /api/v1/internal/prometheus-exporter/accounts  
GET /api/v1/internal/prometheus-exporter/projects  
GET /api/v1/internal/prometheus-exporter/servers  
GET /api/v1/internal/prometheus-exporter/standalone  

### Modified Endpoints: 4
-------------------------
GET /api/v1/accounts/{accountId}/groups/offset-based/dropdown
- Deleted query param: AccessLevel
- Deleted query param: OrganizationId
- Deleted query param: SortBy
- Deleted query param: SortDirection
- Deleted query param: UserId
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - New property: id
                  - New property: name
                  - Deleted property: domainId
                  - Deleted property: groupId
                  - Deleted property: groupName
                  - Deleted property: organizationsCount
                  - Deleted property: roles
                  - Deleted property: usersCount

GET /api/v1/accounts/{accountId}/organizations/offset-based/with/group
- Deleted query param: ProjectId
- Deleted query param: SortDirection
- Deleted query param: UserId
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - New property: accessLevel
                  - New property: isBound
                  - New property: orgId
                  - New property: orgName
                  - New property: projects
                  - Deleted property: domainId
                  - Deleted property: groupsCount
                  - Deleted property: organizationId
                  - Deleted property: organizationName
                  - Deleted property: projectsCount
                  - Deleted property: usersCount

GET /api/v1/accounts/{accountId}/projects/offset-based/dropdown
- Deleted query param: GroupId
- Deleted query param: Health
- Deleted query param: OrganizationId
- Deleted query param: SortBy
- Deleted query param: SortDirection
- Deleted query param: Status
- Deleted query param: UserId
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - New property: id
                  - New property: name
                  - Deleted property: cloudType
                  - Deleted property: domainId
                  - Deleted property: groupsCount
                  - Deleted property: health
                  - Deleted property: organizationId
                  - Deleted property: organizationName
                  - Deleted property: projectId
                  - Deleted property: projectName
                  - Deleted property: status
                  - Deleted property: usersCount

GET /api/v1/accounts/{accountId}/user/offset-based/dropdown
- Deleted query param: FilterGroupId
- Deleted query param: FilterOrganizationId
- Deleted query param: FilterProjectId
- Deleted query param: FilterRole
- Deleted query param: SortBy
- Deleted query param: SortDirection
- Modified query param: Limit
  - Schema changed
    - Default changed from 50 to 100
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Required changed
                  - Deleted required property: globalRole
                  - Deleted required property: groupsCount
                  - Deleted required property: nonGlobalRoles
                  - Deleted required property: organizationsCount
                  - Deleted required property: projectsCount
                  - Deleted required property: userId
                  - Deleted required property: userName
                - Properties changed
                  - New property: displayName
                  - New property: id
                  - New property: name
                  - New property: timestamp
                  - Deleted property: globalRole
                  - Deleted property: groupsCount
                  - Deleted property: nonGlobalRoles
                  - Deleted property: organizationsCount
                  - Deleted property: projectsCount
                  - Deleted property: userId
                  - Deleted property: userName


