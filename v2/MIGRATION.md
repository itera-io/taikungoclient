# Migrating to taikungoclient v2

v2 is generated from the post-3.1 Taikun API. Summary of breaking changes (full per-endpoint
diff: `oasdiff changelog` between the old 3.0.4 spec and the new spec).

## Removed endpoints (~84) — remove all usage
- Cloud providers: **vSphere, Proxmox, OpenShift, Zadara** (credentials, images, flavors).
- Billing/subscription: **payment, subscription, stripe, organization-subscriptions,
  infra-billing-summary, projectinfracost**.
- **Invoices**, **project-templates / server-templates**, **taikun-lb**.

## Account → Domain rename (25 endpoints)
- API/type names: `Account*` → `Domain*` (e.g. `DomainsAPI`, `DomainDetailsDto`,
  `DomainEntity`, `DomainList`, `DomainSsoConfigDto`).
- Path param `accountId` → `domainId`; routes `/api/v1/accounts/*` → `/api/v1/domains/*`.
- Fields: `AccountId`/`AccountName` → `DomainId`/`DomainName`; auth login field
  `accountName` → `domainName` (`LoginCommand.DomainName`).

## Enum value renames (int values unchanged: 2/3)
- `EGlobalRole`, `EFilteredUserRole`, `GlobalRole`: `AccountAdmin` → `DomainAdmin`,
  `AccountOwner` → `DomainOwner` (Go constants `EGLOBALROLE_DOMAIN_ADMIN`, etc.).

## Generator/shape changes
- Client is generated from a **normalized 3.0** spec (`openapi-codegen.json`): numeric ids
  stay plain `integer` (not `integer|string`), nullable maps are `map[...]` + pointer,
  single-element `oneOf` are flattened. This changes some pointer/optional shapes vs the old
  client even on unchanged endpoints — expect compile adjustments across the board.

## Coexistence
Old code keeps importing `github.com/itera-io/taikungoclient` (pin the last pre-v2 tag/commit);
new code imports `github.com/itera-io/taikungoclient/v2`. Migrate package-by-package.
