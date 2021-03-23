# Andy🐼

## Description

Unix CLI package for personal tools.

## Docs
 * #### Options

    * `--help`, `-h` list of commands, options and flags
    * `--version`, `-v` package version

 * #### Commands

    * `server` run http daemon server to serve request info. Server API list [here](#guides_server)

        * `start` run server (default on http://localhost:1991)
        * `start -p NUMBER` set listening port
        * `status` get daemon status
        * `stop` server down

## Tools

 * `make build_app` build package
 * `make deploy_app` move compiled package to `/usr/local/bin` folder
 * `make dev` to start hot-reload development mode

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
