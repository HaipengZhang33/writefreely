&nbsp;
<p align="center">
	<a href="https://writefreely.org"><img src="https://writefreely.org/img/writefreely.svg" width="350px" alt="Write Freely" /></a>
</p>
<hr />
<p align="center">
	<a href="https://github.com/writeas/writefreely/releases/">
		<img src="https://img.shields.io/github/release/writeas/writefreely.svg" alt="Latest release" />
	</a>
	<a href="https://goreportcard.com/report/github.com/writeas/writefreely">
		<img src="https://goreportcard.com/badge/github.com/writeas/writefreely" alt="Go Report Card" />
	</a>
	<a href="https://travis-ci.org/writeas/writefreely">
		<img src="https://travis-ci.org/writeas/writefreely.svg" alt="Build status" />
	</a>
	<a href="http://webchat.freenode.net/?channels=writefreely">
		<img alt="#writefreely on freenode" src="https://img.shields.io/badge/freenode-%23writefreely-blue.svg" />
	</a>
</p>
&nbsp;

WriteFreely is a beautifully pared-down blogging platform that's simple on the surface, yet powerful underneath.

It's designed to be flexible and share your writing widely, so it's built around plain text and can publish to the _fediverse_ via ActivityPub. It's easy to install and lightweight.

**[Start a blog on our instance](https://write.as/new/blog/federated)**

[Try the editor](https://write.as/new)

[Find another instance](https://writefreely.org/instances)

## Features

* Start a blog for yourself, or host a community of writers
* Form larger federated networks, and interact over modern protocols like ActivityPub
* Write on a dead-simple, distraction-free and super fast editor
* Publish drafts and let others proofread them by sharing a private link
* Build more advanced apps and extensions with the [well-documented API](https://developers.write.as/docs/api/)

## Quick start

> **Note** this is currently alpha software. We're quickly moving out of this v0.x stage, but while we're in it, there are no guarantees that this is ready for production use.

First, download the [latest release](https://github.com/writeas/writefreely/releases/latest) for your OS. It includes everything you need to start your blog.

Now extract the files from the archive, change into the directory, and do the following steps:

```bash
# 1) Log into MySQL and run:
# CREATE DATABASE writefreely;
#
# 2) Import the schema with:
mysql -u YOURUSERNAME -p writefreely < schema.sql

# 3) Configure your blog
./writefreely --config

# 4) Generate data encryption keys
./writefreely --gen-keys

# 5) Run
./writefreely

# 6) Check out your site at the URL you specified in the setup process
# 7) There is no Step 7, you're done!
```

For running in production, [see our guide](https://writefreely.org/start#production).

## Development

Ready to hack on your site? Here's a quick overview.

### Prerequisites

* [Go 1.10+](https://golang.org/dl/)
* [Node.js](https://nodejs.org/en/download/)

### Setting up

```bash
go get github.com/writeas/writefreely/cmd/writefreely
```

Create your database, import the schema, and configure your site [as shown above](#quick-start). Then generate the remaining files you'll need:

```bash
make install # Generates encryption keys; installs LESS compiler
make ui      # Generates CSS (run this whenever you update your styles)
make run     # Runs the application
```

### Using Docker

From the cloned git repository, you can quickly stand up a Write Freely instance with Docker and Docker Compose.

First, you'll need to change the password for MariaDB's root user in `docker-compose.yml` from `changeme` to something that is unique to your setup:

```
environment:
  - MYSQL_ROOT_PASSWORD=changeme
```

After that, you can spin up the containers and configure them:

```bash
# 1) Spin up the DB and Write Freely
docker-compose up -d

# 2) Connect to MariaDB container
docker-compose exec db /bin/sh

# 3) Log in to MariaDB, using the password you specified in docker-compose.yml
mysql -u root -p

# 4) Create the database for Write Freely
CREATE DATABASE writefreely;
exit

# 5) Migrate the database
mysql -u root -p writefreely < /tmp/schema.sql
exit

# 6) Generate the configuration and clean up
docker-compose run web writefreely --config
docker stop writefreely_web_run_1 && docker rm writefreely_web_run_1
```

Now you should be able to navigate to http://localhost:8080 and start blogging!

## License

Licensed under the AGPL.
