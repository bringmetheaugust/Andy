# Andy🐼

## Description

Unix CLI package for personal tools.

## Install

Download `andy` file from `/build/andy` folder and paste to your OS executible path (`/usr/local/bin/` for example).

## Deploy

 ##### Required local packages: `Golang v1.15+`

 * `make build_app` build compiled package
 * `make dev` run development live server mode with hot reloading
 * `deploy_app` deploy compiled package to your `/usr/local/bin`

## Docs
 * #### Options💡

    * `--help`, `-h` list of commands, options and flags
    * `--version`, `-v` package version

 * #### Commands🎙

    * `server` run http daemon server to serve request info. Server API list [here](#guides_server)

        * `start` run server (default on http://localhost:1991)
        * `start -p NUMBER` set listening port
        * `status` get daemon status
        * `stop` server down

## Guides

<a name="#guides_server"></a>
<details>
   <summary><code>server</code> daemon API</summary>
   <ul>
        <li>
            <code>GET /docker/ps</code> to get docker containers status (equal to local <code>docker ps -a</code>)
        </li>
   </ul>
</details>
