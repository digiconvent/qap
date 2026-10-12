# cC9 Version

The version of QAP the client speaks. It is always the same as the version of the middleware it belongs to: the release workflow writes one version into both, and a test fails when they differ. [Version mismatch](../layer-4/version-mismatch.md) compares it with the version of the server on every [Discover](../layer-2/discover.md).

---

**Kind:** concept, `version`.

```ts
import { version } from '@digiconvent/qap-client'
```

`-1.-1.-1` between releases, such as `1.0.0` in a release.

---

**Referenced in:**

- [cA9 Version mismatch](../layer-4/version-mismatch.md)
- [cH2 Setup](../app/setup.md)
- [cS1 Discovered](../states/discovered.md)
- [cS7 Refused](../states/refused.md)
