# Keel Webhook Mailer

Keel Webhook Mailer (KWA) is a simple helper for [Keel](https://keel.sh) that exposes a webhook that sends an email for Keel notifications.

If KWA is deployed as a service in the cluster, you can set the `WEBHOOK_ENDPOINT` variable in Keel to `http://kwa.keel.svc.cluster.local/webhook` and an e-mail will be sent for every notification.

## Installation

1. Configure Keel as usual, and set the `WEBHOOK_ENDPOINT` variable to `http://kwa.keel.svc.cluster.local/webhook`.
2. Deploy KWA, see [docs/kwa.yaml](docs/kwa.yaml) for an example YAML file. You need to set the correct credentials, preferably as Kubernetes Secrets. See the available environmental variables in the YAML file.
3. Enjoy getting e-mail from Keel!

## License

    keel-webhook-mailer
    Copyright (C) 2021  Linus Karlsson

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
