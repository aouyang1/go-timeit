[![GoDoc](https://godoc.org/github.com/aouyang1/go-timeit?status.svg)](https://godoc.org/github.com/aouyang1/go-timeit)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

# go-timeit

Golang library to measure execution time of code.

## Contents
- [Installation](#installation)
- [Quick start](#quick-start)
- [Contributing](#contributing)
- [Contact](#contact)
- [License](#license)

## Installation
```sh
$ go get -u github.com/aouyang1/go-timeit
```

## Quick start
```go
package main

import (
  "time"

  "github.com/aouyang1/go-timeit"
)

func main() {
  t := timeit.New()
  for i := 0; i < 10; i++ {
    t.Tic()
    myFunc()
    t.Toc()
  }
  t.Show()
  // 0.24 s ± 0 ms per loop (mean ± std. dev. of 10 runs, 1 loop each)
}

func myFunc() {
  time.Sleep(234 * time.Millisecond)
}
```

## Contributing
* Fork the repository
* Create a new branch (feature_\* or bug_\*)for the new feature or bug fix
* Run tests
* Commit your changes
* Push code and open a new pull request

## Contact
* Austin Ouyang (aouyang1@gmail.com)

## License
The MIT License (MIT). See [LICENSE](https://github.com/matrix-profile-foundation/go-timeit/blob/master/LICENSE) for more details.

Copyright (c) 2018 Austin Ouyang
