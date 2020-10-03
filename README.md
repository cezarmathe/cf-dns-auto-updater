# cf-dns-auto-updater

**DEPRECATED - CHECK https://github.com/cezarmathe/qazx**

A tool that automatically updates Cloudflare A DNS records with new
IPs. It's a hack for overcoming a dynamic IP when you are hosting a
server.

# Install

## Using go

Run `go install github.com/cezarmathe/cf-dns-auto-updater`.

## Downloading a binary

Head to the [Releases](https://github.com/cezarmathe/cf-dns-auto-updater/releases)
page and download the binary for your machine.

# Usage

**cf-dns-auto-updater** makes use of a couple environment variables
described in the `.env` file. Use a systemd service or a cron job.
