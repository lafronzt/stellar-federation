# Stellar Federation Server

One of the beautiful features of Stellar is the ability to create a federation server. This allows you to send payments to accounts on the network without having to know the public key of the account.

## Technical Details

### Architecture

- Built with Go.
- Containerized
- Federation server is a REST API that is exposed to the public.
- Uses the [Stellar Network Foundation](https://developers.stellar.org/docs/glossary/federation/) protocol.

### Features

- Dynamic Memo support
    > Add any memo to any transaction by sending the payment to `tyler+Thank_You*lafronz.com`. The text following `+` is the memo.
    > 
    > Learn more about Memos at [lumenauts.com](https://www.lumenauts.com/explainers/what-are-memos)
- Supports multiple federation domains per servers
- Supports ID to federation reverse lookups

## How to use

1. Create your `stellar.yaml` file

    > A sample file is show in [the repo](stellar.yaml).

2. Create your `stellar.toml` file

    > A sample file is show in [the repo](stellar.toml).

3. Upload your `stellar.toml` files so it can be found by the Stellar Network at `https://YOUR_DOMAIN/.well-known/stellar.toml`

4. Build your Federation Server

    > Replace the stellar.yaml file with your own.
    > 
    > Run `docker build -t docker.io/<username>/stellar-federation-server:latest .`

5. Deploy your Federation Server to your runtime location of your choosing.

    > Sample K8s deployment is show in [the repo](Deployment/k8s.yaml).

## More Information

More information about the federation server can be found in the [Stellar Network Foundation](https://developers.stellar.org/docs/glossary/federation/) documentation.
