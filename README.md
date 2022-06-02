PEQ-Expansions is an effort to make the peq db more progression/era friendly


# Scripts

- make build-dropshift: creates the dropshift binary, places it inside drop. (`usage: dropshift <search> <replace>`, This binary searches for a pattern in *_lde.sql files and places it in *_lte.sql files, "shifting" items)
- make build-dropstrip: creates the dropstrip binary, places it inside drop. (`usage: dropstrip <search>`, This binary searches for a pattern in *_lde.sql files and removes them only)
- make build-dropexport: creates the dropexport binary, places it inside drop. (`usage: drop <zone_shortname>`, This binary queries a database's zone and generates the .sql files noted in drop)