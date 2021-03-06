Cache
=====

A simple cache utility.

**Honorable mentions**

* @benjy - Core concepts
* https://www.npmjs.com/package/npm-cache

## Setup

Configuration for this application is all done through the following YAML file:

```yaml
- directory: vendor
  hash:
  - composer.lock
  - Gemfile.lock
- directory: bin
  hash:
  - composer.lock
  - Gemfile.lock
- directory: node_modules
  hash:
  - packages.json
```

**Concepts**

* **Directory** - The directory which will be snapshotted and restored.
* **Hash** - The files which build up a unqiue hash.

## Usage

**Store folders in cache**

```bash
$ cache snapshot
```

**Check folders are cached**

```bash
$ cache list

HASH                            	FILES                     	DIRECTORY   	STATUS
----                            	-----                     	---------
1846fec897c7639c8303e35d8b9d6cad	composer.lock,Gemfile.lock	vendor      	Cached
1846fec897c7639c8303e35d8b9d6cad	composer.lock,Gemfile.lock	bin         	Cached
NULL                            	packages.json             	node_modules	Not Cached
```

**Restore from cache**

```bash
$ cache restore
```

