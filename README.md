# Keel Mailer

Keel Mailer is a simple helper for [Keel](https://keel.sh) that sends email when there is a new approval waiting.

## Installation

1. Configure Keel as usual.
2. Deploy Keel Mailer, see [docs/keel-mailer.yaml](docs/keel-mailer.yaml) for an example YAML file. You need to set the correct credentials, preferably as Kubernetes Secrets. See the available environmental variables in the YAML file.
3. Enjoy getting e-mail from Keel!

## License

    keel-mailer
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
