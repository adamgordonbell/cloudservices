package main

import (
    "dagger.io/dagger"
)

dagger.#Plan & {
    // Say hello by writing to a file
    actions: hello: #AddHello & {
        dir: client.filesystem.".".read.contents
    }
    client: filesystem: ".": {
        read: contents: dagger.#FS
        write: contents: actions.hello.result
    }
}