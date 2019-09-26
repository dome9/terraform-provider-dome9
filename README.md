# terraform-provider-dome9
Terraform check point cloud guard dome9 provider

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.svg)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-dome9`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-dome9.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-dome9
$ make build
# or if you're on a mac:
$ gnumake build
```

Using the provider
----------------------

Detailed documentation for the Dome9 provider can be found [here](https://www.terraform.io/docs/providers/dome9/index.html).

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.13+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-dome9
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

Acceptance test prerequisites
-----------------------------
In order to successfully run the full suite of acceptance tests, you will need to have the following:

### Dome9 personal access id and secret key
You will need to create a [Dome9 access id & secret key](https://secure.dome9.com/v2/settings/credentials) for
testing. It will need to have a full admin access.

### acceptance test full environment variables list:
Dome9 is a security product, in order to manged the supported clouds a sensitive data must be provided in the on-board staging 
this data is passed using exported environment variables, your environment must set the following:


#### Dome9 environment variables:
- `DOME9_ACCESS_ID=;`
- `DOME9_SECRET_KEY=;`

#### AWS environment variables:
- `ARN=;`
- `SECRET=;`

#### Azure environment variables:
- `SUBSCRIPTION_ID=;`
- `TENANT_ID=;`
- `CLIENT_PASSWORD=;`

#### GCP environment variables:
- `PROJECT_ID`
- `PRIVATE_KEY=;`
- `PRIVATE_KEY_ID=;`
- `CLIENT_EMAIL=;`
- `CLIENT_ID=;`
- `CLIENT_X509_CERT_URL=;`
