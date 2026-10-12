# A12 Version

QAP carries a version value, so it can tell which version of the protocol it speaks. [Discover](../layer-2/discover.md) sends it in the `X-QAP-Version` response header, so a client knows what it talks to before it proposes anything. The client carries the same version and refuses a server that speaks another one.

---

**Kind:** atomic part.

```go
var Version = "-1.-1.-1"
```

`-1.-1.-1` is the version outside a release. Publishing a release on GitHub, with a tag named like `v1.0.0`, starts the release workflow. It checks out the commit of that tag, writes the version into this constant, into the version of the client and into its `package.json`, runs the tests of both, commits "Release 1.0.0" on top, and moves the release tag to that commit. So the release itself carries its version. The Go module sits at the root of the repository, so Go finds the release by that same tag; a tag without the `v` is refused, because Go would not find it. `main` is not changed. The workflow can also be started by hand, with the tag as input. A test fails when the middleware and the client carry different versions.

---

**Referenced in:**

- [I1 Discover](../layer-2/discover.md)
