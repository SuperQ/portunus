# Portunus

Portunus was the ancient Roman god of keys and doors. However, this repo does not
contain the god. It contains Portunus, a small and self-contained user/group
management and authentication service.

TODO explain more

## Running

Once installed, run `portunus-orchestrator` with root privileges. Config is passed to it via the
following environment variables:

| Variable | Default | Explanation |
| -------- | ------- | ----------- |
| `PORTUNUS_LDAP_SUFFIX` | *(required)* | The DN of the topmost entry in your LDAP directory. Must currently be a sequence of `dc=xxx` RDNs. (This requirement may be lifted in future versions.) See [*LDAP directory structure*](#ldap-directory-structure) for details. |
| `PORTUNUS_SERVER_BINARY` | `portunus-server` | Where to find the portunus-server binary. Semantics match those of `execvp(3)`: If the supplied value is not a path containing slashes, `$PATH` will be searched for it. |
| `PORTUNUS_SERVER_GROUP`<br>`PORTUNUS_SERVER_USER` | `portunus` each | The Unix user/group that Portunus' own server will be run as. |
| `PORTUNUS_SERVER_STATE_DIR` | `/var/lib/portunus` | The path where Portunus stores its database. **Set up a backup for this directory.** |
| `PORTUNUS_SLAPD_BINARY` | `slapd` | Where to find the binary of slapd (the OpenLDAP server). Semantics match those of `execvp(3)`: If the supplied value is not a path containing slashes, `$PATH` will be searched for it. |
| `PORTUNUS_SLAPD_GROUP`<br>`PORTUNUS_SLAPD_USER` | `ldap` each | The Unix user/group that slapd will be run as. |
| `PORTUNUS_SLAPD_SCHEMA_DIR` | `/etc/openldap/schema` | Where to find OpenLDAP's schema definitions. |
| `PORTUNUS_SLAPD_STATE_DIR` | `/var/run/portunus-slapd` | The path where slapd stores its database. The contents of this directory are ephemeral and will be wiped when Portunus restarts, so you do not need to back this up. Place this on a tmpfs for optimal performance. |

Root privileges are required for the orchestrator because it needs to setup runtime directories and
bind the LDAP port which is a privileged port (389 without TLS, 636 with TLS). No process managed by
Portunus will offer a network service while running as root:

- LDAP and LDAPS are offered by slapd which is running as `ldap:ldap` by default.
- HTTP is offered by `portunus-server` which is running as `portunus:portunus` by default.

### LDAP directory structure

TODO
