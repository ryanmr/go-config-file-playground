# this is go - do not look

I wanted to parse a `config.yml` file, and eventually generate a new [envoy proxy configuration file](https://www.envoyproxy.io/).

```
go run .

go build .

./generator
```

TODO:

* use [a json schema validator](https://json-schema.org/implementations.html) on the `config.yml` to enforce shapes and rules
* externalize template file from the source

---

Disclaimer:

* I do not know go.
* I did not get very far.
  * This was conceptually ported from a spur of the moment python script.
  * This was eventually ported to `node` javascript with typescript and [swc](https://swc.rs/).