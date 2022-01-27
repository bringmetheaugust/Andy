# Andy🐼

## Description

Unix CLI package for personal tools.

## Install

Download `andy` file from `/build/andy` folder and paste to your OS executible path (`/usr/local/bin/` for example).

## Deploy

 * `make build` build compiled package
 * `deploy` deploy compiled package to your `/usr/local/bin`

 ##### Required global packages: `Golang v1.17+`

## Docs

 * #### Commands🎙

    * `docker_http` run [http server](#guides_server) for getting docker containers info.

        * `start` run server (default on [http://localhost:1991]())
        * `start -p NUMBER` set listening port

    * `window` open GUI interface

    * `init` init package for more friendly usage (generating shell autocomplete)

        * `-s STRING` set shell type for completion (autocomplete)

 * #### Options💡

    * `--help`, `-h` list of commands, options and flags
    * `--version`, `-v` package version

## Guides

<a name="#guides_server"></a>
<details>
   <summary><code>docker http</code> server API</summary>
   <ul>
        <li>
            <details>
               <summary><code>GET /ps</code> to get docker containers status (equal to local <code>docker ps -a</code>)</summary>
               <ul>
                  <li><code>200</code>: success result</li>
                  <li><code>406</code>: failed docker daemon info</li>
               </ul>
            </details>
        </li>
   </ul>
</details>
