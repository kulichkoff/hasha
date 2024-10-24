# hasha

`hasha` is a simple command-line tool for generating hashes from input
data. This tool is useful for verifying data integrity and creating
secure checksums.

## Table of Contents

- [hasha](#hasha)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Contribute](#contribute)

## Installation

To install `hasha`, you can check "releases" page. Here you should download
a version for your platform.

## Usage

You can generate hash using different algorithms. Now the tool supports
SHA256 and MD5. SHA256 algorithm is used by dafault.

To calculate hash for the string just run:

```sh
hasha "some string"
```

or

```sh
hasha -a md5 "some string"
```

ALso, you can calculate hash for the file! You can specify hash algorithm
if necessary.

```sh
hasha path/to/file
```

Or use unix pipes:

```sh
cat file | hasha -a sha256
cat file | hasha
cat file | hasha -a md5
```

You can also use `hasha` output to save in a file:

```sh
hasha path/to/file > chksum_sha256
hasha -a md5 path/to/file > chksum_md5
```

## Contribute

Unfortunately, `hasha` is not ready for community contribution now.
You can make a feature erqurest or bug report on "Issues" page.
