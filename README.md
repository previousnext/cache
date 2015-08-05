Cache
=====

A simple cache utility.

**Honorable mentions**

* @benjy - Core concepts
* https://www.npmjs.com/package/npm-cache

## Setup

Configuration for this application is all done through the following YAML file:

```yaml
- hash_file:
    - composer.lock
    - Gemfile.lock
  restore:
    - vendor
- hash_file:
    - packages.json
  restore:
    - node_modules
```

**Concepts**

* **Hash file** - The files used to generate a hash for computing if the cache has invalidated.
* **Restore** - A list of directories used for restore.

## Usage

**Store folders in cache**

```bash
$ cache snapshot
```

**Check folders are cached**

```bash
$ cache status
                HASH               |     FILE      |    RESTORE    |   STATUS
+----------------------------------+---------------+---------------+------------+
  f7fa752830592c66ed7fd258876ebf38 | Gemfile.lock  | vendor/bundle | Not cached
  NULL                             | composer.lock | vendor        | Not cached
  NULL                             | packages.json | node_modules  | Not cached
```

**Restore from cache**

```bash
$ cache restore
```

