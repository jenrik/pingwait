# pingwait

`pingwait` continuously ping a IP address until it either times out, or receives a reply. This is useful for running a command once a server comes online.

## Usage

Example usage: `pingwait -t 10 1.1.1.1`

`-t` Optional argument specifying the amount of seconds before pingwait times out. Defaults to never timing out.

## Exit codes

0: Received reply<br>
1: Invalid argument specified<br>
2: Timeout was reached<br>
50: Internal error

## Note
This command has to either be run as root or have the capability `CAP_NET_RAW`. The capability can be given to `pingwait` by running `setcap cap_net_raw=+ep pingwait`
