#!/bin/sh

# important
sudo apt update
sudo apt install -y apt-transport-https ca-certificates software-properties-common gnupg lsb-release debian-keyring debian-archive-keyring curl rsync

# repositories
sudo add-apt-repository -y ppa:neovim-ppa/stable # nvim
sudo add-apt-repository -y ppa:longsleep/golang-backports # go
sudo add-apt-repository -y ppa:ubuntu-toolchain-r/test # make
sudo apt-add-repository -y ppa:fish-shell/release-3 # fish
  # docker
sudo curl -fsSL 'https://download.docker.com/linux/ubuntu/gpg' | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
  # caddy
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list

sudo apt update
sudo apt install -y neovim golang-go make fish docker-ce caddy

# go
mkdir -p $HOME/go/bin

# fish
mkdir -p ~/.config/fish
rsync -a ~/zlatoivan_ru/config/init/config.fish ~/.config/fish/config.fish
chsh -s /usr/bin/fish $USER

# starship
curl -sS https://starship.rs/install.sh | sh -s -- -y
rsync -a ~/zlatoivan_ru/config/init/starship.toml ~/.config/starship.toml

# caddy
rsync -a ~/zlatoivan_ru/config/init/Caddyfile /etc/caddy/Caddyfile
sudo systemctl restart caddy

# docker compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# ssh (убрать приветствие)
sudo chmod -x /etc/update-motd.d/*

# new terminal
fish